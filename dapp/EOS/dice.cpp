#include "dice.hpp";

using namespace dicecasino;

void dice::init() {
    require_auth( _self );
    auto configLookup = configs.find(1); //id is always 1
    if (configLookup == configs.end()){
        configs.emplace(_self, [&](auto& newConfig) {
            newConfig.configid  = 1;
            newConfig.locked = asset(double(0), EOS_SYMBOL);
            newConfig.gamestatus = 0;
            newConfig.currentid  = 0;
            newConfig.totalbet = 0;
            newConfig.totalwin = 0;
            newConfig.minbetamount = 1000;
            newConfig.maxbetamount = 500000;
        });
    } else {
        configs.modify(configLookup, _self, [&](auto& config) {
            config.gamestatus = 0;
            config.locked = asset(double(0), EOS_SYMBOL);
            config.currentid  = 0;
            config.totalbet = 0;
            config.totalwin = 0;
            config.minbetamount = 1000;
            config.maxbetamount = 500000;
        });
    }
}

void dice::res(const capi_checksum160& user_seed, const capi_checksum256& seed) {
    uint8_t random_roll = compute_random_roll(seed, user_seed);
    print(random_roll);
}

void dice::reveal(const uint64_t& id, const capi_checksum256& seed) {
    require_auth(_self);
    bet bet = find_or_error(id);
    assert_sha256( (char *)&seed, sizeof(seed), (const capi_checksum256 *)&bet.seed_hash );

    uint8_t random_roll = compute_random_roll(seed, bet.user_seed_hash);
    asset payout = asset(0, EOS_SYMBOL);
    if (random_roll < bet.roll_under) {
        payout = compute_payout(bet.roll_under, bet.amount);
        add_total_win(payout);
        action(permission_level{_self, name("active")},
               name("eosio.token"),
               name("transfer"),
               make_tuple(_self, bet.player, payout, winner_memo(bet)))
            .send();
    }
    unlock(bet.amount);
    if (bet.referrer != _self) {
        // defer trx, no need to rely heavily
        send_defer_action(permission_level{_self, name("active")},
                          name("eosio.token"),
                          name("transfer"),
                          make_tuple(_self,
                                     bet.referrer,
                                     compute_referrer_reward(bet),
                                     referrer_memo(bet)));
    }
    remove_bet(bet);
    const result _result{.bet_id = bet.id,
                     .player = bet.player,
                     .referrer = bet.referrer,
                     .amount = bet.amount,
                     .roll_under = bet.roll_under,
                     .random_roll = random_roll,
                     .seed = seed,
                     .seed_hash = bet.seed_hash,
                     .user_seed_hash = bet.user_seed_hash,
                     .payout = payout};
    save_result(_result);
}

void dice::transfer(name from,name to,asset quantity,string memo) {
    if (from == _self || to != _self) {
        return;
    }

    uint8_t roll_under;
    capi_checksum256 seed_hash;
    capi_checksum160 user_seed_hash;
    uint64_t expiration;
    name referrer;
    capi_signature sig;

    parse_memo(memo, &roll_under, &seed_hash, &user_seed_hash, &expiration, &referrer, &sig);

    //check game
    assert_game_status();

    //check quantity
    assert_quantity(quantity);

    //check roll_under
    assert_roll_under(roll_under, quantity);

    // check seed hash && expiration
    assert_hash(seed_hash, expiration);

    //check referrer
    eosio_assert(referrer != from, "referrer can not be self");

    //check signature
    assert_signature(roll_under, seed_hash, expiration, referrer, sig);

    const bet _bet{.id = next_id(),
                  .player = from,
                  .referrer = referrer,
                  .amount = quantity,
                  .roll_under = roll_under,
                  .seed_hash = seed_hash,
                  .user_seed_hash = user_seed_hash,
                  .created_at = now()};
    save_bet(_bet);
    add_total_bet(quantity);
    lock(quantity);
}

void dice::closegame() {
    require_auth( _self );
    auto configLookup = configs.find(1);

    configs.modify(configLookup, _self, [&](auto& c) {
        c.gamestatus = 1;
    });
}

void dice::opengame() {
    require_auth( _self );
    auto configLookup = configs.find(1);

    configs.modify(configLookup, _self, [&](auto& c) {
        c.gamestatus = 0;
    });
}

void dice::clearconfig() {
    require_auth(_self);
    for(auto itr = configs.begin(); itr != configs.end();) {
        itr = configs.erase(itr);
    }
}

void dice::clearresult() {
    require_auth(_self);
    for(auto itr = results.begin(); itr != results.end();) {
        itr = results.erase(itr);
    }
}

extern "C" void apply(uint64_t receiver, uint64_t code, uint64_t action)
{
	if (code == "eosio.token"_n.value && action == "transfer"_n.value)
	{
		eosio::execute_action(
			eosio::name(receiver), eosio::name(code), &dice::transfer
		);
	}
	else if (code == receiver)
	{
		switch (action)
		{
			EOSIO_DISPATCH_HELPER(dice, (init)(reveal)(opengame)(closegame)(clearresult)(clearconfig)(res));
		}
	}
}