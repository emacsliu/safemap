package safemap

import (
	"github.com/pkg/errors"
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	objects map[string]interface{}
}

var _ SafeMapInterface = &SafeMap{}

func NewSafeMap() *SafeMap {
	return newSafeMap()
}

func newSafeMap() *SafeMap {
	return &SafeMap{
		objects: make(map[string]interface{}),
	}
}

func (m *SafeMap) GetObject(key string) (interface{}, bool) {
	if invalidKey(key) {
		return nil, false
	}
	m.RLock()
	defer m.RUnlock()
	value, ok := m.objects[key]
	return value, ok
}

func (m *SafeMap) SetObject(key string, obj interface{}) error {
	if invalidKey(key) {
		return errors.New("invalid key")
	}
	m.Lock()
	defer m.Unlock()
	m.objects[key] = obj
	return nil
}

func (m *SafeMap) RemoveObject(key string) error {
	if invalidKey(key) {
		return errors.New("invalid key")
	}
	m.Lock()
	defer m.Unlock()
	delete(m.objects, key)
	return nil
}

func (m *SafeMap) Values() []interface{} {
	m.RLock()
	defer m.RUnlock()
	values := []interface{}{}
	for _, v := range m.objects {
		values = append(values, v)
	}
	return values
}

func (m *SafeMap) Keys() []string {
	m.RLock()
	defer m.RUnlock()
	names := []string{}
	for key, _ := range m.objects {
		names = append(names, key)
	}
	return names
}

func invalidKey(key string) bool {
	return len(key) <= 0
}
