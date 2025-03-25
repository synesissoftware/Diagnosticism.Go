package diagnosticism

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

func (d DOOMGram) EventCount() uint64 {
	return d.event_count
}

func (d DOOMGram) EventTimeTotal() (bool, uint64) {
	if d.has_overflowed {
		return false, 0
	} else {
		return true, d.event_time_total
	}
}

func (d DOOMGram) EventTimeTotalRaw() uint64 {
	return d.event_time_total
}

func (d DOOMGram) Overflowed() bool {
	return d.has_overflowed
}

func (d DOOMGram) MinEventTime() (bool, uint64) {
	if 0 == d.event_count {
		return false, 0
	} else {
		return true, d.min_event_time
	}
}

func (d DOOMGram) MaxEventTime() (bool, uint64) {
	if 0 == d.event_count {
		return false, 0
	} else {
		return true, d.max_event_time
	}
}

func (d DOOMGram) NumEventsIn1ns() uint64 {
	return d.num_events_in_1ns
}

func (d DOOMGram) NumEventsIn10ns() uint64 {
	return d.num_events_in_10ns
}

func (d DOOMGram) NumEventsIn100ns() uint64 {
	return d.num_events_in_100ns
}

func (d DOOMGram) NumEventsIn1us() uint64 {
	return d.num_events_in_1us
}

func (d DOOMGram) NumEventsIn10us() uint64 {
	return d.num_events_in_10us
}

func (d DOOMGram) NumEventsIn100us() uint64 {
	return d.num_events_in_100us
}

func (d DOOMGram) NumEventsIn1ms() uint64 {
	return d.num_events_in_1ms
}

func (d DOOMGram) NumEventsIn10ms() uint64 {
	return d.num_events_in_10ms
}

func (d DOOMGram) NumEventsIn100ms() uint64 {
	return d.num_events_in_100ms
}

func (d DOOMGram) NumEventsIn1s() uint64 {
	return d.num_events_in_1s
}

func (d DOOMGram) NumEventsIn10s() uint64 {
	return d.num_events_in_10s
}

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
