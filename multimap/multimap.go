// Package multimap implements a slice based multimap.
//
// Multimap is a collection that maps keys to values, similar to map.
// However, each key may be associated with multiple values.
//
// You can visualize the contents of a multimap either as a map from keys to nonempty collections of values:
//    - a --> 1, 2
//    - b --> 3
// ... or a single "flattened" collection of key-value pairs.
//    - a --> 1
//    - a --> 2
//    - b --> 3
//
// Elements are unordered in the map.
//
// Structure is not thread safe.
//
package multimap

// MultiMap holds the elements in go's native map
type MultiMap struct {
	m map[interface{}][]interface{}
}

// New instantiates a new multimap.
func New() *MultiMap {
	return &MultiMap{m: make(map[interface{}][]interface{})}
}

// Get searches the element in the multimap by key.
// It returns its value or nil if key is not found in multimap.
// Second return parameter is true if key was found, otherwise false.
func (m *MultiMap) Get(key interface{}) (values []interface{}, found bool) {
	values, found = m.m[key]
	return
}

// Put stores a key-value pair in this multimap.
func (m *MultiMap) Put(key interface{}, value interface{}) {
	values, found := m.m[key]
	if found {
		m.m[key] = append(values, value)
	} else {
		values := make([]interface{}, 0)
		m.m[key] = append(values, value)
	}
}

// PutAll stores a key-value pair in this multimap for each of the values, all using the same key key.
func (m *MultiMap) PutAll(key interface{}, values []interface{}) {
	for _, value := range values {
		m.Put(key, value)
	}
}

// Contains returns true if this multimap contains at least one key-value pair with the key key and the value value.
func (m *MultiMap) Contains(key interface{}, value interface{}) bool {
	values, found := m.m[key]
	for _, v := range values {
		if v == value {
			return true && found
		}
	}
	return false && found
}

// ContainsKey returns true if this multimap contains at least one key-value pair with the key key.
func (m *MultiMap) ContainsKey(key interface{}) (found bool) {
	_, found = m.m[key]
	return
}

// ContainsValue returns true if this multimap contains at least one key-value pair with the value value.
func (m *MultiMap) ContainsValue(value interface{}) bool {
	for _, values := range m.m {
		for _, v := range values {
			if v == value {
				return true
			}
		}
	}
	return false
}

// Remove removes a single key-value pair from this multimap, if such exists.
func (m *MultiMap) Remove(key interface{}, value interface{}) {
	values, found := m.m[key]
	if found {
		for i, v := range values {
			if v == value {
				m.m[key] = append(values[:i], values[i+1:]...)
			}
		}
	}
	if len(m.m[key]) == 0 {
		delete(m.m, key)
	}
}

// RemoveAll removes all values associated with the key from the multimap.
func (m *MultiMap) RemoveAll(key interface{}) {
	delete(m.m, key)
}

// Empty returns true if multimap does not contain any key-value pairs.
func (m *MultiMap) Empty() bool {
	return m.Size() == 0
}

// Size returns number of key-value pairs in the multimap.
func (m *MultiMap) Size() int {
	size := 0
	for _, values := range m.m {
		size += len(values)
	}
	return size
}

// Keys returns a view collection containing the key from each key-value pair in this multimap.
// This is done without collapsing duplicates.
func (m *MultiMap) Keys() []interface{} {
	keys := make([]interface{}, m.Size())
	count := 0
	for key, values := range m.m {
		for range values {
			keys[count] = key
			count++
		}
	}
	return keys
}

// KeySet returns all distinct keys contained in this multimap.
func (m *MultiMap) KeySet() []interface{} {
	keys := make([]interface{}, len(m.m))
	count := 0
	for key := range m.m {
		keys[count] = key
		count++
	}
	return keys
}

// Values returns all values from each key-value pair contained in this multimap.
// This is done without collapsing duplicates. (size of Values() = MultiMap.Size()).
func (m *MultiMap) Values() []interface{} {
	values := make([]interface{}, m.Size())
	count := 0
	for _, vs := range m.m {
		for _, value := range vs {
			values[count] = value
			count++
		}
	}
	return values
}

// Entries view collection of all key-value pairs contained in this multimap.
// The return type is a slice of interface{}, which will be a struct of key/value instances.
// Retrieving the key and value from the entries result will be as trivial as:
//   - var entry = m.Entries()[0]
//   - var key = entry.key
//   - var value = entry.value
func (m *MultiMap) Entries() []interface{} {
	entries := make([]interface{}, m.Size())
	count := 0
	for key, values := range m.m {
		for _, value := range values {
			entries[count] = struct {
				key   interface{}
				value interface{}
			}{key, value}
			count++
		}
	}
	return entries
}

// Clear removes all elements from the map.
func (m *MultiMap) Clear() {
	m.m = make(map[interface{}][]interface{})
}
