package goami2

func getAmiFixtureCall() []string {
	amiPack := make([]string, 0)

	// pack # 0
	amiPack = append(amiPack, "Event: DeviceStateChange\r\n"+
		"Privilege: call,all\r\n"+
		"Device: SIP/9170-12\r\n"+
		"State: INUSE\r\n\r\n")

	// pack # 1
	amiPack = append(amiPack, "Event: Newchannel\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 0\r\n"+
		"ChannelStateDesc: Down\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n\r\n")

	// pack # 2
	amiPack = append(amiPack, "Event: Newstate\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 4\r\n"+
		"ChannelStateDesc: Ring\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: fr\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n\r\n")

	// pack # 3
	amiPack = append(amiPack, "Event: SuccessfulAuth\r\n"+
		"Privilege: security,all\r\n"+
		"EventTV: 2020-08-31T11:28:01.428-0400\r\n"+
		"Severity: Informational\r\n"+
		"Service: SIP\r\n"+
		"EventVersion: 1\r\n"+
		"AccountID: 9898\r\n"+
		"SessionID: 0x7f678401d848\r\n"+
		"LocalAddress: IPV4/UDP/10.0.0.181/5060\r\n"+
		"RemoteAddress: IPV4/UDP/10.0.0.185/5060\r\n"+
		"UsingPassword: 0\r\n\r\n")

	// pack # 4
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 4\r\n"+
		"ChannelStateDesc: Ring\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: fr\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Extension: 9898\r\n"+
		"Application: NoOp\r\n"+
		"AppData: Hit dialplan for 9898 @ okon.ferry.sip.com REALM: okon.ferry.sip.com / , CALLID: b5ser03agv7huo7faai5\r\n\r\n")

	// pack # 5
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 4\r\n"+
		"ChannelStateDesc: Ring\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: fr\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 2\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Extension: 9898\r\n"+
		"Application: GotoIf\r\n"+
		"AppData: 0?continue\r\n\r\n")

	// pack # 6
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 4\r\n"+
		"ChannelStateDesc: Ring\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: fr\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 3\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Extension: 9898\r\n"+
		"Application: Set\r\n"+
		"AppData: __realm=okon.ferry.sip.com\r\n\r\n")

	// pack # 7
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 4\r\n"+
		"ChannelStateDesc: Ring\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: fr\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 4\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Extension: 9898\r\n"+
		"Application: Set\r\n"+
		"AppData: __LANGUAGE=en\r\n\r\n")

	// pack # 8
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 4\r\n"+
		"ChannelStateDesc: Ring\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: fr\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 5\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Extension: 9898\r\n"+
		"Application: GotoIf\r\n"+
		"AppData: 1?continue\r\n\r\n")

	// pack # 9
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 4\r\n"+
		"ChannelStateDesc: Ring\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: fr\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 8\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Extension: 9898\r\n"+
		"Application: Set\r\n"+
		"AppData: __SIPFROMDOMAIN=okon.ferry.sip.com\r\n\r\n")

	// pack # 10
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 4\r\n"+
		"ChannelStateDesc: Ring\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: fr\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 9\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Extension: 9898\r\n"+
		"Application: AGI\r\n"+
		"AppData: zonkey.agi,9898\r\n\r\n")

	// pack # 11
	amiPack = append(amiPack, "Event: Newchannel\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 0\r\n"+
		"ChannelStateDesc: Down\r\n"+
		"CallerIDNum: <unknown>\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: s\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n\r\n")

	// pack # 12
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 0\r\n"+
		"ChannelStateDesc: Down\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"Extension: 9898\r\n"+
		"Application: AppDial\r\n"+
		"AppData: (Outgoing Line)\r\n\r\n")

	// pack # 13
	amiPack = append(amiPack, "Event: NewCallerid\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 0\r\n"+
		"ChannelStateDesc: Down\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"CID-CallingPres: 0 (Presentation Allowed, Not Screened)\r\n\r\n")

	// pack # 14
	amiPack = append(amiPack, "Event: NewConnectedLine\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 0\r\n"+
		"ChannelStateDesc: Down\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n\r\n")

	// pack # 15
	amiPack = append(amiPack, "Event: DialBegin\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 4\r\n"+
		"ChannelStateDesc: Ring\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 9\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"DestChannel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"DestChannelState: 0\r\n"+
		"DestChannelStateDesc: Down\r\n"+
		"DestCallerIDNum: 9898\r\n"+
		"DestCallerIDName: <unknown>\r\n"+
		"DestConnectedLineNum: 9170\r\n"+
		"DestConnectedLineName: Bud Heller\r\n"+
		"DestLanguage: en\r\n"+
		"DestAccountCode: \r\n"+
		"DestContext: default\r\n"+
		"DestExten: 9898\r\n"+
		"DestPriority: 1\r\n"+
		"DestUniqueid: 1598887681.61\r\n"+
		"DestLinkedid: 1598887681.60\r\n"+
		"DestChanVariable: realm=okon.ferry.sip.com\r\n"+
		"DestChanVariable: SIPDOMAIN=\r\n"+
		"DestChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"DialString: 9898@okon.ferry.sip.com\r\n\r\n")

	// pack # 16
	amiPack = append(amiPack, "Event: Newstate\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 5\r\n"+
		"ChannelStateDesc: Ringing\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n\r\n")

	// pack # 17
	amiPack = append(amiPack, "Event: DeviceStateChange\r\n"+
		"Privilege: call,all\r\n"+
		"Device: SIP/okon.ferry.sip.com\r\n"+
		"State: INVALID\r\n\r\n")

	// pack # 18
	amiPack = append(amiPack, "Event: Newstate\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n\r\n")

	// pack # 19
	amiPack = append(amiPack, "Event: DeviceStateChange\r\n"+
		"Privilege: call,all\r\n"+
		"Device: SIP/okon.ferry.sip.com\r\n"+
		"State: INVALID\r\n\r\n")

	// pack # 20
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: macro-saveDialedChannel\r\n"+
		"Exten: s\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"Extension: s\r\n"+
		"Application: Macro\r\n"+
		"AppData: saveDialedChannel\r\n\r\n")

	// pack # 21
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: macro-saveDialedChannel\r\n"+
		"Exten: s\r\n"+
		"Priority: 2\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"Extension: s\r\n"+
		"Application: Macro\r\n"+
		"AppData: saveDialedChannel\r\n\r\n")

	// pack # 22
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: macro-saveDialedChannel\r\n"+
		"Exten: s\r\n"+
		"Priority: 3\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"Extension: s\r\n"+
		"Application: Macro\r\n"+
		"AppData: saveDialedChannel\r\n\r\n")

	// pack # 23
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: macro-saveDialedChannel\r\n"+
		"Exten: s\r\n"+
		"Priority: 4\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"Extension: s\r\n"+
		"Application: Macro\r\n"+
		"AppData: saveDialedChannel\r\n\r\n")

	// pack # 24
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: macro-saveDialedChannel\r\n"+
		"Exten: s\r\n"+
		"Priority: 5\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"Extension: s\r\n"+
		"Application: Macro\r\n"+
		"AppData: saveDialedChannel\r\n\r\n")

	// pack # 25
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: macro-saveDialedChannel\r\n"+
		"Exten: s\r\n"+
		"Priority: 6\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"Extension: s\r\n"+
		"Application: Macro\r\n"+
		"AppData: saveDialedChannel\r\n\r\n")

	// pack # 26
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: macro-saveDialedChannel\r\n"+
		"Exten: s\r\n"+
		"Priority: 8\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"Extension: s\r\n"+
		"Application: Macro\r\n"+
		"AppData: saveDialedChannel\r\n\r\n")

	// pack # 27
	amiPack = append(amiPack, "Event: UserEvent\r\n"+
		"Privilege: user,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: macro-saveDialedChannel\r\n"+
		"Exten: s\r\n"+
		"Priority: 8\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"UserEvent: CallMonitor\r\n"+
		"Recording: on\r\n"+
		"MonFile: /var/spool/asterisk/monitor/okon.ferry.sip.com/20200831112801-9656c8c56eba099fef31a8e2eb694483.wav\r\n"+
		"MonCmd: /etc/zonkey/bin/monitor_script.sh dummy1 dummy2 okon.ferry.sip.com 365 /var/spool/asterisk/monitor/okon.ferry.sip.com/20200831112801-9656c8c56eba099fef31a8e2eb694483\r\n"+
		"SipInitClid: b5ser03agv7huo7faai5\r\n\r\n")

	// pack # 28
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: macro-saveDialedChannel\r\n"+
		"Exten: s\r\n"+
		"Priority: 9\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"Extension: s\r\n"+
		"Application: Macro\r\n"+
		"AppData: saveDialedChannel\r\nEvent: DialEnd\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 4\r\n"+
		"ChannelStateDesc: Ring\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 9\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"DestChannel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"DestChannelState: 6\r\n"+
		"DestChannelStateDesc: Up\r\n"+
		"DestCallerIDNum: 9898\r\n"+
		"DestCallerIDName: <unknown>\r\n"+
		"DestConnectedLineNum: 9170\r\n"+
		"DestConnectedLineName: Bud Heller\r\n"+
		"DestLanguage: en\r\n"+
		"DestAccountCode: \r\n"+
		"DestContext: default\r\n"+
		"DestExten: 9898\r\n"+
		"DestPriority: 1\r\n"+
		"DestUniqueid: 1598887681.61\r\n"+
		"DestLinkedid: 1598887681.60\r\n"+
		"DestChanVariable: realm=okon.ferry.sip.com\r\n"+
		"DestChanVariable: SIPDOMAIN=\r\n"+
		"DestChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"DialStatus: ANSWER\r\n\r\n")

	// pack # 29
	amiPack = append(amiPack, "Event: Newstate\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 9\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n\r\n")

	// pack # 30
	amiPack = append(amiPack, "Event: BridgeCreate\r\n"+
		"Privilege: call,all\r\n"+
		"BridgeUniqueid: 208db8a2-3e19-49b4-96df-f27944815f52\r\n"+
		"BridgeType: basic\r\n"+
		"BridgeTechnology: simple_bridge\r\n"+
		"BridgeCreator: <unknown>\r\n"+
		"BridgeName: <unknown>\r\n"+
		"BridgeNumChannels: 0\r\n"+
		"BridgeVideoSourceMode: none\r\n\r\n")

	// pack # 31
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"Extension: 9898\r\n"+
		"Application: AppDial\r\n"+
		"AppData: (Outgoing Line)\r\n\r\n")

	// pack # 32
	amiPack = append(amiPack, "Event: BridgeEnter\r\n"+
		"Privilege: call,all\r\n"+
		"BridgeUniqueid: 208db8a2-3e19-49b4-96df-f27944815f52\r\n"+
		"BridgeType: basic\r\n"+
		"BridgeTechnology: simple_bridge\r\n"+
		"BridgeCreator: <unknown>\r\n"+
		"BridgeName: <unknown>\r\n"+
		"BridgeNumChannels: 1\r\n"+
		"BridgeVideoSourceMode: none\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n\r\n")

	// pack # 33
	amiPack = append(amiPack, "Event: BridgeEnter\r\n"+
		"Privilege: call,all\r\n"+
		"BridgeUniqueid: 208db8a2-3e19-49b4-96df-f27944815f52\r\n"+
		"BridgeType: basic\r\n"+
		"BridgeTechnology: simple_bridge\r\n"+
		"BridgeCreator: <unknown>\r\n"+
		"BridgeName: <unknown>\r\n"+
		"BridgeNumChannels: 2\r\n"+
		"BridgeVideoSourceMode: none\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 9\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n\r\n")

	// pack # 34
	amiPack = append(amiPack, "Event: HangupRequest\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 9\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n\r\n")

	// pack # 35
	amiPack = append(amiPack, "Event: BridgeLeave\r\n"+
		"Privilege: call,all\r\n"+
		"BridgeUniqueid: 208db8a2-3e19-49b4-96df-f27944815f52\r\n"+
		"BridgeType: basic\r\n"+
		"BridgeTechnology: simple_bridge\r\n"+
		"BridgeCreator: <unknown>\r\n"+
		"BridgeName: <unknown>\r\n"+
		"BridgeNumChannels: 1\r\n"+
		"BridgeVideoSourceMode: none\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 9\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n\r\n")

	// pack # 36
	amiPack = append(amiPack, "Event: BridgeLeave\r\n"+
		"Privilege: call,all\r\n"+
		"BridgeUniqueid: 208db8a2-3e19-49b4-96df-f27944815f52\r\n"+
		"BridgeType: basic\r\n"+
		"BridgeTechnology: simple_bridge\r\n"+
		"BridgeCreator: <unknown>\r\n"+
		"BridgeName: <unknown>\r\n"+
		"BridgeNumChannels: 0\r\n"+
		"BridgeVideoSourceMode: none\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n\r\n")

	// pack # 37
	amiPack = append(amiPack, "Event: BridgeDestroy\r\n"+
		"Privilege: call,all\r\n"+
		"BridgeUniqueid: 208db8a2-3e19-49b4-96df-f27944815f52\r\n"+
		"BridgeType: basic\r\n"+
		"BridgeTechnology: simple_bridge\r\n"+
		"BridgeCreator: <unknown>\r\n"+
		"BridgeName: <unknown>\r\n"+
		"BridgeNumChannels: 0\r\n"+
		"BridgeVideoSourceMode: none\r\n\r\n")

	// pack # 38
	amiPack = append(amiPack, "Event: Hangup\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000003d\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9898\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.61\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=3aff7b257f578f093ed98d6c45dd1514@10.0.0.181:5060\r\n"+
		"Cause: 16\r\n"+
		"Cause-txt: Normal Clearing\r\n\r\n")

	// pack # 39
	amiPack = append(amiPack, "Event: DeviceStateChange\r\n"+
		"Privilege: call,all\r\n"+
		"Device: SIP/okon.ferry.sip.com\r\n"+
		"State: INVALID\r\n\r\n")

	// pack # 40
	amiPack = append(amiPack, "Event: SoftHangupRequest\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: 9898\r\n"+
		"Priority: 9\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Cause: 16\r\n\r\n")

	// pack # 41
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: h\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Extension: h\r\n"+
		"Application: NoOp\r\n"+
		"AppData: default hangup hit\r\n\r\n")

	// pack # 42
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: h\r\n"+
		"Priority: 2\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Extension: h\r\n"+
		"Application: Set\r\n"+
		"AppData: count=\r\n\r\n")

	// pack # 43
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: h\r\n"+
		"Priority: 3\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Extension: h\r\n"+
		"Application: Set\r\n"+
		"AppData: CDR(dialstatus)=ANSWER\r\n\r\n")

	// pack # 44
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: h\r\n"+
		"Priority: 4\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Extension: h\r\n"+
		"Application: GotoIf\r\n"+
		"AppData: 1?billing\r\nEvent: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: h\r\n"+
		"Priority: 7\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Extension: h\r\n"+
		"Application: AGI\r\n"+
		"AppData: zonkey.agi,billing\r\n\r\n")

	// pack # 45
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: h\r\n"+
		"Priority: 8\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Extension: h\r\n"+
		"Application: GotoIf\r\n"+
		"AppData: 0?vmcallback:terminate\r\n\r\n")

	// pack # 46
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: h\r\n"+
		"Priority: 12\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Extension: h\r\n"+
		"Application: Congestion\r\n"+
		"AppData: \r\n\r\n")

	// pack # 47
	amiPack = append(amiPack, "Event: DeviceStateChange\r\n"+
		"Privilege: call,all\r\n"+
		"Device: SIP/9170-12\r\n"+
		"State: NOT_INUSE\r\n\r\n")

	// pack # 48
	amiPack = append(amiPack, "Event: Hangup\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/9170-12-0000003c\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: h\r\n"+
		"Priority: 12\r\n"+
		"Uniqueid: 1598887681.60\r\n"+
		"Linkedid: 1598887681.60\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n"+
		"Cause: 16\r\n"+
		"Cause-txt: Normal Clearing\r\n\r\n")

	return amiPack
}

func getAmiFixtureQueues() []string {
	amiPack := make([]string, 0)

	// pack # 0
	amiPack = append(amiPack, "Action: Command\r\n"+
		"Command: queue show\r\n"+
		"Amictx: queue\r\n"+
		"Actionid: bef2eb20211b9@okon.ferry.sip.com\r\n\r\n")

	// pack # 1
	amiPack = append(amiPack, "Action: QueueStatus\r\n"+
		"Queue: Shipping_12\r\n"+
		"Actionid: 175a511275bae@okon.ferry.sip.com\r\n\r\n")

	// pack # 2
	amiPack = append(amiPack, "Action: QueueStatus\r\n"+
		"Actionid: e87c0c83a9c4c@okon.ferry.sip.com\r\n"+
		"Queue: Books_12\r\n\r\n")

	// pack # 3
	amiPack = append(amiPack, "Action: QueueStatus\r\n"+
		"Actionid: b8b7528f22e16@okon.ferry.sip.com\r\n"+
		"Queue: Sales_12\r\n\r\n")

	// pack # 4
	amiPack = append(amiPack, "Response: Follows\r\n"+
		"Privilege: Command\r\n"+
		"ActionID: bef2eb20211b9@okon.ferry.sip.com\r\n\r\n")

	// pack # 5
	amiPack = append(amiPack, "Action: QueueStatus\r\n"+
		"Queue: Financing_12\r\n"+
		"Actionid: d3dbdf0621528@okon.ferry.sip.com\r\n\r\n")

	// pack # 6
	amiPack = append(amiPack, "Action: QueueStatus\r\n"+
		"Queue: ISupport02_12\r\n"+
		"Actionid: 109418f7f141f@okon.ferry.sip.com\r\n\r\n")

	// pack # 7
	amiPack = append(amiPack, "Action: QueueStatus\r\n"+
		"Queue: Bar_12\r\n"+
		"Actionid: ff56787833bc5@okon.ferry.sip.com\r\n\r\n")

	// pack # 8
	amiPack = append(amiPack, "ISupport02_12 has 0 calls (max unlimited) in 'ringall' strategy (0s holdtime, 0s talktime), W:0, C:0, A:0, SL:0.0% within 0s\n"+
		"	Members:\n"+
		"		7855@okon.ferry.sip.com (Local/7855-12@from-queue/n from SIP/7855-12) (ringinuse disabled) (realtime) (paused) (Not in use) has taken no calls yet\n"+
		"		7853@okon.ferry.sip.com (Local/7853-12@from-queue/n from SIP/7853-12) (ringinuse enabled) (realtime) (paused) (Not in use) has taken no calls yet\n"+
		"	No Callers\r\n"+
		"Sales_12 has 0 calls (max unlimited) in 'ringall' strategy (0s holdtime, 0s talktime), W:0, C:0, A:0, SL:0.0% within 0s\n"+
		"	No Members\n"+
		"	No Callers\r\n"+
		"Financing_12 has 0 calls (max unlimited) in 'ringall' strategy (0s holdtime, 0s talktime), W:0, C:0, A:0, SL:0.0% within 0s\n"+
		"	Members:"+
		"		8895@okon.ferry.sip.com (Local/8895-12@from-queue/n from SIP/8895-12) (ringinuse disabled) (realtime) (Not in use) has taken no calls yet\n"+
		"		9068@okon.ferry.sip.com (Local/9068-12@from-queue/n from SIP/9068-12) (ringinuse disabled) (realtime) (Not in use) has taken no calls yet\n"+
		"	No Callers\r\n"+
		"Shipping_12 has 0 calls (max unlimited) in 'ringall' strategy (0s holdtime, 0s talktime), W:0, C:0, A:0, SL:0.0% within 0s\n"+
		"	Members:"+
		"		9170@okon.ferry.sip.com (Local/9170-12@from-queue/n from SIP/9170-12) (ringinuse disabled) (realtime) (Not in use) has taken no calls yet\n"+
		"		8194@okon.ferry.sip.com (Local/8194-12@from-queue/n from SIP/8194-12) (ringinuse disabled) (realtime) (Not in use) has taken no calls yet\n"+
		"	No Callers\r\n"+
		"Bar_12 has 0 calls (max unlimited) in 'ringall' strategy (0s holdtime, 0s talktime), W:0, C:0, A:0, SL:0.0% within 0s\n"+
		"	Members:\n"+
		"		1245@okon.ferry.sip.com (Local/1245-12@from-queue/n from SIP/1245-12) (ringinuse enabled) (realtime) (paused) (Not in use) has taken no calls yet\n"+
		"	No Callers\r\n"+
		"Books_12 has 0 calls (max unlimited) in 'ringall' strategy (0s holdtime, 0s talktime), W:0, C:0, A:0, SL:0.0% within 0s\n"+
		"	Members:\n"+
		"		9921@okon.ferry.sip.com (Local/9921-12@from-queue/n from SIP/9921-12) (ringinuse disabled) (realtime) (Not in use) has taken no calls yet\n"+
		"		9922@okon.ferry.sip.com (Local/9922-12@from-queue/n from SIP/9922-12) (ringinuse disabled) (realtime) (Not in use) has taken no calls yet\n"+
		"	No Callers\r\n"+
		"--END COMMAND--\r\n\r\n")

	// pack # 9
	amiPack = append(amiPack, "Response: Success\r\n"+
		"ActionID: 175a511275bae@okon.ferry.sip.com\r\n"+
		"EventList: start\r\n"+
		"Message: Queue status will follow\r\n\r\n")

	// pack # 10
	amiPack = append(amiPack, "Event: QueueParams\r\n"+
		"Queue: Shipping_12\r\n"+
		"Max: 0\r\n"+
		"Strategy: ringall\r\n"+
		"Calls: 0\r\n"+
		"Holdtime: 0\r\n"+
		"TalkTime: 0\r\n"+
		"Completed: 0\r\n"+
		"Abandoned: 0\r\n"+
		"ServiceLevel: 0\r\n"+
		"ServicelevelPerf: 0.0\r\n"+
		"Weight: 0\r\n"+
		"ActionID: 175a511275bae@okon.ferry.sip.com\r\n\r\n")

	// pack # 11
	amiPack = append(amiPack, "Event: QueueMember\r\n"+
		"Queue: Shipping_12\r\n"+
		"Name: 9170@okon.ferry.sip.com\r\n"+
		"Location: Local/9170-12@from-queue/n\r\n"+
		"StateInterface: SIP/9170-12\r\n"+
		"Membership: static\r\n"+
		"Penalty: 0\r\n"+
		"CallsTaken: 0\r\n"+
		"LastCall: 0\r\n"+
		"InCall: 0\r\n"+
		"Status: 1\r\n"+
		"Paused: 0\r\n"+
		"PausedReason: \r\n"+
		"ActionID: 175a511275bae@okon.ferry.sip.com\r\nEvent: QueueMember\r\n"+
		"Queue: Shipping_12\r\n"+
		"Name: 8194@okon.ferry.sip.com\r\n"+
		"Location: Local/8194-12@from-queue/n\r\n"+
		"StateInterface: SIP/8194-12\r\n"+
		"Membership: static\r\n"+
		"Penalty: 0\r\n"+
		"CallsTaken: 0\r\n"+
		"LastCall: 0\r\n"+
		"InCall: 0\r\n"+
		"Status: 1\r\n"+
		"Paused: 0\r\n"+
		"PausedReason: \r\n"+
		"ActionID: 175a511275bae@okon.ferry.sip.com\r\n\r\n")

	// pack # 12
	amiPack = append(amiPack, "Event: QueueStatusComplete\r\n"+
		"ActionID: 175a511275bae@okon.ferry.sip.com\r\n"+
		"EventList: Complete\r\n"+
		"ListItems: 3\r\nResponse: Success\r\n"+
		"ActionID: e87c0c83a9c4c@okon.ferry.sip.com\r\n"+
		"EventList: start\r\n"+
		"Message: Queue status will follow\r\n\r\n")

	// pack # 13
	amiPack = append(amiPack, "Event: QueueParams\r\n"+
		"Queue: Books_12\r\n"+
		"Max: 0\r\n"+
		"Strategy: ringall\r\n"+
		"Calls: 0\r\n"+
		"Holdtime: 0\r\n"+
		"TalkTime: 0\r\n"+
		"Completed: 0\r\n"+
		"Abandoned: 0\r\n"+
		"ServiceLevel: 0\r\n"+
		"ServicelevelPerf: 0.0\r\n"+
		"Weight: 0\r\n"+
		"ActionID: e87c0c83a9c4c@okon.ferry.sip.com\r\n\r\n")

	// pack # 14
	amiPack = append(amiPack, "Event: QueueMember\r\n"+
		"Queue: Books_12\r\n"+
		"Name: 9921@okon.ferry.sip.com\r\n"+
		"Location: Local/9921-12@from-queue/n\r\n"+
		"StateInterface: SIP/9921-12\r\n"+
		"Membership: static\r\n"+
		"Penalty: 0\r\n"+
		"CallsTaken: 0\r\n"+
		"LastCall: 0\r\n"+
		"InCall: 0\r\n"+
		"Status: 1\r\n"+
		"Paused: 0\r\n"+
		"PausedReason: \r\n"+
		"ActionID: e87c0c83a9c4c@okon.ferry.sip.com\r\n\r\n")

	// pack # 15
	amiPack = append(amiPack, "Event: QueueMember\r\n"+
		"Queue: Books_12\r\n"+
		"Name: 9922@okon.ferry.sip.com\r\n"+
		"Location: Local/9922-12@from-queue/n\r\n"+
		"StateInterface: SIP/9922-12\r\n"+
		"Membership: static\r\n"+
		"Penalty: 0\r\n"+
		"CallsTaken: 0\r\n"+
		"LastCall: 0\r\n"+
		"InCall: 0\r\n"+
		"Status: 1\r\n"+
		"Paused: 0\r\n"+
		"PausedReason: \r\n"+
		"ActionID: e87c0c83a9c4c@okon.ferry.sip.com\r\n\r\n")

	// pack # 16
	amiPack = append(amiPack, "Event: QueueStatusComplete\r\n"+
		"ActionID: e87c0c83a9c4c@okon.ferry.sip.com\r\n"+
		"EventList: Complete\r\n"+
		"ListItems: 3\r\n\r\n")

	// pack # 17
	amiPack = append(amiPack, "Response: Success\r\n"+
		"ActionID: b8b7528f22e16@okon.ferry.sip.com\r\n"+
		"EventList: start\r\n"+
		"Message: Queue status will follow\r\n\r\n")

	// pack # 18
	amiPack = append(amiPack, "Event: QueueParams\r\n"+
		"Queue: Sales_12\r\n"+
		"Max: 0\r\n"+
		"Strategy: ringall\r\n"+
		"Calls: 0\r\n"+
		"Holdtime: 0\r\n"+
		"TalkTime: 0\r\n"+
		"Completed: 0\r\n"+
		"Abandoned: 0\r\n"+
		"ServiceLevel: 0\r\n"+
		"ServicelevelPerf: 0.0\r\n"+
		"Weight: 0\r\n"+
		"ActionID: b8b7528f22e16@okon.ferry.sip.com\r\n\r\n")

	// pack # 19
	amiPack = append(amiPack, "Event: QueueStatusComplete\r\n"+
		"ActionID: b8b7528f22e16@okon.ferry.sip.com\r\n"+
		"EventList: Complete\r\n"+
		"ListItems: 1\r\n\r\n")

	// pack # 20
	amiPack = append(amiPack, "Response: Success\r\n"+
		"ActionID: d3dbdf0621528@okon.ferry.sip.com\r\n"+
		"EventList: start\r\n"+
		"Message: Queue status will follow\r\n\r\n")

	// pack # 21
	amiPack = append(amiPack, "Event: QueueParams\r\n"+
		"Queue: Financing_12\r\n"+
		"Max: 0\r\n"+
		"Strategy: ringall\r\n"+
		"Calls: 0\r\n"+
		"Holdtime: 0\r\n"+
		"TalkTime: 0\r\n"+
		"Completed: 0\r\n"+
		"Abandoned: 0\r\n"+
		"ServiceLevel: 0\r\n"+
		"ServicelevelPerf: 0.0\r\n"+
		"Weight: 0\r\n"+
		"ActionID: d3dbdf0621528@okon.ferry.sip.com\r\n\r\n")

	// pack # 22
	amiPack = append(amiPack, "Event: QueueMember\r\n"+
		"Queue: Financing_12\r\n"+
		"Name: 8895@okon.ferry.sip.com\r\n"+
		"Location: Local/8895-12@from-queue/n\r\n"+
		"StateInterface: SIP/8895-12\r\n"+
		"Membership: static\r\n"+
		"Penalty: 0\r\n"+
		"CallsTaken: 0\r\n"+
		"LastCall: 0\r\n"+
		"InCall: 0\r\n"+
		"Status: 1\r\n"+
		"Paused: 0\r\n"+
		"PausedReason: \r\n"+
		"ActionID: d3dbdf0621528@okon.ferry.sip.com\r\n\r\n")

	// pack # 23
	amiPack = append(amiPack, "Event: QueueMember\r\n"+
		"Queue: Financing_12\r\n"+
		"Name: 9068@okon.ferry.sip.com\r\n"+
		"Location: Local/9068-12@from-queue/n\r\n"+
		"StateInterface: SIP/9068-12\r\n"+
		"Membership: static\r\n"+
		"Penalty: 0\r\n"+
		"CallsTaken: 0\r\n"+
		"LastCall: 0\r\n"+
		"InCall: 0\r\n"+
		"Status: 1\r\n"+
		"Paused: 0\r\n"+
		"PausedReason: \r\n"+
		"ActionID: d3dbdf0621528@okon.ferry.sip.com\r\n\r\n")

	// pack # 24
	amiPack = append(amiPack, "Event: QueueStatusComplete\r\n"+
		"ActionID: d3dbdf0621528@okon.ferry.sip.com\r\n"+
		"EventList: Complete\r\n"+
		"ListItems: 3\r\n\r\n")

	// pack # 25
	amiPack = append(amiPack, "Response: Success\r\n"+
		"ActionID: 109418f7f141f@okon.ferry.sip.com\r\n"+
		"EventList: start\r\n"+
		"Message: Queue status will follow\r\n\r\n")

	// pack # 26
	amiPack = append(amiPack, "Event: QueueParams\r\n"+
		"Queue: ISupport02_12\r\n"+
		"Max: 0\r\n"+
		"Strategy: ringall\r\n"+
		"Calls: 0\r\n"+
		"Holdtime: 0\r\n"+
		"TalkTime: 0\r\n"+
		"Completed: 0\r\n"+
		"Abandoned: 0\r\n"+
		"ServiceLevel: 0\r\n"+
		"ServicelevelPerf: 0.0\r\n"+
		"Weight: 0\r\n"+
		"ActionID: 109418f7f141f@okon.ferry.sip.com\r\nEvent: QueueMember\r\n"+
		"Queue: ISupport02_12\r\n"+
		"Name: 7855@okon.ferry.sip.com\r\n"+
		"Location: Local/7855-12@from-queue/n\r\n"+
		"StateInterface: SIP/7855-12\r\n"+
		"Membership: static\r\n"+
		"Penalty: 0\r\n"+
		"CallsTaken: 0\r\n"+
		"LastCall: 0\r\n"+
		"InCall: 0\r\n"+
		"Status: 1\r\n"+
		"Paused: 1\r\n"+
		"PausedReason: \r\n"+
		"ActionID: 109418f7f141f@okon.ferry.sip.com\r\n\r\n")

	// pack # 27
	amiPack = append(amiPack, "Event: QueueMember\r\n"+
		"Queue: ISupport02_12\r\n"+
		"Name: 7853@okon.ferry.sip.com\r\n"+
		"Location: Local/7853-12@from-queue/n\r\n"+
		"StateInterface: SIP/7853-12\r\n"+
		"Membership: static\r\n"+
		"Penalty: 0\r\n"+
		"CallsTaken: 0\r\n"+
		"LastCall: 0\r\n"+
		"InCall: 0\r\n"+
		"Status: 1\r\n"+
		"Paused: 1\r\n"+
		"PausedReason: \r\n"+
		"ActionID: 109418f7f141f@okon.ferry.sip.com\r\n\r\n")

	// pack # 28
	amiPack = append(amiPack, "Event: QueueStatusComplete\r\n"+
		"ActionID: 109418f7f141f@okon.ferry.sip.com\r\n"+
		"EventList: Complete\r\n"+
		"ListItems: 3\r\n\r\n")

	// pack # 29
	amiPack = append(amiPack, "Response: Success\r\n"+
		"ActionID: ff56787833bc5@okon.ferry.sip.com\r\n"+
		"EventList: start\r\n"+
		"Message: Queue status will follow\r\n\r\n")

	// pack # 30
	amiPack = append(amiPack, "Event: QueueParams\r\n"+
		"Queue: Bar_12\r\n"+
		"Max: 0\r\n"+
		"Strategy: ringall\r\n"+
		"Calls: 0\r\n"+
		"Holdtime: 0\r\n"+
		"TalkTime: 0\r\n"+
		"Completed: 0\r\n"+
		"Abandoned: 0\r\n"+
		"ServiceLevel: 0\r\n"+
		"ServicelevelPerf: 0.0\r\n"+
		"Weight: 0\r\n"+
		"ActionID: ff56787833bc5@okon.ferry.sip.com\r\n\r\n")

	// pack # 31
	amiPack = append(amiPack, "Event: QueueMember\r\n"+
		"Queue: Bar_12\r\n"+
		"Name: 1245@okon.ferry.sip.com\r\n"+
		"Location: Local/1245-12@from-queue/n\r\n"+
		"StateInterface: SIP/1245-12\r\n"+
		"Membership: static\r\n"+
		"Penalty: 0\r\n"+
		"CallsTaken: 0\r\n"+
		"LastCall: 0\r\n"+
		"InCall: 0\r\n"+
		"Status: 1\r\n"+
		"Paused: 1\r\n"+
		"PausedReason: \r\n"+
		"ActionID: ff56787833bc5@okon.ferry.sip.com\r\n\r\n")

	// pack # 32
	amiPack = append(amiPack, "Event: QueueStatusComplete\r\n"+
		"ActionID: ff56787833bc5@okon.ferry.sip.com\r\n"+
		"EventList: Complete\r\n"+
		"ListItems: 2\r\n\r\n")

	// pack # 33
	amiPack = append(amiPack, "Action: ConfbridgeList\r\n"+
		"Conference: 20201*12\r\n"+
		"Actionid: 8870f2cf65cb3@okon.ferry.sip.com\r\n\r\n")

	// pack # 34
	amiPack = append(amiPack, "Response: Error\r\n"+
		"ActionID: 8870f2cf65cb3@okon.ferry.sip.com\r\n"+
		"Message: No active conferences\r\n\r\n")

	// pack # 35
	amiPack = append(amiPack, "Conference: 20202*12\r\n"+
		"Actionid: c0b6d56f20c8d@okon.ferry.sip.com\r\n"+
		"Action: ConfbridgeList\r\n\r\n")

	// pack # 36
	amiPack = append(amiPack, "Action: ConfbridgeList\r\n"+
		"Conference: 987*12\r\n"+
		"Actionid: 710ce47395cf@okon.ferry.sip.com\r\n\r\n")

	// pack # 37
	amiPack = append(amiPack, "Action: ConfbridgeList\r\n"+
		"Conference: 998*12\r\n"+
		"Actionid: 54df01cacad02@okon.ferry.sip.com\r\n\r\n")

	// pack # 38
	amiPack = append(amiPack, "Action: ConfbridgeList\r\n"+
		"Conference: 5555*12\r\n"+
		"Actionid: e223098fe1fd9@okon.ferry.sip.com\r\n\r\n")

	// pack # 39
	amiPack = append(amiPack, "Response: Error\r\n"+
		"ActionID: c0b6d56f20c8d@okon.ferry.sip.com\r\n"+
		"Message: No active conferences\r\n\r\n")

	// pack # 40
	amiPack = append(amiPack, "Response: Error\r\n"+
		"ActionID: 710ce47395cf@okon.ferry.sip.com\r\n"+
		"Message: No active conferences\r\n\r\n")

	// pack # 41
	amiPack = append(amiPack, "Response: Error\r\n"+
		"ActionID: 54df01cacad02@okon.ferry.sip.com\r\n"+
		"Message: No active conferences\r\n\r\n")

	// pack # 42
	amiPack = append(amiPack, "Response: Error\r\n"+
		"ActionID: e223098fe1fd9@okon.ferry.sip.com\r\n"+
		"Message: No active conferences\r\n\r\n")

	return amiPack
}

func getAmiFixtureQueueMembers() []string {
	amiPack := make([]string, 0)

	// pack # 0
	amiPack = append(amiPack, "Actionid: 419e32ed9b2ab@okon.ferry.sip.com\r\n"+
		"Action: QueuePause\r\n"+
		"Interface: Local/9921-12@from-queue/n\r\n"+
		"Paused: 0\r\n"+
		"Queue: Books_12\r\n"+
		"Reason: <nil>\r\n\r\n")

	// pack # 1
	amiPack = append(amiPack, "Response: Success\r\n"+
		"ActionID: 419e32ed9b2ab@okon.ferry.sip.com\r\n"+
		"Message: Interface unpaused successfully\r\n\r\n")

	// pack # 2
	amiPack = append(amiPack, "Event: DeviceStateChange\r\n"+
		"Privilege: call,all\r\n"+
		"Device: Queue:Books_12_pause_Local/9921-12@from-queue/n\r\n"+
		"State: NOT_INUSE\r\n\r\n")

	// pack # 3
	amiPack = append(amiPack, "Event: DeviceStateChange\r\n"+
		"Privilege: call,all\r\n"+
		"Device: Queue:Books_12_avail\r\n"+
		"State: NOT_INUSE\r\n\r\n")

	// pack # 4
	amiPack = append(amiPack, "Event: QueueMemberPause\r\n"+
		"Privilege: agent,all\r\n"+
		"Queue: Books_12\r\n"+
		"MemberName: 9921@okon.ferry.sip.com\r\n"+
		"Interface: Local/9921-12@from-queue/n\r\n"+
		"StateInterface: SIP/9921-12\r\n"+
		"Membership: realtime\r\n"+
		"Penalty: 0\r\n"+
		"CallsTaken: 0\r\n"+
		"LastCall: 0\r\n"+
		"InCall: 0\r\n"+
		"Status: 1\r\n"+
		"Paused: 0\r\n"+
		"PausedReason: \r\n"+
		"Ringinuse: 0\r\n"+
		"Reason: <nil>\r\n\r\n")

	// pack # 5
	amiPack = append(amiPack, "Reason: Break\r\n"+
		"Actionid: 647f57eeccf34@okon.ferry.sip.com\r\n"+
		"Action: QueuePause\r\n"+
		"Interface: Local/9921-12@from-queue/n\r\n"+
		"Paused: 1\r\n"+
		"Queue: Books_12\r\n\r\n")

	// pack # 6
	amiPack = append(amiPack, "Response: Success\r\n"+
		"ActionID: 647f57eeccf34@okon.ferry.sip.com\r\n"+
		"Message: Interface paused successfully\r\n\r\n")

	// pack # 7
	amiPack = append(amiPack, "Event: DeviceStateChange\r\n"+
		"Privilege: call,all\r\n"+
		"Device: Queue:Books_12_pause_Local/9921-12@from-queue/n\r\n"+
		"State: INUSE\r\n\r\n")

	// pack # 8
	amiPack = append(amiPack, "Event: QueueMemberPause\r\n"+
		"Privilege: agent,all\r\n"+
		"Queue: Books_12\r\n"+
		"MemberName: 9921@okon.ferry.sip.com\r\n"+
		"Interface: Local/9921-12@from-queue/n\r\n"+
		"StateInterface: SIP/9921-12\r\n"+
		"Membership: realtime\r\n"+
		"Penalty: 0\r\n"+
		"CallsTaken: 0\r\n"+
		"LastCall: 0\r\n"+
		"InCall: 0\r\n"+
		"Status: 1\r\n"+
		"Paused: 1\r\n"+
		"PausedReason: Break\r\n"+
		"Ringinuse: 0\r\n"+
		"Reason: Break\r\n\r\n")
	return amiPack
}

func getAmiFixtureConfBridge() []string {
	amiPack := make([]string, 0)

	// pack # 0
	amiPack = append(amiPack, "Response: Success\r\n"+
		"ActionID: 9c923c4dadb37@okon.ferry.sip.com\r\n"+
		"Message: Originate successfully queued\r\n\r\n")

	// pack # 1
	amiPack = append(amiPack, "Event: Newchannel\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;1\r\n"+
		"ChannelState: 0\r\n"+
		"ChannelStateDesc: Down\r\n"+
		"CallerIDNum: <unknown>\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: confbridge-invite\r\n"+
		"Exten: 7711\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598968093.901\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\nEvent: Newchannel\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;2\r\n"+
		"ChannelState: 4\r\n"+
		"ChannelStateDesc: Ring\r\n"+
		"CallerIDNum: <unknown>\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: confbridge-invite\r\n"+
		"Exten: 7711\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598968093.902\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n\r\n")

	// pack # 2
	amiPack = append(amiPack, "Event: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;1\r\n"+
		"ChannelState: 0\r\n"+
		"ChannelStateDesc: Down\r\n"+
		"CallerIDNum: <unknown>\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: confbridge-invite\r\n"+
		"Exten: 7711\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598968093.901\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n"+
		"Extension: 7711\r\n"+
		"Application: AppDial2\r\n"+
		"AppData: (Outgoing Line)\r\n\r\n")

	// pack # 3
	amiPack = append(amiPack, "Event: LocalBridge\r\n"+
		"Privilege: call,all\r\n"+
		"LocalOneChannel: Local/7711@confbridge-invite-0000000d;1\r\n"+
		"LocalOneChannelState: 0\r\n"+
		"LocalOneChannelStateDesc: Down\r\n"+
		"LocalOneCallerIDNum: 9170\r\n"+
		"LocalOneCallerIDName: Bud Heller\r\n"+
		"LocalOneConnectedLineNum: 9170\r\n"+
		"LocalOneConnectedLineName: Bud Heller\r\n"+
		"LocalOneLanguage: en\r\n"+
		"LocalOneAccountCode: \r\n"+
		"LocalOneContext: confbridge-invite\r\n"+
		"LocalOneExten: 7711\r\n"+
		"LocalOnePriority: 1\r\n"+
		"LocalOneUniqueid: 1598968093.901\r\n"+
		"LocalOneLinkedid: 1598968093.901\r\n"+
		"LocalOneChanVariable: realm=okon.ferry.sip.com\r\n"+
		"LocalOneChanVariable: SIPDOMAIN=\r\n"+
		"LocalOneChanVariable: SIPCALLID=\r\n"+
		"LocalTwoChannel: Local/7711@confbridge-invite-0000000d;2\r\n"+
		"LocalTwoChannelState: 4\r\n"+
		"LocalTwoChannelStateDesc: Ring\r\n"+
		"LocalTwoCallerIDNum: 9170\r\n"+
		"LocalTwoCallerIDName: Bud Heller\r\n"+
		"LocalTwoConnectedLineNum: 9170\r\n"+
		"LocalTwoConnectedLineName: Bud Heller\r\n"+
		"LocalTwoLanguage: en\r\n"+
		"LocalTwoAccountCode: \r\n"+
		"LocalTwoContext: confbridge-invite\r\n"+
		"LocalTwoExten: 7711\r\n"+
		"LocalTwoPriority: 1\r\n"+
		"LocalTwoUniqueid: 1598968093.902\r\n"+
		"LocalTwoLinkedid: 1598968093.901\r\n"+
		"LocalTwoChanVariable: realm=okon.ferry.sip.com\r\n"+
		"LocalTwoChanVariable: SIPDOMAIN=\r\n"+
		"LocalTwoChanVariable: SIPCALLID=\r\n"+
		"Context: confbridge-invite\r\n"+
		"Exten: 7711\r\n"+
		"LocalOptimization: Yes\r\n\r\n")

	// pack # 4
	amiPack = append(amiPack, "Event: DialBegin\r\n"+
		"Privilege: call,all\r\n"+
		"DestChannel: Local/7711@confbridge-invite-0000000d;1\r\n"+
		"DestChannelState: 0\r\n"+
		"DestChannelStateDesc: Down\r\n"+
		"DestCallerIDNum: 9170\r\n"+
		"DestCallerIDName: Bud Heller\r\n"+
		"DestConnectedLineNum: 9170\r\n"+
		"DestConnectedLineName: Bud Heller\r\n"+
		"DestLanguage: en\r\n"+
		"DestAccountCode: \r\n"+
		"DestContext: confbridge-invite\r\n"+
		"DestExten: 7711\r\n"+
		"DestPriority: 1\r\n"+
		"DestUniqueid: 1598968093.901\r\n"+
		"DestLinkedid: 1598968093.901\r\n"+
		"DestChanVariable: realm=okon.ferry.sip.com\r\n"+
		"DestChanVariable: SIPDOMAIN=\r\n"+
		"DestChanVariable: SIPCALLID=\r\n"+
		"DialString: 7711@confbridge-invite\r\nEvent: Newexten\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;2\r\n"+
		"ChannelState: 4\r\n"+
		"ChannelStateDesc: Ring\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: confbridge-invite\r\n"+
		"Exten: 7711\r\n"+
		"Priority: 3\r\n"+
		"Uniqueid: 1598968093.902\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n"+
		"Extension: 7711\r\n"+
		"Application: Set\r\n"+
		"AppData: __SIPFROMDOMAIN=okon.ferry.sip.com\r\n\r\n")

	// pack # 5
	amiPack = append(amiPack, "Event: Newchannel\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000036b\r\n"+
		"ChannelState: 0\r\n"+
		"ChannelStateDesc: Down\r\n"+
		"CallerIDNum: <unknown>\r\n"+
		"CallerIDName: <unknown>\r\n"+
		"ConnectedLineNum: <unknown>\r\n"+
		"ConnectedLineName: <unknown>\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: default\r\n"+
		"Exten: s\r\n"+
		"Priority: 1\r\n"+
		"Uniqueid: 1598968093.903\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n\r\n")

	// pack # 6
	amiPack = append(amiPack, "Event: DeviceStateChange\r\n"+
		"Privilege: call,all\r\n"+
		"Device: SIP/okon.ferry.sip.com\r\n"+
		"State: INVALID\r\n\r\n")

	// pack # 7
	amiPack = append(amiPack, "Event: DialEnd\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;2\r\n"+
		"ChannelState: 4\r\n"+
		"ChannelStateDesc: Ring\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: confbridge-invite\r\n"+
		"Exten: 7711\r\n"+
		"Priority: 4\r\n"+
		"Uniqueid: 1598968093.902\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n"+
		"DestChannel: SIP/okon.ferry.sip.com-0000036b\r\n"+
		"DestChannelState: 6\r\n"+
		"DestChannelStateDesc: Up\r\n"+
		"DestCallerIDNum: 7711\r\n"+
		"DestCallerIDName: Bud Heller\r\n"+
		"DestConnectedLineNum: 9170\r\n"+
		"DestConnectedLineName: Bud Heller\r\n"+
		"DestLanguage: en\r\n"+
		"DestAccountCode: \r\n"+
		"DestContext: confbridge-invite\r\n"+
		"DestExten: 7711\r\n"+
		"DestPriority: 1\r\n"+
		"DestUniqueid: 1598968093.903\r\n"+
		"DestLinkedid: 1598968093.901\r\n"+
		"DestChanVariable: realm=okon.ferry.sip.com\r\n"+
		"DestChanVariable: SIPDOMAIN=\r\n"+
		"DestChanVariable: SIPCALLID=0953f094680db5a312fc686902360965@10.0.0.184:5060\r\n"+
		"DialStatus: ANSWER\r\n\r\n")

	// pack # 8
	amiPack = append(amiPack, "Event: BridgeEnter\r\n"+
		"Privilege: call,all\r\n"+
		"BridgeUniqueid: ea17f412-5a17-449d-bf52-1d0fc68a13d2\r\n"+
		"BridgeType: basic\r\n"+
		"BridgeTechnology: simple_bridge\r\n"+
		"BridgeCreator: <unknown>\r\n"+
		"BridgeName: <unknown>\r\n"+
		"BridgeNumChannels: 1\r\n"+
		"BridgeVideoSourceMode: none\r\n"+
		"Channel: SIP/okon.ferry.sip.com-0000036b\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 7711\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName:\r\n\r\n")

	// pack # 9
	amiPack = append(amiPack, "Event: ConfbridgeJoin\r\n"+
		"Privilege: call,all\r\n"+
		"Conference: 9170*12\r\n"+
		"BridgeUniqueid: 0974b4b3-d5ea-441d-9f4e-acd2da843ca1\r\n"+
		"BridgeType: base\r\n"+
		"BridgeTechnology: softmix\r\n"+
		"BridgeCreator: ConfBridge\r\n"+
		"BridgeName: 9170*12\r\n"+
		"BridgeNumChannels: 2\r\n"+
		"BridgeVideoSourceMode: none\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;1\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 7711\r\n"+
		"CallerIDName: Evan Beatty\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: confbridge-join-member\r\n"+
		"Exten: MEETME_9170*12\r\n"+
		"Priority: 6\r\n"+
		"Uniqueid: 1598968093.901\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n"+
		"Admin: No\r\n"+
		"Muted: No\r\n\r\n")

	// pack # 10
	amiPack = append(amiPack, "Event: ConfbridgeTalking\r\n"+
		"Privilege: call,all\r\n"+
		"Conference: 9170*12\r\n"+
		"BridgeUniqueid: 0974b4b3-d5ea-441d-9f4e-acd2da843ca1\r\n"+
		"BridgeType: base\r\n"+
		"BridgeTechnology: softmix\r\n"+
		"BridgeCreator: ConfBridge\r\n"+
		"BridgeName: 9170*12\r\n"+
		"BridgeNumChannels: 3\r\n"+
		"BridgeVideoSourceMode: none\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;1\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 7711\r\n"+
		"CallerIDName: Evan Beatty\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: confbridge-join-member\r\n"+
		"Exten: MEETME_9170*12\r\n"+
		"Priority: 6\r\n"+
		"Uniqueid: 1598968093.901\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n"+
		"TalkingStatus: on\r\n"+
		"Admin: No\r\n\r\n")

	// pack # 11
	amiPack = append(amiPack, "Event: ConfbridgeTalking\r\n"+
		"Privilege: call,all\r\n"+
		"Conference: 9170*12\r\n"+
		"BridgeUniqueid: 0974b4b3-d5ea-441d-9f4e-acd2da843ca1\r\n"+
		"BridgeType: base\r\n"+
		"BridgeTechnology: softmix\r\n"+
		"BridgeCreator: ConfBridge\r\n"+
		"BridgeName: 9170*12\r\n"+
		"BridgeNumChannels: 3\r\n"+
		"BridgeVideoSourceMode: none\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;1\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 7711\r\n"+
		"CallerIDName: Evan Beatty\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: confbridge-join-member\r\n"+
		"Exten: MEETME_9170*12\r\n"+
		"Priority: 6\r\n"+
		"Uniqueid: 1598968093.901\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n"+
		"TalkingStatus: off\r\n"+
		"Admin: No\r\n\r\n")

	// pack # 12
	amiPack = append(amiPack, "Action: ConfbridgeKick\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;1\r\n"+
		"Actionid: 5ad12abbacea8@okon.ferry.sip.com\r\n"+
		"Conference: 9170*12\r\n\r\n")

	// pack # 13
	amiPack = append(amiPack, "Response: Success\r\n"+
		"ActionID: 5ad12abbacea8@okon.ferry.sip.com\r\n"+
		"Message: User kicked\r\n\r\n")

	// pack # 14
	amiPack = append(amiPack, "Event: BridgeLeave\r\n"+
		"Privilege: call,all\r\n"+
		"BridgeUniqueid: 0974b4b3-d5ea-441d-9f4e-acd2da843ca1\r\n"+
		"BridgeType: base\r\n"+
		"BridgeTechnology: softmix\r\n"+
		"BridgeCreator: ConfBridge\r\n"+
		"BridgeName: 9170*12\r\n"+
		"BridgeNumChannels: 2\r\n"+
		"BridgeVideoSourceMode: none\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;1\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 7711\r\n"+
		"CallerIDName: Evan Beatty\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: confbridge-join-member\r\n"+
		"Exten: MEETME_9170*12\r\n"+
		"Priority: 6\r\n"+
		"Uniqueid: 1598968093.901\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n\r\n")

	// pack # 15
	amiPack = append(amiPack, "Event: ConfbridgeLeave\r\n"+
		"Privilege: call,all\r\n"+
		"Conference: 9170*12\r\n"+
		"BridgeUniqueid: 0974b4b3-d5ea-441d-9f4e-acd2da843ca1\r\n"+
		"BridgeType: base\r\n"+
		"BridgeTechnology: softmix\r\n"+
		"BridgeCreator: ConfBridge\r\n"+
		"BridgeName: 9170*12\r\n"+
		"BridgeNumChannels: 2\r\n"+
		"BridgeVideoSourceMode: none\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;1\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 7711\r\n"+
		"CallerIDName: Evan Beatty\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: confbridge-join-member\r\n"+
		"Exten: MEETME_9170*12\r\n"+
		"Priority: 6\r\n"+
		"Uniqueid: 1598968093.901\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n"+
		"Admin: No\r\n\r\n")

	// pack # 16
	amiPack = append(amiPack, "Event: SoftHangupRequest\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;1\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 7711\r\n"+
		"CallerIDName: Evan Beatty\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: confbridge-join-member\r\n"+
		"Exten: MEETME_9170*12\r\n"+
		"Priority: 7\r\n"+
		"Uniqueid: 1598968093.901\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n"+
		"Cause: 16\r\n\r\n")

	// pack # 17
	amiPack = append(amiPack, "Event: HangupRequest\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;2\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: confbridge-invite\r\n"+
		"Exten: 7711\r\n"+
		"Priority: 4\r\n"+
		"Uniqueid: 1598968093.902\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n"+
		"Cause: 16\r\n\r\n")

	// pack # 18
	amiPack = append(amiPack, "Event: Hangup\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;1\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 7711\r\n"+
		"CallerIDName: Evan Beatty\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: confbridge-join-member\r\n"+
		"Exten: MEETME_9170*12\r\n"+
		"Priority: 7\r\n"+
		"Uniqueid: 1598968093.901\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n"+
		"Cause: 16\r\n"+
		"Cause-txt: Normal Clearing\r\n\r\n")

	// pack # 19
	amiPack = append(amiPack, "Event: DeviceStateChange\r\n"+
		"Privilege: call,all\r\n"+
		"Device: Local/7711@confbridge-invite\r\n"+
		"State: NOT_INUSE\r\n\r\n")

	// pack # 20
	amiPack = append(amiPack, "Event: BridgeLeave\r\n"+
		"Privilege: call,all\r\n"+
		"BridgeUniqueid: ea17f412-5a17-449d-bf52-1d0fc68a13d2\r\n"+
		"BridgeType: basic\r\n"+
		"BridgeTechnology: simple_bridge\r\n"+
		"BridgeCreator: <unknown>\r\n"+
		"BridgeName: <unknown>\r\n"+
		"BridgeNumChannels: 1\r\n"+
		"BridgeVideoSourceMode: none\r\n"+
		"Channel: Local/7711@confbridge-invite-0000000d;2\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: confbridge-invite\r\n"+
		"Exten: 7711\r\n"+
		"Priority: 4\r\n"+
		"Uniqueid: 1598968093.902\r\n"+
		"Linkedid: 1598968093.901\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n\r\n")

	// pack # 21
	amiPack = append(amiPack, "Event: BridgeDestroy\r\n"+
		"Privilege: call,all\r\n"+
		"BridgeUniqueid: ea17f412-5a17-449d-bf52-1d0fc68a13d2\r\n"+
		"BridgeType: basic\r\n"+
		"BridgeTechnology: simple_bridge\r\n"+
		"BridgeCreator: <unknown>\r\n"+
		"BridgeName: <unknown>\r\n"+
		"BridgeNumChannels: 0\r\n"+
		"BridgeVideoSourceMode: none\r\n\r\n")

	// pack # 22
	amiPack = append(amiPack, "Event: MessageWaiting\r\n"+
		"Privilege: call,all\r\n"+
		"Channel: Local/9922@confbridge-invite-0000000b;2\r\n"+
		"ChannelState: 6\r\n"+
		"ChannelStateDesc: Up\r\n"+
		"CallerIDNum: 9170\r\n"+
		"CallerIDName: Bud Heller\r\n"+
		"ConnectedLineNum: 9170\r\n"+
		"ConnectedLineName: Bud Heller\r\n"+
		"Language: en\r\n"+
		"AccountCode: \r\n"+
		"Context: voicemail\r\n"+
		"Exten: 9922\r\n"+
		"Priority: 10\r\n"+
		"Uniqueid: 1598968054.897\r\n"+
		"Linkedid: 1598968054.896\r\n"+
		"ChanVariable: realm=okon.ferry.sip.com\r\n"+
		"ChanVariable: SIPDOMAIN=\r\n"+
		"ChanVariable: SIPCALLID=\r\n"+
		"Mailbox: 9922@okon.ferry.sip.com\r\n"+
		"Waiting: 1\r\n"+
		"New: 11\r\n"+
		"Old: 4\r\n\r\n")
	return amiPack
}
