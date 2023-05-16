package json

type ByteString []byte

func (s ByteString) MarshalJSON() ([]byte, error) {
	return Marshal(s)
}

func (s *ByteString) UnmarshalJSON(b []byte) error {
	return Unmarshal(b, s)
}
