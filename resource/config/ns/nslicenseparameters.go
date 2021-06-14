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

package ns

/**
* Configuration for licenseparameters resource.
*/
type Nslicenseparameters struct {
	/**
	* If ADC remains in grace for the configured hours then first grace alert will be raised
	*/
	Alert1gracetimeout uint32 `json:"alert1gracetimeout,omitempty"`
	/**
	* If ADC remains in grace for the configured hours then major grace alert will be raised
	*/
	Alert2gracetimeout uint32 `json:"alert2gracetimeout,omitempty"`

}
