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

package appfw

/**
* Statistics for application firewall profile resource.
*/

type Appfwprofilestats struct {
	/**
	* Name of the application firewall profile.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Appfirewallrequestsperprofile uint64 `json:"appfirewallrequestsperprofile,omitempty"`
	/**
	* HTTP/HTTPS requests sent to your protected web servers via the Application Firewall.
	*/
	Appfirewallrequestsperprofilerate int32 `json:"appfirewallrequestsperprofilerate,omitempty"`
	Appfirewallreqbytesperprofile uint64 `json:"appfirewallreqbytesperprofile,omitempty"`
	/**
	* Number of bytes transfered for requests
	*/
	Appfirewallreqbytesperprofilerate int32 `json:"appfirewallreqbytesperprofilerate,omitempty"`
	Appfirewallresponsesperprofile uint64 `json:"appfirewallresponsesperprofile,omitempty"`
	/**
	* HTTP/HTTPS responses sent by your protected web servers via the Application Firewall.
	*/
	Appfirewallresponsesperprofilerate int32 `json:"appfirewallresponsesperprofilerate,omitempty"`
	Appfirewallresbytesperprofile uint64 `json:"appfirewallresbytesperprofile,omitempty"`
	/**
	* Number of bytes transfered for responses
	*/
	Appfirewallresbytesperprofilerate int32 `json:"appfirewallresbytesperprofilerate,omitempty"`
	Appfirewallabortsperprofile uint64 `json:"appfirewallabortsperprofile,omitempty"`
	/**
	* Incomplete HTTP/HTTPS requests aborted by the client before the Application Firewall could finish processing them.
	*/
	Appfirewallabortsperprofilerate int32 `json:"appfirewallabortsperprofilerate,omitempty"`
	Appfirewallredirectsperprofile uint64 `json:"appfirewallredirectsperprofile,omitempty"`
	/**
	* HTTP/HTTPS requests redirected by the Application Firewall to a different Web page or web server. (HTTP 302)
	*/
	Appfirewallredirectsperprofilerate int32 `json:"appfirewallredirectsperprofilerate,omitempty"`
	Appfirewalllongavgresptimeperprofile uint64 `json:"appfirewalllongavgresptimeperprofile,omitempty"`
	Appfirewallshortavgresptimeperprofile uint64 `json:"appfirewallshortavgresptimeperprofile,omitempty"`
	Appfirewallviolstarturlperprofile uint64 `json:"appfirewallviolstarturlperprofile,omitempty"`
	/**
	* Number of Start URL security check violations seen by the Application Firewall.
	*/
	Appfirewallviolstarturlperprofilerate int32 `json:"appfirewallviolstarturlperprofilerate,omitempty"`
	Appfirewallvioldenyurlperprofile uint64 `json:"appfirewallvioldenyurlperprofile,omitempty"`
	/**
	* Number of Deny URL security check violations seen by the Application Firewall.
	*/
	Appfirewallvioldenyurlperprofilerate int32 `json:"appfirewallvioldenyurlperprofilerate,omitempty"`
	Appfirewallviolrefererheaderperprofile uint64 `json:"appfirewallviolrefererheaderperprofile,omitempty"`
	/**
	* Number of Referer Header security check violations seen by the Application Firewall.
	*/
	Appfirewallviolrefererheaderperprofilerate int32 `json:"appfirewallviolrefererheaderperprofilerate,omitempty"`
	Appfirewallviolbufferoverflowperprofile uint64 `json:"appfirewallviolbufferoverflowperprofile,omitempty"`
	/**
	* Number of Buffer Overflow security check violations seen by the Application Firewall.
	*/
	Appfirewallviolbufferoverflowperprofilerate int32 `json:"appfirewallviolbufferoverflowperprofilerate,omitempty"`
	Appfirewallpostbodylimitviolationsperprofile uint64 `json:"appfirewallpostbodylimitviolationsperprofile,omitempty"`
	/**
	* Number of Post Body Limit security check violations seen by the Application Firewall.
	*/
	Appfirewallpostbodylimitviolationsperprofilerate int32 `json:"appfirewallpostbodylimitviolationsperprofilerate,omitempty"`
	Appfirewallviolcookieperprofile uint64 `json:"appfirewallviolcookieperprofile,omitempty"`
	/**
	* Number of Cookie Consistency security check violations seen by the Application Firewall.
	*/
	Appfirewallviolcookieperprofilerate int32 `json:"appfirewallviolcookieperprofilerate,omitempty"`
	Appfirewallviolcookiehijackperprofile uint64 `json:"appfirewallviolcookiehijackperprofile,omitempty"`
	/**
	* Number of Cookie Hijacking security violations seen by the Application Firewall.
	*/
	Appfirewallviolcookiehijackperprofilerate int32 `json:"appfirewallviolcookiehijackperprofilerate,omitempty"`
	Appfirewallviolcsrftagperprofile uint64 `json:"appfirewallviolcsrftagperprofile,omitempty"`
	/**
	* Number of Cross Site Request Forgery form tag security check violations seen by the Application Firewall.
	*/
	Appfirewallviolcsrftagperprofilerate int32 `json:"appfirewallviolcsrftagperprofilerate,omitempty"`
	Appfirewallviolxssperprofile uint64 `json:"appfirewallviolxssperprofile,omitempty"`
	/**
	* Number of HTML Cross-Site Scripting security check violations seen by the Application Firewall.
	*/
	Appfirewallviolxssperprofilerate int32 `json:"appfirewallviolxssperprofilerate,omitempty"`
	Appfirewallviolsqlperprofile uint64 `json:"appfirewallviolsqlperprofile,omitempty"`
	/**
	* Number of HTML SQL Injection security check violations seen by the Application Firewall.
	*/
	Appfirewallviolsqlperprofilerate int32 `json:"appfirewallviolsqlperprofilerate,omitempty"`
	Appfirewallviolfieldformatperprofile uint64 `json:"appfirewallviolfieldformatperprofile,omitempty"`
	/**
	* Number of Field Format security check violations seen by the Application Firewall.
	*/
	Appfirewallviolfieldformatperprofilerate int32 `json:"appfirewallviolfieldformatperprofilerate,omitempty"`
	Appfirewallviolfieldconsistencyperprofile uint64 `json:"appfirewallviolfieldconsistencyperprofile,omitempty"`
	/**
	* Number of Field Consistency security check violations seen by the Application Firewall.
	*/
	Appfirewallviolfieldconsistencyperprofilerate int32 `json:"appfirewallviolfieldconsistencyperprofilerate,omitempty"`
	Appfirewallviolcreditcardperprofile uint64 `json:"appfirewallviolcreditcardperprofile,omitempty"`
	/**
	* Number of Credit Card security check violations seen by the Application Firewall.
	*/
	Appfirewallviolcreditcardperprofilerate int32 `json:"appfirewallviolcreditcardperprofilerate,omitempty"`
	Appfirewallviolsafeobjectperprofile uint64 `json:"appfirewallviolsafeobjectperprofile,omitempty"`
	/**
	* Number of Safe Object security check violations seen by the Application Firewall.
	*/
	Appfirewallviolsafeobjectperprofilerate int32 `json:"appfirewallviolsafeobjectperprofilerate,omitempty"`
	Appfirewallviolsignatureperprofile uint64 `json:"appfirewallviolsignatureperprofile,omitempty"`
	/**
	* Number of Signature violations seen by the Application Firewall.
	*/
	Appfirewallviolsignatureperprofilerate int32 `json:"appfirewallviolsignatureperprofilerate,omitempty"`
	Appfirewallviolcontenttypeperprofile uint64 `json:"appfirewallviolcontenttypeperprofile,omitempty"`
	/**
	* Number of Content Type security check violations seen by the Application Firewall.
	*/
	Appfirewallviolcontenttypeperprofilerate int32 `json:"appfirewallviolcontenttypeperprofilerate,omitempty"`
	Appfirewallvioljsondosperprofile uint64 `json:"appfirewallvioljsondosperprofile,omitempty"`
	/**
	* Number of JSON Denial-of-Service security check violations seen by the Application Firewall.
	*/
	Appfirewallvioljsondosperprofilerate int32 `json:"appfirewallvioljsondosperprofilerate,omitempty"`
	Appfirewallvioljsonsqlperprofile uint64 `json:"appfirewallvioljsonsqlperprofile,omitempty"`
	/**
	* Number of JSON SQL Injection security check violations seen by the Application Firewall.
	*/
	Appfirewallvioljsonsqlperprofilerate int32 `json:"appfirewallvioljsonsqlperprofilerate,omitempty"`
	Appfirewallvioljsonxssperprofile uint64 `json:"appfirewallvioljsonxssperprofile,omitempty"`
	/**
	* Number of JSON Cross-Site Scripting (XSS) security check violations seen by the Application Firewall.
	*/
	Appfirewallvioljsonxssperprofilerate int32 `json:"appfirewallvioljsonxssperprofilerate,omitempty"`
	Appfirewallvioljsoncmdperprofile uint64 `json:"appfirewallvioljsoncmdperprofile,omitempty"`
	/**
	* Number of JSON Command Injection security check violations seen by the Application Firewall.
	*/
	Appfirewallvioljsoncmdperprofilerate int32 `json:"appfirewallvioljsoncmdperprofilerate,omitempty"`
	Appfirewallviolfileuploadtypesperprofile uint64 `json:"appfirewallviolfileuploadtypesperprofile,omitempty"`
	/**
	* Number of Field Upload Types security check violations seen by the Application Firewall.
	*/
	Appfirewallviolfileuploadtypesperprofilerate int32 `json:"appfirewallviolfileuploadtypesperprofilerate,omitempty"`
	Appfirewallxmlpayloadcontenttypemismatchperprofile uint64 `json:"appfirewallxmlpayloadcontenttypemismatchperprofile,omitempty"`
	/**
	* Number of Mismatched Content-Type in request with XML Payload security check violations seen by the Application Firewall.
	*/
	Appfirewallxmlpayloadcontenttypemismatchperprofilerate int32 `json:"appfirewallxmlpayloadcontenttypemismatchperprofilerate,omitempty"`
	Appfirewallviolcmdperprofile uint64 `json:"appfirewallviolcmdperprofile,omitempty"`
	/**
	* Number of HTML Command Injection security check violations seen by the Application Firewall.
	*/
	Appfirewallviolcmdperprofilerate int32 `json:"appfirewallviolcmdperprofilerate,omitempty"`
	Appfirewallviolwellformednessviolationsperprofile uint64 `json:"appfirewallviolwellformednessviolationsperprofile,omitempty"`
	/**
	* Number of XML Format security check violations seen by the Application Firewall.
	*/
	Appfirewallviolwellformednessviolationsperprofilerate int32 `json:"appfirewallviolwellformednessviolationsperprofilerate,omitempty"`
	Appfirewallviolxdosviolationsperprofile uint64 `json:"appfirewallviolxdosviolationsperprofile,omitempty"`
	/**
	* Number of XML Denial-of-Service security check violations seen by the Application Firewall.
	*/
	Appfirewallviolxdosviolationsperprofilerate int32 `json:"appfirewallviolxdosviolationsperprofilerate,omitempty"`
	Appfirewallviolmsgvalviolationsperprofile uint64 `json:"appfirewallviolmsgvalviolationsperprofile,omitempty"`
	/**
	* Number of XML Message Validation security check violations seen by the Application Firewall.
	*/
	Appfirewallviolmsgvalviolationsperprofilerate int32 `json:"appfirewallviolmsgvalviolationsperprofilerate,omitempty"`
	Appfirewallviolwsiviolationsperprofile uint64 `json:"appfirewallviolwsiviolationsperprofile,omitempty"`
	/**
	* Number of Web Services Interoperability (WS-I) security check violations seen by the Application Firewall.
	*/
	Appfirewallviolwsiviolationsperprofilerate int32 `json:"appfirewallviolwsiviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlsqlviolationsperprofile uint64 `json:"appfirewallviolxmlsqlviolationsperprofile,omitempty"`
	/**
	* Number of XML SQL Injection security check violations seen by the Application Firewall.
	*/
	Appfirewallviolxmlsqlviolationsperprofilerate int32 `json:"appfirewallviolxmlsqlviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlxssviolationsperprofile uint64 `json:"appfirewallviolxmlxssviolationsperprofile,omitempty"`
	/**
	* Number of XML Cross-Site Scripting (XSS) security check violations seen by the Application Firewall.
	*/
	Appfirewallviolxmlxssviolationsperprofilerate int32 `json:"appfirewallviolxmlxssviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlattachmentviolationsperprofile uint64 `json:"appfirewallviolxmlattachmentviolationsperprofile,omitempty"`
	/**
	* Number of XML Attachment security check violations seen by the Application Firewall.
	*/
	Appfirewallviolxmlattachmentviolationsperprofilerate int32 `json:"appfirewallviolxmlattachmentviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlsoapfaultviolationsperprofile uint64 `json:"appfirewallviolxmlsoapfaultviolationsperprofile,omitempty"`
	/**
	* Number of requests returning soap:fault from the backend server
	*/
	Appfirewallviolxmlsoapfaultviolationsperprofilerate int32 `json:"appfirewallviolxmlsoapfaultviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlgenericviolationsperprofile uint64 `json:"appfirewallviolxmlgenericviolationsperprofile,omitempty"`
	/**
	* Number of requests returning XML generic violation from the backend server
	*/
	Appfirewallviolxmlgenericviolationsperprofilerate int32 `json:"appfirewallviolxmlgenericviolationsperprofilerate,omitempty"`
	Appfirewallviolsqlgramperprofile uint64 `json:"appfirewallviolsqlgramperprofile,omitempty"`
	/**
	* Number of HTML SQL Injection security check violations (reported by SQL grammar) seen by the Application Firewall.
	*/
	Appfirewallviolsqlgramperprofilerate int32 `json:"appfirewallviolsqlgramperprofilerate,omitempty"`
	Appfirewallvioljsonsqlgramperprofile uint64 `json:"appfirewallvioljsonsqlgramperprofile,omitempty"`
	/**
	* Number of JSON SQL Injection security check violations (reported by SQL grammar) seen by the Application Firewall.
	*/
	Appfirewallvioljsonsqlgramperprofilerate int32 `json:"appfirewallvioljsonsqlgramperprofilerate,omitempty"`
	Appfirewalltotalviolperprofile uint64 `json:"appfirewalltotalviolperprofile,omitempty"`
	/**
	* Number of violations seen by the application firewall on per profile basis
	*/
	Appfirewallviolperprofilerate int32 `json:"appfirewallviolperprofilerate,omitempty"`
	Appfirewalllogstarturlperprofile uint64 `json:"appfirewalllogstarturlperprofile,omitempty"`
	/**
	* Number of Start URL security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogstarturlperprofilerate int32 `json:"appfirewalllogstarturlperprofilerate,omitempty"`
	Appfirewalllogdenyurlperprofile uint64 `json:"appfirewalllogdenyurlperprofile,omitempty"`
	/**
	* Number of Deny URL security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogdenyurlperprofilerate int32 `json:"appfirewalllogdenyurlperprofilerate,omitempty"`
	Appfirewalllogrefererheaderperprofile uint64 `json:"appfirewalllogrefererheaderperprofile,omitempty"`
	/**
	* Number of Referer Header security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogrefererheaderperprofilerate int32 `json:"appfirewalllogrefererheaderperprofilerate,omitempty"`
	Appfirewalllogbufferoverflowperprofile uint64 `json:"appfirewalllogbufferoverflowperprofile,omitempty"`
	/**
	* Number of Buffer Overflow security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogbufferoverflowperprofilerate int32 `json:"appfirewalllogbufferoverflowperprofilerate,omitempty"`
	Appfirewallpostbodylimitlogsperprofile uint64 `json:"appfirewallpostbodylimitlogsperprofile,omitempty"`
	/**
	* Number of Post Body Limit security check logs seen by the Application Firewall.
	*/
	Appfirewallpostbodylimitlogsperprofilerate int32 `json:"appfirewallpostbodylimitlogsperprofilerate,omitempty"`
	Appfirewalllogcookieperprofile uint64 `json:"appfirewalllogcookieperprofile,omitempty"`
	/**
	* Number of Cookie Consistency security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogcookieperprofilerate int32 `json:"appfirewalllogcookieperprofilerate,omitempty"`
	Appfirewalllogcookiehijackperprofile uint64 `json:"appfirewalllogcookiehijackperprofile,omitempty"`
	/**
	* Number of Cookie Hijacking security violation log messages generated by the Application Firewall.
	*/
	Appfirewalllogcookiehijackperprofilerate int32 `json:"appfirewalllogcookiehijackperprofilerate,omitempty"`
	Appfirewalllogcsrftagperprofile uint64 `json:"appfirewalllogcsrftagperprofile,omitempty"`
	/**
	* Number of Cross Site Request Forgery form tag security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogcsrftagperprofilerate int32 `json:"appfirewalllogcsrftagperprofilerate,omitempty"`
	Appfirewalllogxssperprofile uint64 `json:"appfirewalllogxssperprofile,omitempty"`
	/**
	* Number of HTML Cross-Site Scripting security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogxssperprofilerate int32 `json:"appfirewalllogxssperprofilerate,omitempty"`
	Appfirewalllogxformxssperprofile uint64 `json:"appfirewalllogxformxssperprofile,omitempty"`
	/**
	* Number of HTML Cross-Site Scripting security check transform log messages generated by the Application Firewall.
	*/
	Appfirewalllogxformxssperprofilerate int32 `json:"appfirewalllogxformxssperprofilerate,omitempty"`
	Appfirewalllogsqlperprofile uint64 `json:"appfirewalllogsqlperprofile,omitempty"`
	/**
	* Number of HTML SQL Injection security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogsqlperprofilerate int32 `json:"appfirewalllogsqlperprofilerate,omitempty"`
	Appfirewalllogxformsqlperprofile uint64 `json:"appfirewalllogxformsqlperprofile,omitempty"`
	/**
	* Number of HTML SQL Injection security check transform log messages generated by the Application Firewall.
	*/
	Appfirewalllogxformsqlperprofilerate int32 `json:"appfirewalllogxformsqlperprofilerate,omitempty"`
	Appfirewalllogfieldformatperprofile uint64 `json:"appfirewalllogfieldformatperprofile,omitempty"`
	/**
	* Number of Field Format security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogfieldformatperprofilerate int32 `json:"appfirewalllogfieldformatperprofilerate,omitempty"`
	Appfirewalllogfieldconsistencyperprofile uint64 `json:"appfirewalllogfieldconsistencyperprofile,omitempty"`
	/**
	* Number of Field Consistency security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogfieldconsistencyperprofilerate int32 `json:"appfirewalllogfieldconsistencyperprofilerate,omitempty"`
	Appfirewalllogcreditcardperprofile uint64 `json:"appfirewalllogcreditcardperprofile,omitempty"`
	/**
	* Number of Credit Card security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogcreditcardperprofilerate int32 `json:"appfirewalllogcreditcardperprofilerate,omitempty"`
	Appfirewallxformlogcreditcardperprofile uint64 `json:"appfirewallxformlogcreditcardperprofile,omitempty"`
	/**
	* Number of Credit Card security check transform log messages generated by the Application Firewall.
	*/
	Appfirewallxformlogcreditcardperprofilerate int32 `json:"appfirewallxformlogcreditcardperprofilerate,omitempty"`
	Appfirewalllogsafeobjectperprofile uint64 `json:"appfirewalllogsafeobjectperprofile,omitempty"`
	/**
	* Number of Safe Object security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogsafeobjectperprofilerate int32 `json:"appfirewalllogsafeobjectperprofilerate,omitempty"`
	Appfirewalllogcontenttypeperprofile uint64 `json:"appfirewalllogcontenttypeperprofile,omitempty"`
	/**
	* Number of Content type security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogcontenttypeperprofilerate int32 `json:"appfirewalllogcontenttypeperprofilerate,omitempty"`
	Appfirewalllogsjsondosperprofile uint64 `json:"appfirewalllogsjsondosperprofile,omitempty"`
	/**
	* Number of JSON Denial-of-Service security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogsjsondosperprofilerate int32 `json:"appfirewalllogsjsondosperprofilerate,omitempty"`
	Appfirewalllogsjsonsqlperprofile uint64 `json:"appfirewalllogsjsonsqlperprofile,omitempty"`
	/**
	* Number of JSON SQL Injection security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogsjsonsqlperprofilerate int32 `json:"appfirewalllogsjsonsqlperprofilerate,omitempty"`
	Appfirewalllogsjsonxssperprofile uint64 `json:"appfirewalllogsjsonxssperprofile,omitempty"`
	/**
	* Number of JSON Cross-Site Scripting (XSS) security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogsjsonxssperprofilerate int32 `json:"appfirewalllogsjsonxssperprofilerate,omitempty"`
	Appfirewalllogsjsoncmdperprofile uint64 `json:"appfirewalllogsjsoncmdperprofile,omitempty"`
	/**
	* Number of JSON Command Injection security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogsjsoncmdperprofilerate int32 `json:"appfirewalllogsjsoncmdperprofilerate,omitempty"`
	Appfirewalllogfileuploadtypesperprofile uint64 `json:"appfirewalllogfileuploadtypesperprofile,omitempty"`
	/**
	* Number of File Upload Types security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogfileuploadtypesperprofilerate int32 `json:"appfirewalllogfileuploadtypesperprofilerate,omitempty"`
	Appfirewallloginfercontenttypexmlpayloadperprofile uint64 `json:"appfirewallloginfercontenttypexmlpayloadperprofile,omitempty"`
	/**
	* Number of Mismatched Content-Type in request with XML Payload security check logs seen by the Application Firewall.
	*/
	Appfirewallloginfercontenttypexmlpayloadperprofilerate int32 `json:"appfirewallloginfercontenttypexmlpayloadperprofilerate,omitempty"`
	Appfirewalllogcmdperprofile uint64 `json:"appfirewalllogcmdperprofile,omitempty"`
	/**
	* Number of HTML Command Injection security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogcmdperprofilerate int32 `json:"appfirewalllogcmdperprofilerate,omitempty"`
	Appfirewallwellformednesslogsperprofile uint64 `json:"appfirewallwellformednesslogsperprofile,omitempty"`
	/**
	* Number of XML Format security check log messages generated by the Application Firewall.
	*/
	Appfirewallwellformednesslogsperprofilerate int32 `json:"appfirewallwellformednesslogsperprofilerate,omitempty"`
	Appfirewallxdoslogsperprofile uint64 `json:"appfirewallxdoslogsperprofile,omitempty"`
	/**
	* Number of XML Denial-of-Service security check log messages generated by the Application Firewall.
	*/
	Appfirewallxdoslogsperprofilerate int32 `json:"appfirewallxdoslogsperprofilerate,omitempty"`
	Appfirewallmsgvallogsperprofile uint64 `json:"appfirewallmsgvallogsperprofile,omitempty"`
	/**
	* Number of XML Message Validation security check log messages generated by the Application Firewall.
	*/
	Appfirewallmsgvallogsperprofilerate int32 `json:"appfirewallmsgvallogsperprofilerate,omitempty"`
	Appfirewallwsilogsperprofile uint64 `json:"appfirewallwsilogsperprofile,omitempty"`
	/**
	* Number of Web Services Interoperability (WS-I) security check log messages generated by the Application Firewall.
	*/
	Appfirewallwsilogsperprofilerate int32 `json:"appfirewallwsilogsperprofilerate,omitempty"`
	Appfirewallxmlsqllogsperprofile uint64 `json:"appfirewallxmlsqllogsperprofile,omitempty"`
	/**
	* Number of XML SQL Injection security check log messages generated by the Application Firewall.
	*/
	Appfirewallxmlsqllogsperprofilerate int32 `json:"appfirewallxmlsqllogsperprofilerate,omitempty"`
	Appfirewallxmlxsslogsperprofile uint64 `json:"appfirewallxmlxsslogsperprofile,omitempty"`
	/**
	* Number of XML Cross-Site Scripting (XSS) security check log messages generated by the Application Firewall.
	*/
	Appfirewallxmlxsslogsperprofilerate int32 `json:"appfirewallxmlxsslogsperprofilerate,omitempty"`
	Appfirewallxmlattachmentlogsperprofile uint64 `json:"appfirewallxmlattachmentlogsperprofile,omitempty"`
	/**
	* Number of XML Attachment security check log messages generated by the Application Firewall.
	*/
	Appfirewallxmlattachmentlogsperprofilerate int32 `json:"appfirewallxmlattachmentlogsperprofilerate,omitempty"`
	Appfirewallxmlsoapfaultlogsperprofile uint64 `json:"appfirewallxmlsoapfaultlogsperprofile,omitempty"`
	/**
	* Number of requests generating soap:fault log messages
	*/
	Appfirewallxmlsoapfaultlogsperprofilerate int32 `json:"appfirewallxmlsoapfaultlogsperprofilerate,omitempty"`
	Appfirewallxmlgenericlogsperprofile uint64 `json:"appfirewallxmlgenericlogsperprofile,omitempty"`
	/**
	* Number of requests generating XML Generic log messages
	*/
	Appfirewallxmlgenericlogsperprofilerate int32 `json:"appfirewallxmlgenericlogsperprofilerate,omitempty"`
	Appfirewalllogsqlgramperprofile uint64 `json:"appfirewalllogsqlgramperprofile,omitempty"`
	/**
	* Number of HTML SQL Injection security check log messages (reported by SQL grammar) generated by the Application Firewall.
	*/
	Appfirewalllogsqlgramperprofilerate int32 `json:"appfirewalllogsqlgramperprofilerate,omitempty"`
	Appfirewalllogsjsonsqlgramperprofile uint64 `json:"appfirewalllogsjsonsqlgramperprofile,omitempty"`
	/**
	* Number of JSON SQL Injection security check log messages (reported by SQL grammar) generated by the Application Firewall.
	*/
	Appfirewalllogsjsonsqlgramperprofilerate int32 `json:"appfirewalllogsjsonsqlgramperprofilerate,omitempty"`
	Appfirewalltotallogperprofile uint64 `json:"appfirewalltotallogperprofile,omitempty"`
	/**
	* Number of log messages generated by the application firewall on per profile basis
	*/
	Appfirewalllogperprofilerate int32 `json:"appfirewalllogperprofilerate,omitempty"`
	Appfirewallret4xxperprofile uint64 `json:"appfirewallret4xxperprofile,omitempty"`
	/**
	* Number of requests returning HTTP 4xx from the backend server
	*/
	Appfirewallret4xxperprofilerate int32 `json:"appfirewallret4xxperprofilerate,omitempty"`
	Appfirewallret5xxperprofile uint64 `json:"appfirewallret5xxperprofile,omitempty"`
	/**
	* Number of requests returning HTTP 5xx from the backend server
	*/
	Appfirewallret5xxperprofilerate int32 `json:"appfirewallret5xxperprofilerate,omitempty"`

}
