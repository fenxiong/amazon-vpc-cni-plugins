// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package plugin

import (
	"time"

	"github.com/cihub/seelog"
	"github.com/pkg/errors"
	"github.com/sparrc/go-ping"
)

const (
	pingNum = 1
)

func waitEndpointAvailable(endpoint string, timeout time.Duration) (*ping.Statistics, error) {
	pinger, err := ping.NewPinger(endpoint)
	if err != nil {
		return nil, err
	}
	pinger.SetPrivileged(true)
	pinger.Count = pingNum
	pinger.Timeout = 300 * time.Millisecond

	timer := time.NewTimer(timeout)
	statsChan := make(chan *ping.Statistics, 1)

	cancelled := false
	go func() {
		var statsToReturn *ping.Statistics
		for !cancelled {
			stats, err := pingWithPinger(pinger)
			if err != nil {
				time.Sleep(100 * time.Millisecond)
				seelog.Errorf("error pinging %s: %v", endpoint, err)
				continue
			}

			statsToReturn = stats
			break
		}
		statsChan <- statsToReturn
	}()

	select {
	case stats := <-statsChan:
		return stats, nil
	case <-timer.C:
		cancelled = true
		return nil, errors.Errorf("Timed out waiting for endpoint '%s' to be available", endpoint)
	}
}

func pingWithPinger(pinger *ping.Pinger) (*ping.Statistics, error) {
	pinger.Run()
	stats := pinger.Statistics()
	if stats.PacketLoss > 0.0 {
		return nil, errors.Errorf("fail to ping endpoint, packet loss: %v", stats.PacketLoss)
	}

	return stats, nil
}

func pingEndpoint(endpoint string) (*ping.Statistics, error) {
	pinger, err := ping.NewPinger(endpoint)
	if err != nil {
		return nil, err
	}
	pinger.SetPrivileged(true)
	pinger.Count = pingNum
	pinger.Timeout = 300 * time.Millisecond

	pinger.Run()

	stats := pinger.Statistics()
	if stats.PacketLoss > 0.0 {
		return nil, errors.Errorf("fail to ping endpoint, packet loss: %v", stats.PacketLoss)
	}

	return stats, nil
}