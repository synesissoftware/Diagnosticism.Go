package diagnosticism

import (
	stegol "github.com/synesissoftware/STEGoL"

	"testing"
)

// -------------------------------------------------------------------------
// Unit-tests
// -------------------------------------------------------------------------

func Test_calc_doom(t *testing.T) {

	{
		v := calc_doom(0)

		stegol.CheckIntegerEqual(t, 0, v)
	}

	{
		for i := uint64(1); i != 10; i++ {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 1, v)
		}
	}

	{
		for i := uint64(10); i != 100; i++ {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 2, v)
		}
	}

	{
		for i := uint64(100); i != 1_000; i++ {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 3, v)
		}
	}

	{
		for i := uint64(1_000); i != 10_000; i++ {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 4, v)
		}
	}

	{
		for i := uint64(10_000); i != 100_000; i++ {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 5, v)
		}
	}

	{
		for i := uint64(100_000); i != 1_000_000; i++ {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 6, v)
		}
	}

	{
		for i := uint64(1_000_000); i != 10_000_000; i += 1_000_000 {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 7, v)
		}
	}

	{
		for i := uint64(10_000_000); i != 100_000_000; i += 10_000_000 {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 8, v)
		}
	}

	{
		for i := uint64(100_000_000); i != 1_000_000_000; i += 100_000_000 {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 9, v)
		}
	}

	{
		for i := uint64(1_000_000_000); i != 10_000_000_000; i += 1_000_000_000 {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 10, v)
		}
	}

	{
		for i := uint64(10_000_000_000); i != 100_000_000_000; i += 10_000_000_000 {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 11, v)
		}
	}

	{
		for i := uint64(100_000_000_000); i != 1_000_000_000_000; i += 100_000_000_000 {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 12, v)
		}
	}

	{
		for i := uint64(1_000_000_000_000); i != 10_000_000_000_000; i += 1_000_000_000_000 {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 13, v)
		}
	}

	{
		for i := uint64(10_000_000_000_000); i != 100_000_000_000_000; i += 10_000_000_000_000 {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 14, v)
		}
	}

	{
		for i := uint64(100_000_000_000_000); i != 1_000_000_000_000_000; i += 100_000_000_000_000 {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 15, v)
		}
	}

	{
		for i := uint64(1_000_000_000_000_000); i != 10_000_000_000_000_000; i += 1_000_000_000_000_000 {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 16, v)
		}
	}

	{
		for i := uint64(10_000_000_000_000_000); i != 100_000_000_000_000_000; i += 10_000_000_000_000_000 {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 17, v)
		}
	}

	{
		for i := uint64(100_000_000_000_000_000); i != 1_000_000_000_000_000_000; i += 100_000_000_000_000_000 {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 18, v)
		}
	}

	{
		for i := uint64(1_000_000_000_000_000_000); i != 10_000_000_000_000_000_000; i += 1_000_000_000_000_000_000 {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 19, v)
		}
	}

	{
		for i := uint64(10_000_000_000_000_000_000); i != 18_000_000_000_000_000_000; i += 1_000_000_000_000_000_000 {
			v := calc_doom(i)

			stegol.CheckIntegerEqual(t, 20, v)
		}
	}
}

// -------------------------------------------------------------------------
// Benchmarks
// -------------------------------------------------------------------------

var (
	sink_uint32 uint32
)

func Benchmark_calc_doom_0(b *testing.B) {

	bmarks := []struct {
		name  string
		value uint64
	}{
		{
			name:  "zero",
			value: 0,
		},
		{
			name:  "1",
			value: 1,
		},
		{
			name:  "9",
			value: 9,
		},
		{
			name:  "10",
			value: 10,
		},
		{
			name:  "99",
			value: 99,
		},
		{
			name:  "100",
			value: 100,
		},
		{
			name:  "999",
			value: 999,
		},
		{
			name:  "1,000",
			value: 1_000,
		},
		{
			name:  "9,999",
			value: 9_999,
		},
		{
			name:  "10,000",
			value: 10_000,
		},
		{
			name:  "99,999",
			value: 99_999,
		},
		{
			name:  "100,000",
			value: 100_000,
		},
		{
			name:  "999,999",
			value: 999_999,
		},
		{
			name:  "1,000,000",
			value: 1_000_000,
		},
		{
			name:  "9,999,999",
			value: 9_999_999,
		},
		{
			name:  "10,000,000",
			value: 10_000_000,
		},
		{
			name:  "99,999,999",
			value: 99_999_999,
		},
		{
			name:  "100,000,000",
			value: 100_000_000,
		},
		{
			name:  "999,999,999",
			value: 999_999_999,
		},
		{
			name:  "1,000,000,000",
			value: 1_000_000_000,
		},
		{
			name:  "9,999,999,999",
			value: 9_999_999_999,
		},
		{
			name:  "10,000,000,000",
			value: 10_000_000_000,
		},
		{
			name:  "99,999,999,999",
			value: 99_999_999_999,
		},
		{
			name:  "100,000,000,000",
			value: 100_000_000_000,
		},
		{
			name:  "999,999,999,999",
			value: 999_999_999_999,
		},
		{
			name:  "1,000,000,000,000",
			value: 1_000_000_000_000,
		},
		{
			name:  "9,999,999,999,999",
			value: 9_999_999_999_999,
		},
		{
			name:  "10,000,000,000,000",
			value: 10_000_000_000_000,
		},
		{
			name:  "99,999,999,999,999",
			value: 99_999_999_999_999,
		},
		{
			name:  "100,000,000,000,000",
			value: 100_000_000_000_000,
		},
		{
			name:  "999,999,999,999,999",
			value: 999_999_999_999_999,
		},
		{
			name:  "1,000,000,000,000,000",
			value: 1_000_000_000_000_000,
		},
		{
			name:  "9,999,999,999,999,999",
			value: 9_999_999_999_999_999,
		},
		{
			name:  "10,000,000,000,000,000",
			value: 10_000_000_000_000_000,
		},
		{
			name:  "99,999,999,999,999,999",
			value: 99_999_999_999_999_999,
		},
		{
			name:  "100,000,000,000,000,000",
			value: 100_000_000_000_000_000,
		},
		{
			name:  "999,999,999,999,999,999",
			value: 999_999_999_999_999_999,
		},
		{
			name:  "1,000,000,000,000,000,000",
			value: 1_000_000_000_000_000_000,
		},
		{
			name:  "9,999,999,999,999,999,999",
			value: 9_999_999_999_999_999_999,
		},
		{
			name:  "10,000,000,000,000,000,000",
			value: 10_000_000_000_000_000_000,
		},
		{
			name:  "18,446,744,073,709,551,615",
			value: 18_446_744_073_709_551_615,
		},
	}

	for _, bb := range bmarks {
		b.Run(bb.name, func(b *testing.B) {

			b.ResetTimer()

			for range b.N {
				sink_uint32 = calc_doom(bb.value)
			}
		})
	}
}
