// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package licenser

// CallbackWatcher defines an addhoc listener for events generated by the manager.
type CallbackWatcher struct {
	New     func(License)
	Stopped func()
}

// OnNewLicense is called when a new license is set in the manager.
func (cb *CallbackWatcher) OnNewLicense(license License) {
	if cb.New == nil {
		return
	}
	cb.New(license)
}

// OnManagerStopped is called when the manager is stopped, watcher are expected to terminates any
// features that depends on a specific license.
func (cb *CallbackWatcher) OnManagerStopped() {
	if cb.Stopped == nil {
		return
	}

	cb.Stopped()
}