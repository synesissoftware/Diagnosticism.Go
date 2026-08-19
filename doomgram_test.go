package diagnosticism_test

import (
	. "github.com/synesissoftware/Diagnosticism.Go"

	"github.com/stretchr/testify/require"

	"testing"
	"time"
)

// NOTE: these functions taken primarily from **Diagnosticism.Rust**.

// -------------------------------------------------------------------------
// Fixtures
// -------------------------------------------------------------------------

// Returns a zero-value DOOMGram with no events recorded.
func emptyDOOMGram() DOOMGram {
	return DOOMGram{}
}

// Returns a DOOMGram heavily loaded in a single bucket, exercising the
// overflow/scaling path in gram_doom_to_char.
func skewedDOOMGram() DOOMGram {
	var dg DOOMGram

	for i := 0; i != 1_000_000; i++ {
		dg.PushEventDuration(1 * time.Millisecond)
	}
	dg.PushEventDuration(1 * time.Nanosecond)

	return dg
}

// Returns a DOOMGram with events only in the microsecond buckets — a
// realistic profile for a fast in-process operation.
func sparseDOOMGram() DOOMGram {
	var dg DOOMGram

	for i := 0; i != 1_000; i++ {
		dg.PushEventDuration(10 * time.Microsecond)
	}
	for i := 0; i != 50; i++ {
		dg.PushEventDuration(100 * time.Microsecond)
	}

	return dg
}

// Returns a DOOMGram with one event recorded in every bucket, giving a
// representative "all buckets active" strip.
func uniformDOOMGram() DOOMGram {
	var dg DOOMGram

	// One sample in each of the 12 OOM buckets.
	durations := []time.Duration{
		1 * time.Nanosecond,
		10 * time.Nanosecond,
		100 * time.Nanosecond,
		1 * time.Microsecond,
		10 * time.Microsecond,
		100 * time.Microsecond,
		1 * time.Millisecond,
		10 * time.Millisecond,
		100 * time.Millisecond,
		1 * time.Second,
		10 * time.Second,
		100 * time.Second,
	}
	for _, dur := range durations {
		dg.PushEventDuration(dur) // NOTE: assumes PushEventDuration(time.Duration) is the recording API.
	}

	return dg
}

// -------------------------------------------------------------------------
// Unit-tests
// -------------------------------------------------------------------------

func Test_DOOMGram_Default(t *testing.T) {

	var dg DOOMGram

	require.Equal(t, uint64(0), dg.EventCount())

	require.Equal(t, uint64(0), dg.EventTimeTotalRaw())

	{
		ok, total_time := dg.EventTimeTotal()

		require.True(t, ok)
		require.Equal(t, uint64(0), total_time)
	}

	require.False(t, dg.Overflowed())

	{
		ok, _ := dg.MinEventTime()

		require.False(t, ok)
	}

	{
		ok, _ := dg.MaxEventTime()

		require.False(t, ok)
	}

	require.Equal(t, uint64(0), dg.NumEventsIn1ns())
	require.Equal(t, uint64(0), dg.NumEventsIn10ns())
	require.Equal(t, uint64(0), dg.NumEventsIn100ns())
	require.Equal(t, uint64(0), dg.NumEventsIn1us())
	require.Equal(t, uint64(0), dg.NumEventsIn10us())
	require.Equal(t, uint64(0), dg.NumEventsIn100us())
	require.Equal(t, uint64(0), dg.NumEventsIn1ms())
	require.Equal(t, uint64(0), dg.NumEventsIn10ms())
	require.Equal(t, uint64(0), dg.NumEventsIn100ms())
	require.Equal(t, uint64(0), dg.NumEventsIn1s())
	require.Equal(t, uint64(0), dg.NumEventsIn10s())
	require.Equal(t, uint64(0), dg.NumEventsIe100s())

	require.Equal(t, "____________", dg.ToStrip())
}

func Test_DOOMGram_SINGLE_TIMING_EVENT(t *testing.T) {

	var dg DOOMGram

	dg.PushEventTimeMs(13)

	require.Equal(t, uint64(1), dg.EventCount())

	require.Equal(t, uint64(13000000), dg.EventTimeTotalRaw())

	{
		ok, total_time := dg.EventTimeTotal()

		require.True(t, ok)
		require.Equal(t, uint64(13000000), total_time)
	}

	require.False(t, dg.Overflowed())

	{
		ok, tm := dg.MinEventTime()

		require.True(t, ok)
		require.Equal(t, tm, uint64(13000000))
	}

	{
		ok, tm := dg.MaxEventTime()

		require.True(t, ok)
		require.Equal(t, tm, uint64(13000000))
	}

	require.Equal(t, uint64(0), dg.NumEventsIn1ns())
	require.Equal(t, uint64(0), dg.NumEventsIn10ns())
	require.Equal(t, uint64(0), dg.NumEventsIn100ns())
	require.Equal(t, uint64(0), dg.NumEventsIn1us())
	require.Equal(t, uint64(0), dg.NumEventsIn10us())
	require.Equal(t, uint64(0), dg.NumEventsIn100us())
	require.Equal(t, uint64(0), dg.NumEventsIn1ms())
	require.Equal(t, uint64(1), dg.NumEventsIn10ms())
	require.Equal(t, uint64(0), dg.NumEventsIn100ms())
	require.Equal(t, uint64(0), dg.NumEventsIn1s())
	require.Equal(t, uint64(0), dg.NumEventsIn10s())
	require.Equal(t, uint64(0), dg.NumEventsIe100s())

	require.Equal(t, "_______a____", dg.ToStrip())
}

func Test_DOOMGram_UNIFORM_SPREAD_TIMINGS_1(t *testing.T) {

	var dg DOOMGram

	dg.PushEventTimeNs(9)
	dg.PushEventTimeNs(80)
	dg.PushEventTimeNs(700)
	dg.PushEventTimeUs(6)
	dg.PushEventTimeUs(50)
	dg.PushEventTimeUs(400)
	dg.PushEventTimeMs(3)
	dg.PushEventTimeMs(20)
	dg.PushEventTimeMs(100)
	dg.PushEventTimeS(9)
	dg.PushEventTimeS(80)
	dg.PushEventTimeS(700)

	require.Equal(t, uint64(12), dg.EventCount())

	require.Equal(t, uint64(789123456789), dg.EventTimeTotalRaw())

	{
		ok, total_time := dg.EventTimeTotal()

		require.True(t, ok)
		require.Equal(t, uint64(789123456789), total_time)
	}

	require.False(t, dg.Overflowed())

	{
		ok, tm := dg.MinEventTime()

		require.True(t, ok)
		require.Equal(t, tm, uint64(9))
	}

	{
		ok, tm := dg.MaxEventTime()

		require.True(t, ok)
		require.Equal(t, tm, uint64(700000000000))
	}

	require.Equal(t, uint64(1), dg.NumEventsIn1ns())
	require.Equal(t, uint64(1), dg.NumEventsIn10ns())
	require.Equal(t, uint64(1), dg.NumEventsIn100ns())
	require.Equal(t, uint64(1), dg.NumEventsIn1us())
	require.Equal(t, uint64(1), dg.NumEventsIn10us())
	require.Equal(t, uint64(1), dg.NumEventsIn100us())
	require.Equal(t, uint64(1), dg.NumEventsIn1ms())
	require.Equal(t, uint64(1), dg.NumEventsIn10ms())
	require.Equal(t, uint64(1), dg.NumEventsIn100ms())
	require.Equal(t, uint64(1), dg.NumEventsIn1s())
	require.Equal(t, uint64(1), dg.NumEventsIn10s())
	require.Equal(t, uint64(1), dg.NumEventsIe100s())

	require.Equal(t, "aaaaaaaaaaaa", dg.ToStrip())
}

func Test_DOOMGram_UNIFORM_SPREAD_TIMINGS_2(t *testing.T) {

	var dg DOOMGram

	dg.PushEventTimeNs(9)
	dg.PushEventTimeNs(80)
	dg.PushEventTimeNs(700)
	dg.PushEventTimeNs(6000)
	dg.PushEventTimeNs(50000)
	dg.PushEventTimeNs(400000)
	dg.PushEventTimeMs(3)
	dg.PushEventTimeMs(20)
	dg.PushEventTimeMs(100)
	dg.PushEventTimeMs(9000)
	dg.PushEventTimeMs(80000)
	dg.PushEventTimeMs(700000)

	require.Equal(t, uint64(12), dg.EventCount())

	require.Equal(t, uint64(789123456789), dg.EventTimeTotalRaw())

	{
		ok, total_time := dg.EventTimeTotal()

		require.True(t, ok)
		require.Equal(t, uint64(789123456789), total_time)
	}

	require.False(t, dg.Overflowed())

	{
		ok, tm := dg.MinEventTime()

		require.True(t, ok)
		require.Equal(t, tm, uint64(9))
	}

	{
		ok, tm := dg.MaxEventTime()

		require.True(t, ok)
		require.Equal(t, tm, uint64(700000000000))
	}

	require.Equal(t, uint64(1), dg.NumEventsIn1ns())
	require.Equal(t, uint64(1), dg.NumEventsIn10ns())
	require.Equal(t, uint64(1), dg.NumEventsIn100ns())
	require.Equal(t, uint64(1), dg.NumEventsIn1us())
	require.Equal(t, uint64(1), dg.NumEventsIn10us())
	require.Equal(t, uint64(1), dg.NumEventsIn100us())
	require.Equal(t, uint64(1), dg.NumEventsIn1ms())
	require.Equal(t, uint64(1), dg.NumEventsIn10ms())
	require.Equal(t, uint64(1), dg.NumEventsIn100ms())
	require.Equal(t, uint64(1), dg.NumEventsIn1s())
	require.Equal(t, uint64(1), dg.NumEventsIn10s())
	require.Equal(t, uint64(1), dg.NumEventsIe100s())

	require.Equal(t, "aaaaaaaaaaaa", dg.ToStrip())
}

func Test_DOOMGram_UNIFORM_SPREAD_TIMINGS_3(t *testing.T) {

	var dg DOOMGram

	dg.PushEventTimeNs(9)
	dg.PushEventTimeNs(80)
	dg.PushEventTimeNs(700)
	dg.PushEventTimeNs(6000)
	dg.PushEventTimeNs(50000)
	dg.PushEventTimeNs(400000)
	dg.PushEventTimeNs(3000000)
	dg.PushEventTimeNs(20000000)
	dg.PushEventTimeNs(100000000)
	dg.PushEventTimeNs(9000000000)
	dg.PushEventTimeNs(80000000000)
	dg.PushEventTimeNs(700000000000)

	require.Equal(t, uint64(12), dg.EventCount())

	require.Equal(t, uint64(789123456789), dg.EventTimeTotalRaw())

	{
		ok, total_time := dg.EventTimeTotal()

		require.True(t, ok)
		require.Equal(t, uint64(789123456789), total_time)
	}

	require.False(t, dg.Overflowed())

	{
		ok, tm := dg.MinEventTime()

		require.True(t, ok)
		require.Equal(t, tm, uint64(9))
	}

	{
		ok, tm := dg.MaxEventTime()

		require.True(t, ok)
		require.Equal(t, tm, uint64(700000000000))
	}

	require.Equal(t, uint64(1), dg.NumEventsIn1ns())
	require.Equal(t, uint64(1), dg.NumEventsIn10ns())
	require.Equal(t, uint64(1), dg.NumEventsIn100ns())
	require.Equal(t, uint64(1), dg.NumEventsIn1us())
	require.Equal(t, uint64(1), dg.NumEventsIn10us())
	require.Equal(t, uint64(1), dg.NumEventsIn100us())
	require.Equal(t, uint64(1), dg.NumEventsIn1ms())
	require.Equal(t, uint64(1), dg.NumEventsIn10ms())
	require.Equal(t, uint64(1), dg.NumEventsIn100ms())
	require.Equal(t, uint64(1), dg.NumEventsIn1s())
	require.Equal(t, uint64(1), dg.NumEventsIn10s())
	require.Equal(t, uint64(1), dg.NumEventsIe100s())

	require.Equal(t, "aaaaaaaaaaaa", dg.ToStrip())
}

func Test_DOOMGram_UNIFORM_SPREAD_TIMINGS_4(t *testing.T) {

	var dg DOOMGram

	dg.PushEventTimeUs(6)
	dg.PushEventTimeUs(50)
	dg.PushEventTimeUs(400)
	dg.PushEventTimeUs(3000)
	dg.PushEventTimeUs(20000)
	dg.PushEventTimeUs(100000)
	dg.PushEventTimeUs(9000000)
	dg.PushEventTimeUs(80000000)
	dg.PushEventTimeUs(700000000)

	require.Equal(t, uint64(9), dg.EventCount())

	require.Equal(t, uint64(789123456000), dg.EventTimeTotalRaw())

	{
		ok, total_time := dg.EventTimeTotal()

		require.True(t, ok)
		require.Equal(t, uint64(789123456000), total_time)
	}

	require.False(t, dg.Overflowed())

	{
		ok, tm := dg.MinEventTime()

		require.True(t, ok)
		require.Equal(t, tm, uint64(6000))
	}

	{
		ok, tm := dg.MaxEventTime()

		require.True(t, ok)
		require.Equal(t, tm, uint64(700000000000))
	}

	require.Equal(t, uint64(0), dg.NumEventsIn1ns())
	require.Equal(t, uint64(0), dg.NumEventsIn10ns())
	require.Equal(t, uint64(0), dg.NumEventsIn100ns())
	require.Equal(t, uint64(1), dg.NumEventsIn1us())
	require.Equal(t, uint64(1), dg.NumEventsIn10us())
	require.Equal(t, uint64(1), dg.NumEventsIn100us())
	require.Equal(t, uint64(1), dg.NumEventsIn1ms())
	require.Equal(t, uint64(1), dg.NumEventsIn10ms())
	require.Equal(t, uint64(1), dg.NumEventsIn100ms())
	require.Equal(t, uint64(1), dg.NumEventsIn1s())
	require.Equal(t, uint64(1), dg.NumEventsIn10s())
	require.Equal(t, uint64(1), dg.NumEventsIe100s())

	require.Equal(t, "___aaaaaaaaa", dg.ToStrip())
}

func Test_DOOMGram_SEVERAL_DISTINCT_TIMINGS(t *testing.T) {

	var dg DOOMGram

	dg.PushEventTimeNs(23)
	dg.PushEventTimeNs(10)
	dg.PushEventTimeUs(7)
	dg.PushEventTimeUs(7)
	dg.PushEventTimeUs(89)
	dg.PushEventTimeMs(248)
	dg.PushEventTimeS(5)
	dg.PushEventTimeS(309)

	require.Equal(t, uint64(8), dg.EventCount())

	require.Equal(t, uint64(314248103033), dg.EventTimeTotalRaw())

	{
		ok, total_time := dg.EventTimeTotal()

		require.True(t, ok)
		require.Equal(t, uint64(314248103033), total_time)
	}

	require.False(t, dg.Overflowed())

	{
		ok, tm := dg.MinEventTime()

		require.True(t, ok)
		require.Equal(t, tm, uint64(10))
	}

	{
		ok, tm := dg.MaxEventTime()

		require.True(t, ok)
		require.Equal(t, tm, uint64(309000000000))
	}

	require.Equal(t, uint64(0), dg.NumEventsIn1ns())
	require.Equal(t, uint64(2), dg.NumEventsIn10ns())
	require.Equal(t, uint64(0), dg.NumEventsIn100ns())
	require.Equal(t, uint64(2), dg.NumEventsIn1us())
	require.Equal(t, uint64(1), dg.NumEventsIn10us())
	require.Equal(t, uint64(0), dg.NumEventsIn100us())
	require.Equal(t, uint64(0), dg.NumEventsIn1ms())
	require.Equal(t, uint64(0), dg.NumEventsIn10ms())
	require.Equal(t, uint64(1), dg.NumEventsIn100ms())
	require.Equal(t, uint64(1), dg.NumEventsIn1s())
	require.Equal(t, uint64(0), dg.NumEventsIn10s())
	require.Equal(t, uint64(1), dg.NumEventsIe100s())

	require.Equal(t, "_a_aa___aa_a", dg.ToStrip())
}

func Test_DOOMGram_SEVERAL_INTERSECTING_TIMINGS(t *testing.T) {

	var dg DOOMGram

	dg.PushEventTimeNs(11)
	dg.PushEventTimeNs(19)
	dg.PushEventTimeNs(19)
	dg.PushEventTimeUs(7)
	dg.PushEventTimeUs(7)
	dg.PushEventTimeUs(89)
	dg.PushEventTimeMs(248)
	dg.PushEventTimeMs(4321)
	dg.PushEventTimeS(5)
	dg.PushEventTimeS(309)

	require.Equal(t, uint64(10), dg.EventCount())

	require.Equal(t, uint64(318569103049), dg.EventTimeTotalRaw())

	{
		ok, total_time := dg.EventTimeTotal()

		require.True(t, ok)
		require.Equal(t, uint64(318569103049), total_time)
	}

	require.False(t, dg.Overflowed())

	{
		ok, tm := dg.MinEventTime()

		require.True(t, ok)
		require.Equal(t, tm, uint64(11))
	}

	{
		ok, tm := dg.MaxEventTime()

		require.True(t, ok)
		require.Equal(t, tm, uint64(309000000000))
	}

	require.Equal(t, uint64(0), dg.NumEventsIn1ns())
	require.Equal(t, uint64(3), dg.NumEventsIn10ns())
	require.Equal(t, uint64(0), dg.NumEventsIn100ns())
	require.Equal(t, uint64(2), dg.NumEventsIn1us())
	require.Equal(t, uint64(1), dg.NumEventsIn10us())
	require.Equal(t, uint64(0), dg.NumEventsIn100us())
	require.Equal(t, uint64(0), dg.NumEventsIn1ms())
	require.Equal(t, uint64(0), dg.NumEventsIn10ms())
	require.Equal(t, uint64(1), dg.NumEventsIn100ms())
	require.Equal(t, uint64(2), dg.NumEventsIn1s())
	require.Equal(t, uint64(0), dg.NumEventsIn10s())
	require.Equal(t, uint64(1), dg.NumEventsIe100s())

	require.Equal(t, "_a_aa___aa_a", dg.ToStrip())
}

func Test_DOOMGram_OVERFLOW_BY_SECONDS(t *testing.T) {

	var dg DOOMGram

	// u64 max:
	//
	// 18,446,744,073,709,551,615 ns
	//     18,446,744,073,709,551 µs
	//         18,446,744,073,709 ms
	//             18,446,744,073  s

	// 18446744073 max

	// add in max # seconds
	{
		ok := dg.PushEventTimeS(18446744073)

		require.True(t, ok)
	}

	// add in 0 more
	{
		ok := dg.PushEventTimeS(0)

		require.True(t, ok)
	}

	// add in 1 more to overflow
	{
		ok := dg.PushEventTimeS(1)

		require.False(t, ok)
	}
}

func Test_DOOMGram_OVERFLOW_BY_MICROSECONDS(t *testing.T) {

	var dg DOOMGram

	// u64 max:
	//
	// 18,446,744,073,709,551,615 ns
	//     18,446,744,073,709,551 µs
	//         18,446,744,073,709 ms
	//             18,446,744,073  s

	// 18446744073 max

	// add in max-1 # microseconds
	{
		ok := dg.PushEventTimeUs(18446744073709550)

		require.True(t, ok)
	}

	// add in 1 more to max
	{
		ok := dg.PushEventTimeUs(1)

		require.True(t, ok)
	}

	// add in 0 more
	{
		ok := dg.PushEventTimeUs(0)

		require.True(t, ok)
	}

	// add in 1 more to overflow
	{
		ok := dg.PushEventTimeUs(1)

		require.False(t, ok)
	}
}

// -------------------------------------------------------------------------
// Benchmarks
// -------------------------------------------------------------------------

var (
	sink_string string
)

// Measures the cost of recording a single sub-10ns event.
func Benchmark_ADD_EVENT_1ns(b *testing.B) {
	var dg DOOMGram

	b.ResetTimer()

	for range b.N {
		dg.PushEventDuration(1 * time.Nanosecond)
	}
}

// Measures the cost of recording a typical fast in-process duration.
func Benchmark_ADD_EVENT_1us(b *testing.B) {
	var dg DOOMGram

	b.ResetTimer()

	for range b.N {
		dg.PushEventDuration(1 * time.Microsecond)
	}
}

// Measures the cost of recording a slow (1s) event, exercising the far end
// of the bucket array.
func Benchmark_ADD_EVENT_1s(b *testing.B) {
	var dg DOOMGram

	b.ResetTimer()

	for range b.N {
		dg.PushEventDuration(1 * time.Second)
	}
}

// -----------------------------------------------------
// Benchmarks: ToStrip — the primary subject of interest
// -----------------------------------------------------

// Measures ToStrip on a zero DOOMGram.
//
// All 12 buckets are zero; exercises the "all underscore" fast path if any.
func Benchmark_ToStrip_EMPTY(b *testing.B) {
	dg := emptyDOOMGram()

	b.ResetTimer()

	for range b.N {
		sink_string = dg.ToStrip()
	}
}

// Measures ToStrip when all 12 buckets are equally populated. This is the
// most representative general case.
func Benchmark_ToStrip_UNIFORM(b *testing.B) {
	dg := uniformDOOMGram()

	b.ResetTimer()

	for range b.N {
		sink_string = dg.ToStrip()
	}
}

// Measures ToStrip with a realistic sparse profile (most buckets empty,
// activity concentrated in two adjacent buckets).
func Benchmark_ToStrip_SPARSE(b *testing.B) {
	dg := sparseDOOMGram()

	b.ResetTimer()

	for range b.N {
		sink_string = dg.ToStrip()
	}
}

// Measures ToStrip when one bucket has vastly more events than the others.
// Tests the scaling arithmetic under extreme ratios.
func Benchmark_ToStrip_SKEWED(b *testing.B) {
	dg := skewedDOOMGram()

	b.ResetTimer()

	for range b.N {
		sink_string = dg.ToStrip()
	}
}

// -----------------------------------------------------
// Benchmarks: combined Add + ToStrip round-trip
// -----------------------------------------------------

// Measures the cost of a realistic usage pattern: record N events, then
// format once. The ratio here (1000:1) is typical for a component that logs
// its histogram periodically.
func Benchmark_ADD_EVENTS_THEN_CALL_ToStrip(b *testing.B) {
	for range b.N {
		var dg DOOMGram

		for j := 0; j < 1_000; j++ {
			dg.PushEventDuration(time.Duration(j+1) * time.Microsecond)
		}

		sink_string = dg.ToStrip()
	}
}

// Measures the pathological case where ToStrip is called after every single
// PushEventDuration — the worst case for any non-optimised allocation.
func Benchmark_ADD_EVENT_AND_CALL_ToStrip_INTERLEAVED(b *testing.B) {
	var dg DOOMGram

	b.ResetTimer()

	for i := 0; i != b.N; i++ {
		dg.PushEventDuration(time.Duration(i+1) * time.Microsecond)

		sink_string = dg.ToStrip()
	}
}

// -----------------------------------------------------
// Allocation probe: assert expected alloc count for ToStrip
// -----------------------------------------------------

// Test_ToStrip_ALLOCATIONS is not a benchmark but a correctness check on
// the allocator behaviour of ToStrip. It documents the current allocation
// count so that any regression (e.g. an extra heap allocation introduced by
// a refactor) is caught immediately.
func Test_ToStrip_ALLOCATIONS(t *testing.T) {
	dg := uniformDOOMGram()

	allocs := testing.AllocsPerRun(1_000, func() {
		sink_string = dg.ToStrip()
	})

	const expectedAllocs = 1
	if allocs != expectedAllocs {
		t.Errorf("ToStrip(): got %.0f allocs/op, want %d — allocation profile has changed", allocs, expectedAllocs)
	}
}
