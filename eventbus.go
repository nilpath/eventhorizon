// Copyright (c) 2014 - Max Ekman <max@looplab.se>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventhorizon

// EventHandler is an interface that all handlers of events should implement.
type EventHandler interface {
	// HandleEvent handles an event.
	HandleEvent(Event)
}

// EventBus is an interface defining an event bus for distributing events.
type EventBus interface {
	// PublishEvent publishes an event on the event bus.
	PublishEvent(Event)
	// AddHandler adds a handler for a specific local event.
	AddHandler(EventHandler, Event)
	// AddLocalHandler adds a handler for local events.
	AddLocalHandler(EventHandler)
	// AddGlobalHandler adds a handler for global (remote) events.
	AddGlobalHandler(EventHandler)
}
