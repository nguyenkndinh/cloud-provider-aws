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

var register sync.Once

var (
	workItemLatency = metrics.NewHistogramVec(
		&metrics.HistogramOpts{
			Name:           "tagging_workitem_latency",
			Help:           "workitem latency",
			StabilityLevel: metrics.ALPHA,
		},
		[]string{"workqueue"})

	workItemError = metrics.NewCounterVec(
		&metrics.CounterOpts{
			Name:           "tagging_workitem_error",
			Help:           "workitem error",
			StabilityLevel: metrics.ALPHA,
		},
		[]string{"workqueue"})
)

// registerMetrics registers tagging-controller metrics.
func registerMetrics() {
	register.Do(func() {
		legacyregistry.MustRegister(workItemLatency)
		legacyregistry.MustRegister(workItemError)
	})
}

func recordWorkItemLatencyMetrics(actionName string, timeTaken float64) {
	workItemLatency.With(metrics.Labels{"workqueue": actionName}).Observe(timeTaken)
}

func recordWorkItemErrorMetrics(actionName string) {
	workItemError.With(metrics.Labels{"workqueue": actionName}).Inc()
}
