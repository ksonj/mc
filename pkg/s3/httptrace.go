/*
 * Mini Object Storage, (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package s3

import (
	"errors"
	"net/http"
)

type HttpTracer interface {
	Request(req *http.Request)
	Response(res *http.Response)
}

type RoundTripTrace struct {
	Trace     HttpTracer
	Transport http.RoundTripper
}

func (t RoundTripTrace) RoundTrip(req *http.Request) (res *http.Response, err error) {
	if t.Trace != nil {
		t.Trace.Request(req)
	}
	if t.Transport != nil {
		res, err = t.Transport.RoundTrip(req)
	} else {
		return nil, errors.New("TraceRoundTrip.Transport is nil")
	}

	if t.Trace != nil {
		t.Trace.Response(res)
	}
	return res, err
}
