// Copyright 2016 The Prometheus Authors
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

package prober

import (
	"crypto/tls"
	"time"
)

func getEarliestCertExpiry(state *tls.ConnectionState) time.Time {
	earliest := time.Time{}
	for _, cert := range state.PeerCertificates {
		if (earliest.IsZero() || cert.NotAfter.Before(earliest)) && !cert.NotAfter.IsZero() {
			earliest = cert.NotAfter
		}
	}
	return earliest
}

func getTLSVersion(state *tls.ConnectionState) string {
	switch state.Version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "unknown"
	}
}
