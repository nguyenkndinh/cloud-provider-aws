/*
Copyright 2020 The Kubernetes Authors.
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

package tagging

import (
	"k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/legacyregistry"
	"sync"
)

const (
	// subSystemName is the name of this subsystem name used for prometheus metrics.
	subSystemName = "tagging_controller"
)

var register sync.Once

// registerMetrics registers tagging-controller metrics.
func registerMetrics() {
	register.Do(func() {
		legacyregistry.MustRegister(nodeTaggedLatency)
		legacyregistry.MustRegister(nodeUntaggedLatency)
	})
}

var (
	nodeTaggedLatency = metrics.NewHistogram(&metrics.HistogramOpts{
		Name:      "node_tag_latency_seconds",
		Subsystem: subSystemName,
		Help:      "A metric measuring the latency for tagging node resources.",
		// Buckets from 1s to 16384s
		Buckets:        metrics.ExponentialBuckets(1, 2, 15),
		StabilityLevel: metrics.ALPHA,
	})

	nodeUntaggedLatency = metrics.NewHistogram(&metrics.HistogramOpts{
		Name:      "node_untag_latency_seconds",
		Subsystem: subSystemName,
		Help:      "A metric measuring the latency for untagging node resources.",
		// Buckets from 1s to 16384s
		Buckets:        metrics.ExponentialBuckets(1, 2, 15),
		StabilityLevel: metrics.ALPHA,
	})
)
