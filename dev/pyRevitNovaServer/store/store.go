package store

import "sync"

// Store is a thread-safe in-memory store for pyRevit records.
type Store struct {
	mu      sync.RWMutex
	scripts []ScriptRecord
	events  []EventRecord
}

// NewStore creates a new in-memory Store.
func NewStore() *Store {
	return &Store{
		scripts: make([]ScriptRecord, 0),
		events:  make([]EventRecord, 0),
	}
}

// AddScript appends a script execution record to the store.
func (s *Store) AddScript(rec ScriptRecord) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.scripts = append(s.scripts, rec)
}

// GetScripts returns a copy of all script execution records.
func (s *Store) GetScripts() []ScriptRecord {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]ScriptRecord, len(s.scripts))
	copy(result, s.scripts)
	return result
}

// AddEvent appends an application event record to the store.
func (s *Store) AddEvent(rec EventRecord) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, rec)
}

// GetEvents returns a copy of all application event records.
func (s *Store) GetEvents() []EventRecord {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]EventRecord, len(s.events))
	copy(result, s.events)
	return result
}

// Stats returns summary statistics for the stored records.
func (s *Store) Stats() map[string]int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return map[string]int{
		"scripts": len(s.scripts),
		"events":  len(s.events),
	}
}
