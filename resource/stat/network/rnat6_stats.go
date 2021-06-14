/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package network

/**
* Statistics for IPv6 RNAT configured route resource.
*/

type Rnat6stats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Rnat6totrxbytes uint64 `json:"rnat6totrxbytes,omitempty"`
	/**
	* Bytes received during RNAT6 sessions.
	*/
	Rnat6rxbytesrate int32 `json:"rnat6rxbytesrate,omitempty"`
	Rnat6tottxbytes uint64 `json:"rnat6tottxbytes,omitempty"`
	/**
	* Bytes sent during RNAT6 sessions.
	*/
	Rnat6txbytesrate int32 `json:"rnat6txbytesrate,omitempty"`
	Rnat6totrxpkts uint64 `json:"rnat6totrxpkts,omitempty"`
	/**
	* Packets received during RNAT6 sessions.
	*/
	Rnat6rxpktsrate int32 `json:"rnat6rxpktsrate,omitempty"`
	Rnat6tottxpkts uint64 `json:"rnat6tottxpkts,omitempty"`
	/**
	* Packets sent during RNAT6 sessions.
	*/
	Rnat6txpktsrate int32 `json:"rnat6txpktsrate,omitempty"`
	Rnat6tottxsyn uint64 `json:"rnat6tottxsyn,omitempty"`
	/**
	* Requests for connections sent during RNAT6 sessions.
	*/
	Rnat6txsynrate int32 `json:"rnat6txsynrate,omitempty"`
	Rnat6cursessions uint32 `json:"rnat6cursessions,omitempty"`

}
