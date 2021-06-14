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
* Statistics for "VXLAN" resource.
*/

type Vxlanstats struct {
	/**
	* An integer specifying the VXLAN identification number (VNID).
	*/
	Id uint32 `json:"id,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Vxlantotrxpkts uint64 `json:"vxlantotrxpkts,omitempty"`
	/**
	* Packets received on the VXLAN.
	*/
	Vxlanrxpktsrate int32 `json:"vxlanrxpktsrate,omitempty"`
	Vxlantotrxbytes uint64 `json:"vxlantotrxbytes,omitempty"`
	/**
	* Bytes of data received on the VXLAN.
	*/
	Vxlanrxbytesrate int32 `json:"vxlanrxbytesrate,omitempty"`
	Vxlantottxpkts uint64 `json:"vxlantottxpkts,omitempty"`
	/**
	* Packets transmitted on the VXLAN.
	*/
	Vxlantxpktsrate int32 `json:"vxlantxpktsrate,omitempty"`
	Vxlantottxbytes uint64 `json:"vxlantottxbytes,omitempty"`
	/**
	* Bytes of data transmitted on the VXLAN.
	*/
	Vxlantxbytesrate int32 `json:"vxlantxbytesrate,omitempty"`

}
