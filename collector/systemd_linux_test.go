// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package collector

import (
	"testing"

	"github.com/coreos/go-systemd/dbus"
	"github.com/prometheus/client_golang/prometheus"
)

// Creates mock UnitLists
func getUnitListFixtures() [][]dbus.UnitStatus {
	fixture1 := []dbus.UnitStatus{
		dbus.UnitStatus{
			Name:        "foo",
			Description: "foo desc",
			LoadState:   "loaded",
			ActiveState: "active",
			SubState:    "running",
			Followed:    "",
			Path:        "/org/freedesktop/systemd1/unit/foo",
			JobId:       0,
			JobType:     "",
			JobPath:     "/",
		},
		dbus.UnitStatus{
			Name:        "bar",
			Description: "bar desc",
			LoadState:   "not-found",
			ActiveState: "inactive",
			SubState:    "dead",
			Followed:    "",
			Path:        "/org/freedesktop/systemd1/unit/bar",
			JobId:       0,
			JobType:     "",
			JobPath:     "/",
		},
	}

	fixture2 := []dbus.UnitStatus{}

	return [][]dbus.UnitStatus{fixture1, fixture2}
}

func TestSystemdCollectorDoesntCrash(t *testing.T) {
	c, err := NewSystemdCollector()
	if err != nil {
		t.Fatal(err)
	}
	sink := make(chan prometheus.Metric)
	go func() {
		for {
			<-sink
		}
	}()

	fixtures := getUnitListFixtures()
	collector := (c).(*systemdCollector)
	for _, units := range fixtures {
		collector.collectUnitStatusMetrics(sink, units)
	}
}
