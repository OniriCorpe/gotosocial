// GoToSocial
// Copyright (C) GoToSocial Authors admin@gotosocial.org
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package queue

import (
	"sync/atomic"

	"codeberg.org/gruf/go-structr"
)

// StructQueue wraps a structr.Queue{} to
// provide simple index caching by name.
type StructQueue[StructType any] struct {
	queue structr.Queue[StructType]
	index map[string]*structr.Index
	wait  atomic.Pointer[chan struct{}]
}

// Init initializes queue with structr.QueueConfig{}.
func (q *StructQueue[T]) Init(config structr.QueueConfig[T]) {
	q.index = make(map[string]*structr.Index, len(config.Indices))
	q.queue = structr.Queue[T]{}
	q.queue.Init(config)
	for _, cfg := range config.Indices {
		q.index[cfg.Fields] = q.queue.Index(cfg.Fields)
	}
}

// Pop: see structr.Queue{}.PopFront().
func (q *StructQueue[T]) Pop() (value T, ok bool) {
	return q.queue.PopFront()
}

// Push wraps structr.Queue{}.PushBack() to awaken those blocking on <-.Wait().
func (q *StructQueue[T]) Push(values ...T) {
	q.queue.PushBack(values...)
	q.broadcast()
}

// Delete pops (and drops!) all queued entries under index with key.
func (q *StructQueue[T]) Delete(index string, key ...any) {
	i := q.index[index]
	_ = q.queue.Pop(i, i.Key(key...))
}

// Len: see structr.Queue{}.Len().
func (q *StructQueue[T]) Len() int {
	return q.queue.Len()
}

// Wait returns current wait channel, which may be
// blocked on to awaken when new value pushed to queue.
func (q *StructQueue[T]) Wait() <-chan struct{} {
	var ch chan struct{}

	for {
		// Get channel ptr.
		ptr := q.wait.Load()
		if ptr != nil {
			return *ptr
		}

		if ch == nil {
			// Allocate new channel.
			ch = make(chan struct{})
		}

		// Try set the new wait channel ptr.
		if q.wait.CompareAndSwap(ptr, &ch) {
			return ch
		}
	}
}

// broadcast safely closes wait channel if
// currently set, releasing waiting goroutines.
func (q *StructQueue[T]) broadcast() {
	if ptr := q.wait.Swap(nil); ptr != nil {
		close(*ptr)
	}
}
