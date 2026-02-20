package storage
type MemoryStore struct{Materials map[string]string;Notes map[string]string}
func NewMemoryStore()*MemoryStore{return &MemoryStore{Materials:map[string]string{},Notes:map[string]string{}}}