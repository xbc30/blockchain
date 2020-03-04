
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package librato

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"time"

	"github.com/ethereum/go-ethereum/metrics"
)

//用于从time.duration.string提取单位的regexp
var unitRegexp = regexp.MustCompile(`[^\\d]+$`)

//将时间.duration转换为库的帮助程序，用于显示计时器度量的属性
func translateTimerAttributes(d time.Duration) (attrs map[string]interface{}) {
	attrs = make(map[string]interface{})
	attrs[DisplayTransform] = fmt.Sprintf("x/%d", int64(d))
	attrs[DisplayUnitsShort] = string(unitRegexp.Find([]byte(d.String())))
	return
}

type Reporter struct {
	Email, Token    string
	Namespace       string
	Source          string
	Interval        time.Duration
	Registry        metrics.Registry
Percentiles     []float64              //要报告的柱状图指标百分比
TimerAttributes map[string]interface{} //显示计时器的单位
	intervalSec     int64
}

func NewReporter(r metrics.Registry, d time.Duration, e string, t string, s string, p []float64, u time.Duration) *Reporter {
	return &Reporter{e, t, "", s, d, r, p, translateTimerAttributes(u), int64(d / time.Second)}
}

func Librato(r metrics.Registry, d time.Duration, e string, t string, s string, p []float64, u time.Duration) {
	NewReporter(r, d, e, t, s, p, u).Run()
}

func (rep *Reporter) Run() {
log.Printf("WARNING: This client has been DEPRECATED! It has been moved to https://github.com/mihasya/go-metrics-libarto，将于2015年8月5日从rcrowley/go metrics中删除）
	ticker := time.Tick(rep.Interval)
	metricsApi := &LibratoClient{rep.Email, rep.Token}
	for now := range ticker {
		var metrics Batch
		var err error
		if metrics, err = rep.BuildRequest(now, rep.Registry); err != nil {
			log.Printf("ERROR constructing librato request body %s", err)
			continue
		}
		if err := metricsApi.PostMetrics(metrics); err != nil {
			log.Printf("ERROR sending metrics to librato %s", err)
			continue
		}
	}
}

//根据度量提供的数据计算平方和。柱状图
//参见http://en.wikipedia.org/wiki/standard_deviation rapid_calculation_methods
func sumSquares(s metrics.Sample) float64 {
	count := float64(s.Count())
	sumSquared := math.Pow(count*s.Mean(), 2)
	sumSquares := math.Pow(count*s.StdDev(), 2) + sumSquared/count
	if math.IsNaN(sumSquares) {
		return 0.0
	}
	return sumSquares
}
func sumSquaresTimer(t metrics.Timer) float64 {
	count := float64(t.Count())
	sumSquared := math.Pow(count*t.Mean(), 2)
	sumSquares := math.Pow(count*t.StdDev(), 2) + sumSquared/count
	if math.IsNaN(sumSquares) {
		return 0.0
	}
	return sumSquares
}

func (rep *Reporter) BuildRequest(now time.Time, r metrics.Registry) (snapshot Batch, err error) {
	snapshot = Batch{
//将时间戳强制为单步执行fn，以便它们在天秤座图中对齐
		MeasureTime: (now.Unix() / rep.intervalSec) * rep.intervalSec,
		Source:      rep.Source,
	}
	snapshot.Gauges = make([]Measurement, 0)
	snapshot.Counters = make([]Measurement, 0)
	histogramGaugeCount := 1 + len(rep.Percentiles)
	r.Each(func(name string, metric interface{}) {
		if rep.Namespace != "" {
			name = fmt.Sprintf("%s.%s", rep.Namespace, name)
		}
		measurement := Measurement{}
		measurement[Period] = rep.Interval.Seconds()
		switch m := metric.(type) {
		case metrics.Counter:
			if m.Count() > 0 {
				measurement[Name] = fmt.Sprintf("%s.%s", name, "count")
				measurement[Value] = float64(m.Count())
				measurement[Attributes] = map[string]interface{}{
					DisplayUnitsLong:  Operations,
					DisplayUnitsShort: OperationsShort,
					DisplayMin:        "0",
				}
				snapshot.Counters = append(snapshot.Counters, measurement)
			}
		case metrics.Gauge:
			measurement[Name] = name
			measurement[Value] = float64(m.Value())
			snapshot.Gauges = append(snapshot.Gauges, measurement)
		case metrics.GaugeFloat64:
			measurement[Name] = name
			measurement[Value] = m.Value()
			snapshot.Gauges = append(snapshot.Gauges, measurement)
		case metrics.Histogram:
			if m.Count() > 0 {
				gauges := make([]Measurement, histogramGaugeCount)
				s := m.Sample()
				measurement[Name] = fmt.Sprintf("%s.%s", name, "hist")
				measurement[Count] = uint64(s.Count())
				measurement[Max] = float64(s.Max())
				measurement[Min] = float64(s.Min())
				measurement[Sum] = float64(s.Sum())
				measurement[SumSquares] = sumSquares(s)
				gauges[0] = measurement
				for i, p := range rep.Percentiles {
					gauges[i+1] = Measurement{
						Name:   fmt.Sprintf("%s.%.2f", measurement[Name], p),
						Value:  s.Percentile(p),
						Period: measurement[Period],
					}
				}
				snapshot.Gauges = append(snapshot.Gauges, gauges...)
			}
		case metrics.Meter:
			measurement[Name] = name
			measurement[Value] = float64(m.Count())
			snapshot.Counters = append(snapshot.Counters, measurement)
			snapshot.Gauges = append(snapshot.Gauges,
				Measurement{
					Name:   fmt.Sprintf("%s.%s", name, "1min"),
					Value:  m.Rate1(),
					Period: int64(rep.Interval.Seconds()),
					Attributes: map[string]interface{}{
						DisplayUnitsLong:  Operations,
						DisplayUnitsShort: OperationsShort,
						DisplayMin:        "0",
					},
				},
				Measurement{
					Name:   fmt.Sprintf("%s.%s", name, "5min"),
					Value:  m.Rate5(),
					Period: int64(rep.Interval.Seconds()),
					Attributes: map[string]interface{}{
						DisplayUnitsLong:  Operations,
						DisplayUnitsShort: OperationsShort,
						DisplayMin:        "0",
					},
				},
				Measurement{
					Name:   fmt.Sprintf("%s.%s", name, "15min"),
					Value:  m.Rate15(),
					Period: int64(rep.Interval.Seconds()),
					Attributes: map[string]interface{}{
						DisplayUnitsLong:  Operations,
						DisplayUnitsShort: OperationsShort,
						DisplayMin:        "0",
					},
				},
			)
		case metrics.Timer:
			measurement[Name] = name
			measurement[Value] = float64(m.Count())
			snapshot.Counters = append(snapshot.Counters, measurement)
			if m.Count() > 0 {
				libratoName := fmt.Sprintf("%s.%s", name, "timer.mean")
				gauges := make([]Measurement, histogramGaugeCount)
				gauges[0] = Measurement{
					Name:       libratoName,
					Count:      uint64(m.Count()),
					Sum:        m.Mean() * float64(m.Count()),
					Max:        float64(m.Max()),
					Min:        float64(m.Min()),
					SumSquares: sumSquaresTimer(m),
					Period:     int64(rep.Interval.Seconds()),
					Attributes: rep.TimerAttributes,
				}
				for i, p := range rep.Percentiles {
					gauges[i+1] = Measurement{
						Name:       fmt.Sprintf("%s.timer.%2.0f", name, p*100),
						Value:      m.Percentile(p),
						Period:     int64(rep.Interval.Seconds()),
						Attributes: rep.TimerAttributes,
					}
				}
				snapshot.Gauges = append(snapshot.Gauges, gauges...)
				snapshot.Gauges = append(snapshot.Gauges,
					Measurement{
						Name:   fmt.Sprintf("%s.%s", name, "rate.1min"),
						Value:  m.Rate1(),
						Period: int64(rep.Interval.Seconds()),
						Attributes: map[string]interface{}{
							DisplayUnitsLong:  Operations,
							DisplayUnitsShort: OperationsShort,
							DisplayMin:        "0",
						},
					},
					Measurement{
						Name:   fmt.Sprintf("%s.%s", name, "rate.5min"),
						Value:  m.Rate5(),
						Period: int64(rep.Interval.Seconds()),
						Attributes: map[string]interface{}{
							DisplayUnitsLong:  Operations,
							DisplayUnitsShort: OperationsShort,
							DisplayMin:        "0",
						},
					},
					Measurement{
						Name:   fmt.Sprintf("%s.%s", name, "rate.15min"),
						Value:  m.Rate15(),
						Period: int64(rep.Interval.Seconds()),
						Attributes: map[string]interface{}{
							DisplayUnitsLong:  Operations,
							DisplayUnitsShort: OperationsShort,
							DisplayMin:        "0",
						},
					},
				)
			}
		}
	})
	return
}
