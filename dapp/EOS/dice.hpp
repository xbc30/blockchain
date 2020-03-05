#include <algorithm>
#include <eosiolib/transaction.hpp>
#include <eosiolib/crypto.h>
#include <eosiolib/asset.hpp>
#include <eosiolib/eosio.hpp>
#include <eosiolib/singleton.hpp>
#include <eosiolib/time.hpp>

using namespace eosio;
using namespace std;

#define EOS_SYMBOL symbol("EOS", 4)

static const string PUB_KEY = "EOS8X3wChnMTX6fBwpSzxqkRmw1wRjUPbRf8oCiENqQNjzwtGCnTP";

/** All alphanumeric characters except for "0", "I", "O", and "l" */
static const char* pszBase58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz";
static const int8_t mapBase58[256] = {
    -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
    -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
    -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0,  1,  2,  3,  4,  5,  6,  7,
    8,  -1, -1, -1, -1, -1, -1, -1, 9,  10, 11, 12, 13, 14, 15, 16, -1, 17, 18,
    19, 20, 21, -1, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, -1, -1, -1, -1,
    -1, -1, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, -1, 44, 45, 46, 47, 48,
    49, 50, 51, 52, 53, 54, 55, 56, 57, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
    -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
    -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
    -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
    -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
    -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
    -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
    -1, -1, -1, -1, -1, -1, -1, -1, -1,
};

string uint64_string(uint64_t input) {
    string result;
    uint8_t base = 10;
    do {
        char c = input % base;
        input /= base;
        if (c < 10)
            c += '0';
        else
            c += 'A' - 10;
        result = c + result;
    } while (input);
    return result;
}

uint8_t from_hex(char c) {
    if (c >= '0' && c <= '9') return c - '0';
    if (c >= 'a' && c <= 'f') return c - 'a' + 10;
    if (c >= 'A' && c <= 'F') return c - 'A' + 10;
    eosio_assert(false, "Invalid hex character");
    return 0;
}

size_t from_hex(const string& hex_str, char* out_data, size_t out_data_len) {
    auto i = hex_str.begin();
    uint8_t* out_pos = (uint8_t*)out_data;
    uint8_t* out_end = out_pos + out_data_len;
    while (i != hex_str.end() && out_end != out_pos) {
        *out_pos = from_hex((char)(*i)) << 4;
        ++i;
        if (i != hex_str.end()) {
            *out_pos |= from_hex((char)(*i));
            ++i;
        }
        ++out_pos;
    }
    return out_pos - (uint8_t*)out_data;
}

string to_hex(const char* d, uint32_t s) {
    std::string r;
    const char* to_hex = "0123456789abcdef";
    uint8_t* c = (uint8_t*)d;
    for (uint32_t i = 0; i < s; ++i)
        (r += to_hex[(c[i] >> 4)]) += to_hex[(c[i] & 0x0f)];
    return r;
}

string sha256_to_hex(const capi_checksum256& sha256) {
    return to_hex((char*)sha256.hash, sizeof(sha256.hash));
}

string sha1_to_hex(const capi_checksum160& sha1) {
    return to_hex((char*)sha1.hash, sizeof(sha1.hash));
}

uint64_t uint64_hash(const string& hash) {
    return std::hash<string>{}(hash);
}

uint64_t uint64_hash(const capi_checksum256& hash) {
    return uint64_hash(sha256_to_hex(hash));
}

capi_checksum256 hex_to_sha256(const string& hex_str) {
    eosio_assert(hex_str.length() == 64, "invalid sha256");
    capi_checksum256 checksum;
    from_hex(hex_str, (char*)checksum.hash, sizeof(checksum.hash));
    return checksum;
}

capi_checksum160 hex_to_sha1(const string& hex_str) {
    eosio_assert(hex_str.length() == 40, "invalid sha1");
    capi_checksum160 checksum;
    from_hex(hex_str, (char*)checksum.hash, sizeof(checksum.hash));
    return checksum;
}

size_t sub2sep(const string& input, string* output, const char& separator, const size_t& first_pos = 0, const bool& required = false) {
    eosio_assert(first_pos != string::npos, "invalid first pos");
    auto pos = input.find(separator, first_pos);
    if (pos == string::npos) {
        eosio_assert(!required, "parse memo error");
        return string::npos;
    }
    *output = input.substr(first_pos, pos - first_pos);
    return pos;
}

bool DecodeBase58(const char* psz, std::vector<unsigned char>& vch) {
    // Skip leading spaces.
    while (*psz && isspace(*psz)) psz++;
    // Skip and count leading '1's.
    int zeroes = 0;
    int length = 0;
    while (*psz == '1') {
        zeroes++;
        psz++;
    }
    // Allocate enough space in big-endian base256 representation.
    int size = strlen(psz) * 733 / 1000 + 1;  // log(58) / log(256), rounded up.
    std::vector<unsigned char> b256(size);
    // Process the characters.
    static_assert(
        sizeof(mapBase58) / sizeof(mapBase58[0]) == 256,
        "mapBase58.size() should be 256");  // guarantee not out of range
    while (*psz && !isspace(*psz)) {
        // Decode base58 character
        int carry = mapBase58[(uint8_t)*psz];
        if (carry == -1)  // Invalid b58 character
            return false;
        int i = 0;
        for (std::vector<unsigned char>::reverse_iterator it = b256.rbegin();
             (carry != 0 || i < length) && (it != b256.rend());
             ++it, ++i) {
            carry += 58 * (*it);
            *it = carry % 256;
            carry /= 256;
        }
        assert(carry == 0);
        length = i;
        psz++;
    }
    // Skip trailing spaces.
    while (isspace(*psz)) psz++;
    if (*psz != 0) return false;
    // Skip leading zeroes in b256.
    std::vector<unsigned char>::iterator it = b256.begin() + (size - length);
    while (it != b256.end() && *it == 0) it++;
    // Copy result into output vector.
    vch.reserve(zeroes + (b256.end() - it));
    vch.assign(zeroes, 0x00);
    while (it != b256.end()) vch.push_back(*(it++));
    return true;
}

bool decode_base58(const string& str, vector<unsigned char>& vch) {
    return DecodeBase58(str.c_str(), vch);
}
// Copied from https://github.com/bitcoin/bitcoin

capi_signature str_to_sig(const string& sig, const bool& checksumming = true) {
    const auto pivot = sig.find('_');
    eosio_assert(pivot != string::npos, "No delimiter in signature");
    const auto prefix_str = sig.substr(0, pivot);
    eosio_assert(prefix_str == "SIG", "Signature Key has invalid prefix");
    const auto next_pivot = sig.find('_', pivot + 1);
    eosio_assert(next_pivot != string::npos, "No curve in signature");
    const auto curve = sig.substr(pivot + 1, next_pivot - pivot - 1);
    eosio_assert(curve == "K1" || curve == "R1", "Incorrect curve");
    const bool k1 = curve == "K1";
    auto data_str = sig.substr(next_pivot + 1);
    eosio_assert(!data_str.empty(), "Signature has no data");
    vector<unsigned char> vch;

    eosio_assert(decode_base58(data_str, vch), "Decode signature failed");

    eosio_assert(vch.size() == 69, "Invalid signature");

    if (checksumming) {
        array<unsigned char, 67> check_data;
        copy_n(vch.begin(), 65, check_data.begin());
        check_data[65] = k1 ? 'K' : 'R';
        check_data[66] = '1';

        capi_checksum160 check_sig;
        ripemd160(reinterpret_cast<char*>(check_data.data()), 67, &check_sig);

        eosio_assert(memcmp(&check_sig.hash, &vch.end()[-4], 4) == 0, "Signature checksum mismatch");
    }

    capi_signature _sig;
    unsigned int type = k1 ? 0 : 1;
    _sig.data[0] = (uint8_t)type;
    for (int i = 1; i < sizeof(_sig.data); i++) {
        _sig.data[i] = vch[i - 1];
    }
    return _sig;
}

capi_public_key str_to_pub(const string& pubkey, const bool& checksumming = true) {
    string pubkey_prefix("EOS");
    auto base58substr = pubkey.substr(pubkey_prefix.length());
    vector<unsigned char> vch;
    eosio_assert(decode_base58(base58substr, vch), "Decode public key failed");
    eosio_assert(vch.size() == 37, "Invalid public key");
    if (checksumming) {

        array<unsigned char, 33> pubkey_data;
        copy_n(vch.begin(), 33, pubkey_data.begin());

        capi_checksum160 check_pubkey;
        ripemd160(reinterpret_cast<char*>(pubkey_data.data()), 33, &check_pubkey);

        eosio_assert(memcmp(&check_pubkey, &vch.end()[-4], 4) == 0, "Public key checksum mismatch");
    }
    capi_public_key _pub_key;
    unsigned int type = 0;
    _pub_key.data[0] = (char)type;
    for (int i = 1; i < sizeof(_pub_key.data); i++) {
        _pub_key.data[i] = vch[i - 1];
    }
    return _pub_key;
}

namespace dicecasino {
    CONTRACT dice : public contract {
        public:
          using contract::contract;
            dice ( name receiver, name code, eosio::datastream<const char*> ds ) : contract(receiver,code,ds),
                bets(receiver, receiver.value),
                hashs(receiver, receiver.value),
                results(receiver, receiver.value),
                configs(receiver, receiver.value){};

            // @abi action
            [[eosio::action]]
            void init();

            // @abi action
            [[eosio::action]]
            void closegame();

            // @abi action
            [[eosio::action]]
            void opengame();

            // @abi action
            [[eosio::action]]
            void clearconfig();

            // @abi action
            [[eosio::action]]
            void clearresult();

            // @abi action
            [[eosio::action]]
            void res(const capi_checksum160& user_seed, const capi_checksum256& seed);

            // @abi action
            [[eosio::action]]
            void transfer(name from, name to, asset quantity, string memo);

            // @abi action
            [[eosio::action]]
            void reveal(const uint64_t& id, const capi_checksum256& seed);

        private:
            struct [[eosio::table]] bet {
                uint64_t id;
                name player;
                name referrer;
                asset amount;
                uint8_t roll_under;
                capi_checksum256 seed_hash;
                capi_checksum160 user_seed_hash;
                uint64_t created_at;
                uint64_t primary_key() const { return id; }

                EOSLIB_SERIALIZE( bet, (id)(player)(referrer)(amount)(roll_under)(seed_hash)(user_seed_hash)(created_at))
            };

            struct [[eosio::table]] result {
                uint64_t bet_id;
                name player;
                name referrer;
                asset amount;
                uint8_t roll_under;
                uint8_t random_roll;
                capi_checksum256 seed;
                capi_checksum256 seed_hash;
                capi_checksum160 user_seed_hash;
                asset payout;
                uint64_t primary_key() const { return bet_id; }

                EOSLIB_SERIALIZE( result, (bet_id)(player)(referrer)(amount)(roll_under)(random_roll)(seed)(seed_hash)(user_seed_hash)(payout))
            };

            // @abi table hash i64
            struct [[eosio::table]] sthash {
                capi_checksum256 hash;
                uint64_t expiration;
                uint64_t primary_key() const { return uint64_hash(hash); }

                uint64_t by_expiration() const { return expiration; }

                EOSLIB_SERIALIZE( sthash, (hash)(expiration))
            };

            struct [[eosio::table]] config {
                uint8_t   configid;
                asset     locked;
                uint64_t  currentid;
                uint8_t   gamestatus;
                uint64_t  totalbet;
                uint64_t  totalwin;
                uint64_t  minbetamount;
                uint64_t  maxbetamount;

                uint64_t primary_key() const { return configid; }
                EOSLIB_SERIALIZE(config, (configid)(locked)(currentid)(gamestatus)(totalbet)(totalwin)(minbetamount)(maxbetamount))
            };

            typedef multi_index<name("result"), result> result_index;
            typedef multi_index<name("bet"), bet> bet_index;
            typedef multi_index<name("config"), config> config_index;
            typedef multi_index<"sthash"_n,sthash,indexed_by<"byexpiration"_n,const_mem_fun<sthash, uint64_t, &sthash::by_expiration>>> hash_index;

            result_index results;
            bet_index bets;
            config_index configs;
            hash_index hashs;

         // 参数解析
        void parse_memo(string memo,
                        uint8_t* roll_under,
                        capi_checksum256* seed_hash,
                        capi_checksum160* user_seed_hash,
                        uint64_t* expiration,
                        name* referrer,
                        capi_signature* sig) {
            // remove space
            memo.erase(std::remove_if(memo.begin(),
                                      memo.end(),
                                      [](unsigned char x) { return std::isspace(x); }),
                       memo.end());
            size_t sep_count = std::count(memo.begin(), memo.end(), '-');
            eosio_assert(sep_count == 5, "invalid memo");
            size_t pos;
            string container;
            pos = sub2sep(memo, &container, '-', 0, true);
            eosio_assert(!container.empty(), "no roll under");
            *roll_under = stoi(container);
            pos = sub2sep(memo, &container, '-', ++pos, true);
            eosio_assert(!container.empty(), "no seed hash");
            *seed_hash = hex_to_sha256(container);
            pos = sub2sep(memo, &container, '-', ++pos, true);
            eosio_assert(!container.empty(), "no user seed hash");
            *user_seed_hash = hex_to_sha1(container);
            pos = sub2sep(memo, &container, '-', ++pos, true);
            eosio_assert(!container.empty(), "no expiration");
            *expiration = stoull(container);
            pos = sub2sep(memo, &container, '-', ++pos, true);
            eosio_assert(!container.empty(), "no referrer");
            *referrer = name(container.c_str());
            container = memo.substr(++pos);
            eosio_assert(!container.empty(), "no signature");
            *sig = str_to_sig(container);
        }
        bet find_or_error(const uint64_t& id) {
            auto itr = bets.find(id);
            eosio_assert(itr != bets.end(), "bet not found");
            return *itr;
        }
        void save_bet(const bet& betobj) {
            bets.emplace(_self, [&](bet& b) {
                b.id = betobj.id;
                b.player = betobj.player;
                b.referrer = betobj.referrer;
                b.amount = betobj.amount;
                b.roll_under = betobj.roll_under;
                b.seed_hash = betobj.seed_hash;
                b.user_seed_hash = betobj.user_seed_hash;
                b.created_at = betobj.created_at;
            });
        }
        void remove_bet(const bet& bet) { bets.erase(bet); }
        void save_result(const result& resultobj) {
            results.emplace(_self, [&](result& r) {
                r.bet_id = resultobj.bet_id;
                r.player = resultobj.player;
                r.referrer = resultobj.referrer;
                r.amount = resultobj.amount;
                r.roll_under = resultobj.roll_under;
                r.random_roll = resultobj.random_roll;
                r.seed = resultobj.seed;
                r.seed_hash = resultobj.seed_hash;
                r.user_seed_hash = resultobj.user_seed_hash;
            });
        }
        // 迭代id
        uint64_t next_id() {
            auto configLookup = configs.find(1); //id is always 1
            configs.modify(configLookup, _self, [&](auto& newConfig) {
                newConfig.currentid += 1;
            });
            return configLookup->currentid + 1;
        }
        void lock(const asset& amount) {
            auto configLookup = configs.find(1); //id is always 1
            configs.modify(configLookup, _self, [&](auto& newConfig) {
                newConfig.locked += amount;
            });
        }
        void unlock(const asset& amount) {
            auto configLookup = configs.find(1); //id is always 1
            configs.modify(configLookup, _self, [&](auto& newConfig) {
                newConfig.locked -= amount;
            });
        }
        void add_total_bet(const asset& quantity) {
            auto configLookup = configs.find(1); //id is always 1
            configs.modify(configLookup, _self, [&](auto& newConfig) {
               newConfig.totalbet += quantity.amount;
            });
        }
        void add_total_win(const asset& quantity) {
            auto configLookup = configs.find(1); //id is always 1
            configs.modify(configLookup, _self, [&](auto& newConfig) {
               newConfig.totalwin += quantity.amount;
            });
        }
        // 拼接邀请人memo
        string referrer_memo(const bet& bet) {
            string memo = "bet id:";
            string id = uint64_string(bet.id);
            memo.append(id);
            memo.append(" player: ");
            string player = name{bet.player}.to_string();
            memo.append(player);
            memo.append(" referral reward! - unicorn.bi/eos/dice");
            return memo;
        }
        // 拼接胜利memo
        string winner_memo(const bet& bet) {
            string memo = "bet id:";
            string id = uint64_string(bet.id);
            memo.append(id);
            memo.append(" player: ");
            string player = name{bet.player}.to_string();
            memo.append(player);
            memo.append(" winner! - unicorn.bi/eos/dice");
            return memo;
        }
        // 根据两个提供的种子计算幸运数字
        uint8_t compute_random_roll(const capi_checksum256& seed1, const capi_checksum160& seed2) {
            string mixed_seed = sha256_to_hex(seed1);
            mixed_seed += sha1_to_hex(seed2);
            const char* mixed_seed_cstr = mixed_seed.c_str();
            capi_checksum256 source;
            sha256(mixed_seed_cstr, strlen(mixed_seed_cstr), &source);
            return (source.hash[0] + source.hash[1] + source.hash[2] + source.hash[3] + source.hash[4] + source.hash[5]) % 100 + 1;
        }
        // 计算邀请人的奖励
        asset compute_referrer_reward(const bet& bet) { return bet.amount / 200; }
        asset compute_payout(const uint8_t& roll_under, const asset& offer) {
            return min(max_payout(roll_under, offer), max_bonus());
        }
        asset max_payout(const uint8_t& roll_under, const asset& offer) {
            const double ODDS = 98.0 / ((double)roll_under - 1.0);
            return asset(ODDS * offer.amount, offer.symbol);
        }
        asset max_bonus() { return available_balance(); }
        asset available_balance() {
            auto currentconfig = configs.get(1);
            const asset total = asset(double(5000000), EOS_SYMBOL);
            const asset locked = currentconfig.locked;
            const asset available = total - locked;
            eosio_assert(available.amount >= 0, "fund pool overdraw");
            return available;
        }
        // 判断游戏是否进行
        void assert_game_status() {
            auto currentconfig = configs.get(1);
            eosio_assert(currentconfig.gamestatus == 0, "game had paused, please wait open!");
        }
        // 判断hash是否存在和过期
        void assert_hash(const capi_checksum256& seed_hash, const uint64_t& expiration) {
                const uint32_t _now = now();
                // check expiratin
                eosio_assert(expiration > _now, "seed hash expired");
                // check hash duplicate
                const uint64_t key = uint64_hash(seed_hash);
                auto itr = hashs.find(key);
                eosio_assert(itr == hashs.end(), "hash duplicate");
                // clean up
                auto index = hashs.get_index<name("byexpiration")>();
                auto upper_itr = index.upper_bound(_now);
                auto begin_itr = index.begin();
                while (begin_itr != upper_itr) {
                    begin_itr = index.erase(begin_itr);
                }
                // save hash
                hashs.emplace(_self, [&](sthash& r) {
                    r.hash = seed_hash;
                    r.expiration = expiration;
                });
            }
        // 判断投注是否正确
        void assert_quantity(const asset& quantity) {
            auto currentconfig = configs.get(1);
            eosio_assert(quantity.symbol == EOS_SYMBOL, "only EOS token allowed");
            eosio_assert(quantity.is_valid(), "quantity invalid");
            eosio_assert(quantity.amount >= currentconfig.minbetamount, "transfer quantity must be greater than 0.1");
            eosio_assert(quantity.amount <= currentconfig.maxbetamount, "transfer quantity must be less than 50");
        }
        // 判断投注数字是否正确
        void assert_roll_under(const uint8_t& roll_under, const asset& quantity) {
            eosio_assert(roll_under >= 3 && roll_under <= 96,
                         "roll under overflow, must be greater than 3 and less than 96");
            eosio_assert(
                max_payout(roll_under, quantity) <= max_bonus(),
                "offered overflow, expected earning is greater than the maximum bonus");
        }
        // 判断签名是否正确
        void assert_signature(const uint8_t& roll_under,
                              const capi_checksum256& seed_hash,
                              const uint64_t& expiration,
                              const name& referrer,
                              const capi_signature& sig) {
            string data = uint64_string(roll_under);
            data += "-";
            data += sha256_to_hex(seed_hash);
            data += "-";
            data += uint64_string(expiration);
            data += "-";
            data += name{referrer}.to_string();
            capi_checksum256 digest;
            const char* data_cstr = data.c_str();
            sha256(data_cstr, strlen(data_cstr), &digest);
            capi_public_key key = str_to_pub(PUB_KEY);
            assert_recover_key(&digest, (char*)&sig, sizeof(sig), (char*)&key, sizeof(key));
        }
        template <typename... Args>
        void send_defer_action(Args&&... args) {
            transaction trx;
            trx.actions.emplace_back(std::forward<Args>(args)...);
            trx.send(next_id(), _self, false);
        }
    };
}