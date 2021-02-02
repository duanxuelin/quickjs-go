package quickjs

import "sync"

type runtimeTracks struct {
	mutex sync.RWMutex
	rts   map[uintptr]*Runtime
}

var tracks = &runtimeTracks{sync.RWMutex{}, make(map[uintptr]*Runtime, 32)}

func (t *runtimeTracks) get(id uintptr) (*Runtime, bool) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	rt, ok := tracks.rts[id]
	return rt, ok
}

func (t *runtimeTracks) add(id uintptr, rt *Runtime) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.rts[id] = rt
}

func (t *runtimeTracks) free(id uintptr) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	delete(t.rts, id)
}
