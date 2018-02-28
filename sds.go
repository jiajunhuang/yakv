package yakv

// SDS is same as simple dynamic string in redis/sds.{h,c}
type SDS struct {
	noCopy
	data []byte
}

func (s *SDS) Append(a []byte) int {
	s.data = append(s.data, a...)
	return len(s.data)
}

func (s *SDS) Len() int {
	return len(s.data)
}

func (s *SDS) GetRange(left, right int) []byte {
	length := len(s.data)

	if left < 0 {
		if left < -length {
			left = 0
		} else {
			left += length
		}
	}

	right++ // because redis GETRANGE include the right edge
	if right > length {
		right = length
	}

	if left <= right && right <= length && left >= 0 {
		return s.data[left:right]
	}

	return []byte("")
}
