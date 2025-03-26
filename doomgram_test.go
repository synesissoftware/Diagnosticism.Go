package diagnosticism_test

import (
	"github.com/stretchr/testify/require"
	. "github.com/synesissoftware/Diagnosticism.Go"

	"testing"
)

// NOTE: these functions taken primarily from **Diagnosticism.Rust**.

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

func TEST_doomgram_OVERFLOW_BY_SECONDS(t *testing.T) {

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

func TEST_doomgram_OVERFLOW_BY_MICROSECONDS(t *testing.T) {

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
