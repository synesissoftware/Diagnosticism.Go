package diagnosticism

// Decimal Order-Of-Magnitude frequency histoGRAM
//
// # Note:
// This is a Go port of the equivalent `stlsoft::doomgram` class from the
// **STLSoft** libraries (https://github.com/synesissoftware/STLSoft-1.11).
type DOOMGram struct {
	event_count         uint64
	event_time_total    uint64
	has_overflowed      bool
	min_event_time      uint64
	max_event_time      uint64
	num_events_in_1ns   uint64
	num_events_in_10ns  uint64
	num_events_in_100ns uint64
	num_events_in_1us   uint64
	num_events_in_10us  uint64
	num_events_in_100us uint64
	num_events_in_1ms   uint64
	num_events_in_10ms  uint64
	num_events_in_100ms uint64
	num_events_in_1s    uint64
	num_events_in_10s   uint64
	num_events_ge_100s  uint64
}

// Number of events counted.
func (d DOOMGram) EventCount() uint64 {
	return d.event_count
}

// Attempts to obtain the total event time (in nanoseconds).
func (d DOOMGram) EventTimeTotal() (bool, uint64) {
	if d.has_overflowed {
		return false, 0
	} else {
		return true, d.event_time_total
	}
}

// Obtains the total event time (in nanoseconds), regardless of whether
// overflow has occurred.
func (d DOOMGram) EventTimeTotalRaw() uint64 {
	return d.event_time_total
}

// Indicates whether overflow has occurred.
func (d DOOMGram) Overflowed() bool {
	return d.has_overflowed
}

// Attempts to obtain the minimum event time.
func (d DOOMGram) MinEventTime() (bool, uint64) {
	if 0 == d.event_count {
		return false, 0
	} else {
		return true, d.min_event_time
	}
}

// Attempts to obtain the maximum event time.
func (d DOOMGram) MaxEventTime() (bool, uint64) {
	if 0 == d.event_count {
		return false, 0
	} else {
		return true, d.max_event_time
	}
}

// Number of events counted in the interval [1ns, 10ns).
func (d DOOMGram) NumEventsIn1ns() uint64 {
	return d.num_events_in_1ns
}

// Number of events counted in the interval [10ns, 100ns).
func (d DOOMGram) NumEventsIn10ns() uint64 {
	return d.num_events_in_10ns
}

// Number of events counted in the interval [100ns, 1µs).
func (d DOOMGram) NumEventsIn100ns() uint64 {
	return d.num_events_in_100ns
}

// Number of events counted in the interval [1µs, 10µs).
func (d DOOMGram) NumEventsIn1us() uint64 {
	return d.num_events_in_1us
}

// Number of events counted in the interval [10µs, 100µs).
func (d DOOMGram) NumEventsIn10us() uint64 {
	return d.num_events_in_10us
}

// Number of events counted in the interval [100µs, 1ms).
func (d DOOMGram) NumEventsIn100us() uint64 {
	return d.num_events_in_100us
}

// Number of events counted in the interval [1ms, 10ms).
func (d DOOMGram) NumEventsIn1ms() uint64 {
	return d.num_events_in_1ms
}

// Number of events counted in the interval [10ms, 100ms).
func (d DOOMGram) NumEventsIn10ms() uint64 {
	return d.num_events_in_10ms
}

// Number of events counted in the interval [100ms, 1s).
func (d DOOMGram) NumEventsIn100ms() uint64 {
	return d.num_events_in_100ms
}

// Number of events counted in the interval [1s, 10s).
func (d DOOMGram) NumEventsIn1s() uint64 {
	return d.num_events_in_1s
}

// Number of events counted in the interval [10s, 100s).
func (d DOOMGram) NumEventsIn10s() uint64 {
	return d.num_events_in_10s
}

// Number of events counted in the interval [100s, ∞).
func (d DOOMGram) NumEventsIe100s() uint64 {
	return d.num_events_ge_100s
}

func (d *DOOMGram) PushEventTimeNs(time_in_ns uint64) bool {
	if d.tryAddNsToTotalAndUpdateMinmaxAndCount(time_in_ns) {

		d.event_count += 1

		d.pushEventTimeNs_(time_in_ns)

		return true
	} else {
		return false
	}
}

func (d *DOOMGram) PushEventTimeUs(time_in_us uint64) bool {
	time_in_ns := 1000 * time_in_us

	if d.tryAddNsToTotalAndUpdateMinmaxAndCount(time_in_ns) {

		d.event_count += 1

		d.pushEventTimeNs_(time_in_ns)

		return true
	} else {
		return false
	}
}

func (d *DOOMGram) PushEventTimeMs(time_in_ms uint64) bool {
	time_in_ns := 1000 * 1000 * time_in_ms

	if d.tryAddNsToTotalAndUpdateMinmaxAndCount(time_in_ns) {

		d.event_count += 1

		d.pushEventTimeNs_(time_in_ns)

		return true
	} else {
		return false
	}
}

func (d *DOOMGram) PushEventTimeS(time_in_s uint64) bool {
	time_in_ns := 1000 * 1000 * 1000 * time_in_s

	if d.tryAddNsToTotalAndUpdateMinmaxAndCount(time_in_ns) {

		d.event_count += 1

		d.pushEventTimeNs_(time_in_ns)

		return true
	} else {
		return false
	}
}

func (d *DOOMGram) pushEventTimeNs_(time_in_ns uint64) {
	if time_in_ns >= 1000000 {
		// >= 1ms

		if time_in_ns >= 1000000000 {
			// >= 1s

			if time_in_ns >= 100000000000 {
				d.num_events_ge_100s += 1
			} else if time_in_ns >= 10000000000 {
				d.num_events_in_10s += 1
			} else {
				d.num_events_in_1s += 1
			}
		} else {
			// < 1s

			if time_in_ns >= 100000000 {
				d.num_events_in_100ms += 1
			} else if time_in_ns >= 10000000 {
				d.num_events_in_10ms += 1
			} else {
				d.num_events_in_1ms += 1
			}
		}
	} else {
		// < 1ms

		if time_in_ns >= 1000 {
			// >= 1µs

			if time_in_ns >= 100000 {
				d.num_events_in_100us += 1
			} else if time_in_ns >= 10000 {
				d.num_events_in_10us += 1
			} else {
				d.num_events_in_1us += 1
			}
		} else {
			// < 1µs

			if time_in_ns >= 100 {
				d.num_events_in_100ns += 1
			} else if time_in_ns >= 10 {
				d.num_events_in_10ns += 1
			} else if time_in_ns >= 1 {
				d.num_events_in_1ns += 1
			}
		}
	}
}

func (d *DOOMGram) tryAddNsToTotalAndUpdateMinmaxAndCount(time_in_ns uint64) bool {
	if d.has_overflowed {
		return false
	}

	new_total := d.event_time_total + time_in_ns

	if new_total < d.event_time_total || new_total < time_in_ns {
		d.has_overflowed = true

		return false
	} else {

		d.event_time_total = new_total

		if 0 == d.event_count {
			d.min_event_time = time_in_ns
			d.max_event_time = time_in_ns
		} else {
			if time_in_ns < d.min_event_time {
				d.min_event_time = time_in_ns
			}
			if time_in_ns > d.max_event_time {
				d.max_event_time = time_in_ns
			}
		}

		return true
	}
}

// NOTE: taken directly from Diagnosticism.Rust
func calc_doom(v uint64) uint32 {

	if v >= 100000000 {
		//    return count_decimal_digits(v);
	} else {
		if v >= 10000 {
			if v >= 1000000 {
				if v >= 10000000 {
					return 8
				} else {
					return 7
				}
			} else {
				if v >= 100000 {
					return 6
				} else {
					return 5
				}
			}
		} else {
			if v >= 100 {
				if v >= 1000 {
					return 4
				} else {
					return 3
				}
			} else {
				if v >= 10 {
					return 2
				} else {
					if v > 0 {
						return 1
					} else {
						return 0
					}
				}
			}
		}
	}

	if 0 == v {
		return 0
	} else {
		r := uint32(0)
		v := v

		for 0 != v {
			v /= 10
			r += 1
		}

		return r
	}
}

// NOTE: taken directly from Diagnosticism.Rust
func gram_doom_to_char(
	doom uint32,
	ch_0 byte,
	ch_overflow uint8,
	rng []byte,
) byte {
	if 0 == doom {
		return ch_0
	} else {
		if doom > uint32(len(rng)) {
			return ch_overflow
		} else {
			return rng[doom-1]
		}
	}
}

// Converts a [DOOMGram] instance into a timing strip string
// representation.
func (d DOOMGram) ToStrip() string {
	var strip [12]byte

	ch_0 := byte('_')
	ch_overflow := byte('*')
	rng := []byte("abcdefghijklmnopqrstuvwxyz")

	strip[0] = gram_doom_to_char(calc_doom(d.NumEventsIn1ns()), ch_0, ch_overflow, rng)
	strip[1] = gram_doom_to_char(calc_doom(d.NumEventsIn10ns()), ch_0, ch_overflow, rng)
	strip[2] = gram_doom_to_char(calc_doom(d.NumEventsIn100ns()), ch_0, ch_overflow, rng)
	strip[3] = gram_doom_to_char(calc_doom(d.NumEventsIn1us()), ch_0, ch_overflow, rng)
	strip[4] = gram_doom_to_char(calc_doom(d.NumEventsIn10us()), ch_0, ch_overflow, rng)
	strip[5] = gram_doom_to_char(calc_doom(d.NumEventsIn100us()), ch_0, ch_overflow, rng)
	strip[6] = gram_doom_to_char(calc_doom(d.NumEventsIn1ms()), ch_0, ch_overflow, rng)
	strip[7] = gram_doom_to_char(calc_doom(d.NumEventsIn10ms()), ch_0, ch_overflow, rng)
	strip[8] = gram_doom_to_char(calc_doom(d.NumEventsIn100ms()), ch_0, ch_overflow, rng)
	strip[9] = gram_doom_to_char(calc_doom(d.NumEventsIn1s()), ch_0, ch_overflow, rng)
	strip[10] = gram_doom_to_char(calc_doom(d.NumEventsIn10s()), ch_0, ch_overflow, rng)
	strip[11] = gram_doom_to_char(calc_doom(d.NumEventsIe100s()), ch_0, ch_overflow, rng)

	return string(strip[:])
}
