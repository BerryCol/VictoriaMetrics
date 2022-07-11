// Code generated by qtc from "query_range_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/prometheus/query_range_response.qtpl:1
package prometheus

//line app/vmselect/prometheus/query_range_response.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmselect/netstorage"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/querytracer"
)

// QueryRangeResponse generates response for /api/v1/query_range.See https://prometheus.io/docs/prometheus/latest/querying/api/#range-queries

//line app/vmselect/prometheus/query_range_response.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/prometheus/query_range_response.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/prometheus/query_range_response.qtpl:9
func StreamQueryRangeResponse(qw422016 *qt422016.Writer, rs []netstorage.Result, qt *querytracer.Tracer, qtDone func()) {
//line app/vmselect/prometheus/query_range_response.qtpl:9
	qw422016.N().S(`{`)
//line app/vmselect/prometheus/query_range_response.qtpl:12
	seriesCount := len(rs)
	pointsCount := 0

//line app/vmselect/prometheus/query_range_response.qtpl:14
	qw422016.N().S(`"status":"success","data":{"resultType":"matrix","result":[`)
//line app/vmselect/prometheus/query_range_response.qtpl:19
	if len(rs) > 0 {
//line app/vmselect/prometheus/query_range_response.qtpl:20
		streamqueryRangeLine(qw422016, &rs[0])
//line app/vmselect/prometheus/query_range_response.qtpl:21
		pointsCount += len(rs[0].Values)

//line app/vmselect/prometheus/query_range_response.qtpl:22
		rs = rs[1:]

//line app/vmselect/prometheus/query_range_response.qtpl:23
		for i := range rs {
//line app/vmselect/prometheus/query_range_response.qtpl:23
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/query_range_response.qtpl:24
			streamqueryRangeLine(qw422016, &rs[i])
//line app/vmselect/prometheus/query_range_response.qtpl:25
			pointsCount += len(rs[i].Values)

//line app/vmselect/prometheus/query_range_response.qtpl:26
		}
//line app/vmselect/prometheus/query_range_response.qtpl:27
	}
//line app/vmselect/prometheus/query_range_response.qtpl:27
	qw422016.N().S(`]}`)
//line app/vmselect/prometheus/query_range_response.qtpl:31
	qt.Printf("generate /api/v1/query_range response for series=%d, points=%d", seriesCount, pointsCount)
	qtDone()

//line app/vmselect/prometheus/query_range_response.qtpl:34
	streamdumpQueryTrace(qw422016, qt)
//line app/vmselect/prometheus/query_range_response.qtpl:34
	qw422016.N().S(`}`)
//line app/vmselect/prometheus/query_range_response.qtpl:36
}

//line app/vmselect/prometheus/query_range_response.qtpl:36
func WriteQueryRangeResponse(qq422016 qtio422016.Writer, rs []netstorage.Result, qt *querytracer.Tracer, qtDone func()) {
//line app/vmselect/prometheus/query_range_response.qtpl:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/query_range_response.qtpl:36
	StreamQueryRangeResponse(qw422016, rs, qt, qtDone)
//line app/vmselect/prometheus/query_range_response.qtpl:36
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/query_range_response.qtpl:36
}

//line app/vmselect/prometheus/query_range_response.qtpl:36
func QueryRangeResponse(rs []netstorage.Result, qt *querytracer.Tracer, qtDone func()) string {
//line app/vmselect/prometheus/query_range_response.qtpl:36
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/query_range_response.qtpl:36
	WriteQueryRangeResponse(qb422016, rs, qt, qtDone)
//line app/vmselect/prometheus/query_range_response.qtpl:36
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/query_range_response.qtpl:36
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/query_range_response.qtpl:36
	return qs422016
//line app/vmselect/prometheus/query_range_response.qtpl:36
}

//line app/vmselect/prometheus/query_range_response.qtpl:38
func streamqueryRangeLine(qw422016 *qt422016.Writer, r *netstorage.Result) {
//line app/vmselect/prometheus/query_range_response.qtpl:38
	qw422016.N().S(`{"metric":`)
//line app/vmselect/prometheus/query_range_response.qtpl:40
	streammetricNameObject(qw422016, &r.MetricName)
//line app/vmselect/prometheus/query_range_response.qtpl:40
	qw422016.N().S(`,"values":`)
//line app/vmselect/prometheus/query_range_response.qtpl:41
	streamvaluesWithTimestamps(qw422016, r.Values, r.Timestamps)
//line app/vmselect/prometheus/query_range_response.qtpl:41
	qw422016.N().S(`}`)
//line app/vmselect/prometheus/query_range_response.qtpl:43
}

//line app/vmselect/prometheus/query_range_response.qtpl:43
func writequeryRangeLine(qq422016 qtio422016.Writer, r *netstorage.Result) {
//line app/vmselect/prometheus/query_range_response.qtpl:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/query_range_response.qtpl:43
	streamqueryRangeLine(qw422016, r)
//line app/vmselect/prometheus/query_range_response.qtpl:43
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/query_range_response.qtpl:43
}

//line app/vmselect/prometheus/query_range_response.qtpl:43
func queryRangeLine(r *netstorage.Result) string {
//line app/vmselect/prometheus/query_range_response.qtpl:43
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/query_range_response.qtpl:43
	writequeryRangeLine(qb422016, r)
//line app/vmselect/prometheus/query_range_response.qtpl:43
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/query_range_response.qtpl:43
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/query_range_response.qtpl:43
	return qs422016
//line app/vmselect/prometheus/query_range_response.qtpl:43
}
