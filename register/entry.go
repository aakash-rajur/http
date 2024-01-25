package register

import "net/http"

type Entry struct {
	segments segments
	Handler  http.Handler
}

func (e Entry) cmp(other Entry) int {
	a, b := e.segments, other.segments

	al, bl := len(a), len(b)

	maxLength := max(al, bl)

	for i := 0; i < maxLength; i += 1 {
		if i >= al {
			return -1
		}

		if i >= bl {
			return 1
		}

		segmentA, segmentB := a[i], b[i]

		aParam, bParam := segmentA.isParam(), segmentB.isParam()

		if aParam && !bParam {
			return 1
		}

		if !aParam && bParam {
			return -1
		}

		comparison := segmentA.cmp(segmentB)

		if comparison == 0 {
			continue
		}

		return comparison
	}

	return 0

	/*
		lengthDelta := al - bl

		if lengthDelta != 0 {
			return lengthDelta
		}

		pal, bal := 0, 0

		for i := 0; i < al; i += 1 {
			weight := 10 ^ (al - i)

			pa, pb := a[i], b[i]

			if pa.isParam() {
				pal += weight
			}

			if pb.isParam() {
				bal += weight
			}
		}

		paramDelta := pal - bal

		if paramDelta != 0 {
			return paramDelta
		}

		for i := 0; i < al; i += 1 {
			pa, pb := a[i], b[i]

			comparison := pa.cmp(pb)

			if comparison == 0 {
				continue
			}

			return comparison
		}

		return 0
	*/
}
