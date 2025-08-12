package datastructure

type SimpleString struct {
	hashTable map[string]string
	size      int
}

func NewString() *SimpleString {
	return &SimpleString{
		hashTable: make(map[string]string),
		size:      10,
	}
}

// TODO
func (s *SimpleString) ReHashing(key string) {}

// TODO
func (s *SimpleString) SetTTL(key, value string, ttl int64) {}

func (s *SimpleString) Set(key, value string) (bool, error) {
	s.hashTable[key] = value

	return true, nil
}

func (s *SimpleString) Get(key string) (string, error) {
	return s.hashTable[key], nil
}

func (s *SimpleString) Del(key string) (bool, error) {
	delete(s.hashTable, key)

	return true, nil
}
