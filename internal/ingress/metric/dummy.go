/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metric

// DummyCollector dummy implementation for mocks in tests
type DummyCollector struct{}

// ConfigSuccess ...
func (dc DummyCollector) ConfigSuccess(uint64, bool) {}

// IncReloadCount ...
func (dc DummyCollector) IncReloadCount() {}

// IncReloadErrorCount ...
func (dc DummyCollector) IncReloadErrorCount() {}

// RemoveMetrics ...
func (dc DummyCollector) RemoveMetrics(ingresses, endpoints []string) {}

// Start ...
func (dc DummyCollector) Start() {}

// Stop ...
func (dc DummyCollector) Stop() {}
