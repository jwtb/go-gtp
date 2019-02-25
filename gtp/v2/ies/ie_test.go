// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	v2 "github.com/wmnsk/go-gtp/gtp/v2"
	"github.com/wmnsk/go-gtp/gtp/v2/ies"
)

func TestIEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ies.IE
		serialized  []byte
	}{
		{
			"IMSI",
			ies.NewIMSI("123451234567890"),
			[]byte{0x01, 0x00, 0x08, 0x00, 0x21, 0x43, 0x15, 0x32, 0x54, 0x76, 0x98, 0xf0},
		}, {
			"Cause",
			ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil),
			[]byte{0x02, 0x00, 0x02, 0x00, 0x10, 0x00},
		}, {
			"CauseIMSIIMEINotKnown",
			ies.NewCause(v2.CauseIMSIIMEINotKnown, 1, 0, 0, ies.NewIMSI("")),
			[]byte{0x02, 0x00, 0x03, 0x00, 0x60, 0x04, 0x01},
		}, {
			"Recovery",
			ies.NewRecovery(0xff),
			[]byte{0x03, 0x00, 0x01, 0x00, 0xff},
		}, {
			"AccessPointName",
			ies.NewAccessPointName("some.apn.example"),
			[]byte{0x47, 0x00, 0x11, 0x00, 0x04, 0x73, 0x6f, 0x6d, 0x65, 0x03, 0x61, 0x70, 0x6e, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65},
		}, {
			"AggregateMaximumBitRate",
			ies.NewAggregateMaximumBitRate(0x11111111, 0x22222222),
			[]byte{0x48, 0x00, 0x08, 0x00, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22},
		}, {
			"EPSBearerID",
			ies.NewEPSBearerID(0x05),
			[]byte{0x49, 0x00, 0x01, 0x00, 0x05},
		}, {
			"IPAddress/v4",
			ies.NewIPAddress("1.1.1.1"),
			[]byte{0x4a, 0x00, 0x04, 0x00, 0x01, 0x01, 0x01, 0x01},
		}, {
			"IPAddress/v6",
			ies.NewIPAddress("2001::1"),
			[]byte{0x4a, 0x00, 0x10, 0x00, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		}, {
			"MobileEquipmentIdentity",
			ies.NewMobileEquipmentIdentity("123450123456789"),
			[]byte{0x4b, 0x00, 0x08, 0x00, 0x21, 0x43, 0x05, 0x21, 0x43, 0x65, 0x87, 0xf9},
		}, {
			"MSISDN",
			ies.NewMSISDN("123450123456789"),
			[]byte{0x4c, 0x00, 0x08, 0x00, 0x21, 0x43, 0x05, 0x21, 0x43, 0x65, 0x87, 0xf9},
		}, {
			"Indication",
			ies.NewIndication(
				1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0,
				1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0,
			),
			[]byte{0x4d, 0x00, 0x07, 0x00, 0xa1, 0x08, 0x15, 0x10, 0x88, 0x81, 0x40},
		}, {
			"IndicationFromBitSequence",
			ies.NewIndicationFromBitSequence("10100001000010000001010100010000100010001000000101000"),
			[]byte{0x4d, 0x00, 0x07, 0x00, 0xa1, 0x08, 0x15, 0x10, 0x88, 0x81, 0x40},
		}, {
			"IndicationFromOctets/Full",
			ies.NewIndicationFromOctets(0xa1, 0x08, 0x15, 0x10, 0x88, 0x81, 0x40),
			[]byte{0x4d, 0x00, 0x07, 0x00, 0xa1, 0x08, 0x15, 0x10, 0x88, 0x81, 0x40},
		}, {
			"IndicationFromOctets/Short",
			ies.NewIndicationFromOctets(0xa1, 0x08),
			[]byte{0x4d, 0x00, 0x02, 0x00, 0xa1, 0x08},
		}, {
			"ProtocolConfigurationOptions",
			ies.NewProtocolConfigurationOptions(
				v2.ConfigProtocolPPPWithIP,
				ies.NewConfigurationProtocolOption(v2.ProtoIDIPCP, []byte{0x01, 0x00, 0x00, 0x10, 0x81, 0x06, 0x00, 0x00, 0x00, 0x00, 0x83, 0x06, 0x00, 0x00, 0x00, 0x00}),
				ies.NewConfigurationProtocolOption(v2.ContIDMSSupportofNetworkRequestedBearerControlIndicator, nil),
				ies.NewConfigurationProtocolOption(v2.ContIDIPaddressAllocationViaNASSignalling, nil),
				ies.NewConfigurationProtocolOption(v2.ContIDDNSServerIPv4AddressRequest, nil),
				ies.NewConfigurationProtocolOption(v2.ContIDIPv4LinkMTURequest, nil),
			),
			[]byte{
				0x4e, 0x00, 0x20, 0x00,
				// ConfigurationProtocol
				0x80,
				// IPCP
				0x80, 0x21, 0x10, 0x01, 0x00, 0x00, 0x10, 0x81, 0x06, 0x00, 0x00, 0x00, 0x00, 0x83, 0x06, 0x00, 0x00, 0x00, 0x00,
				// Bearer control indicator
				0x00, 0x05, 0x00,
				// IP alloc via NAS
				0x00, 0x0a, 0x00,
				// DNS server request
				0x00, 0x0d, 0x00,
				// IPv4 link MTU request
				0x00, 0x10, 0x00,
			},
		}, {
			"PDNAddressAllocation/v4",
			ies.NewPDNAddressAllocation("1.1.1.1"),
			[]byte{0x4f, 0x00, 0x05, 0x00, 0x01, 0x01, 0x01, 0x01, 0x01},
		},
		/* XXX - needs fix in NewPDNAddressAllocation!
		{
			"PDNAddressAllocation/v6",
			ies.NewPDNAddressAllocation("2001::1"),
			[]byte{0x4f, 0x00, 0x12, 0x00, 0x02, 0x00, 0x20, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		}, */
		{
			"BearerQoS",
			ies.NewBearerQoS(1, 2, 1, 0xff, 0x1111111111, 0x2222222222, 0x1111111111, 0x2222222222),
			[]byte{0x50, 0x00, 0x16, 0x00, 0x49, 0xff, 0x11, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x22, 0x11, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x22},
		}, {
			"FlowQoS",
			ies.NewFlowQoS(0xff, 0x1111111111, 0x2222222222, 0x1111111111, 0x2222222222),
			[]byte{0x51, 0x00, 0x15, 0x00, 0xff, 0x11, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x22, 0x11, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x22},
		}, {
			"RATType",
			ies.NewRATType(v2.RATTypeEUTRAN),
			[]byte{0x52, 0x00, 0x01, 0x00, 0x06},
		}, {
			"ServingNetwork/2-digit",
			ies.NewServingNetwork("123", "45"),
			[]byte{0x53, 0x00, 0x03, 0x00, 0x21, 0xf3, 0x54},
		}, {
			"ServingNetwork/3-digit",
			ies.NewServingNetwork("123", "456"),
			[]byte{0x53, 0x00, 0x03, 0x00, 0x21, 0x63, 0x54},
		},
		/* { XXX - implement!
			"EPSBearerLevelTrafficFlowTemplate",
			ies.NewEPSBearerLevelTrafficFlowTemplate(),
			[]byte{},
		},*/
		/* { XXX - implement! (same as Bearer TFT)
			"TrafficAggregateDescription",
			ies.NewTrafficAggregateDescription(),
			[]byte{},
		},*/
		{
			"UserLocationInformation/Lazy-1",
			ies.NewUserLocationInformationLazy(
				"123", "45",
				0x1111, 0x2222, 0x3333, -1, 0x5555, 0x666666, -1, 0x22222222,
			),
			[]byte{
				0x56, 0x00, 0x26, 0x00,
				// Flags
				0xbb,
				// CGI
				0x21, 0xf3, 0x54, 0x11, 0x11, 0x22, 0x22,
				// SAI
				0x21, 0xf3, 0x54, 0x11, 0x11, 0x33, 0x33,
				// TAI
				0x21, 0xf3, 0x54, 0x55, 0x55,
				// ECGI
				0x21, 0xf3, 0x54, 0x00, 0x06, 0x66, 0x66,
				// RAI
				0x21, 0xf3, 0x54, 0x11, 0x11,
				// Extended Macro eNB ID
				0x21, 0xf3, 0x54, 0x22, 0x22, 0x22,
			},
		}, {
			"UserLocationInformation/Lazy-2",
			ies.NewUserLocationInformationLazy(
				"123", "45",
				0x1111, 0x2222, 0x3333, 0x4444, 0x5555, 0x666666, 0x11111111, 0x22222222,
			),
			[]byte{
				0x56, 0x00, 0x33, 0x00,
				// Flags
				0xff,
				// CGI
				0x21, 0xf3, 0x54, 0x11, 0x11, 0x22, 0x22,
				// SAI
				0x21, 0xf3, 0x54, 0x11, 0x11, 0x33, 0x33,
				// RAI
				0x21, 0xf3, 0x54, 0x11, 0x11, 0x44, 0x44,
				// TAI
				0x21, 0xf3, 0x54, 0x55, 0x55,
				// ECGI
				0x21, 0xf3, 0x54, 0x00, 0x06, 0x66, 0x66,
				// RAI
				0x21, 0xf3, 0x54, 0x11, 0x11,
				// Macro eNB ID
				0x21, 0xf3, 0x54, 0x11, 0x11, 0x11,
				// Extended Macro eNB ID
				0x21, 0xf3, 0x54, 0x22, 0x22, 0x22,
			},
		}, {
			"UserLocationInformation/Full",
			ies.NewUserLocationInformation(
				1, 1, 1, 1, 1, 1, 1, 1, "123", "45",
				0x1111, 0x2222, 0x3333, 0x4444, 0x5555, 0x666666, 0x11111111, 0x22222222,
			),
			[]byte{
				0x56, 0x00, 0x33, 0x00,
				// Flags
				0xff,
				// CGI
				0x21, 0xf3, 0x54, 0x11, 0x11, 0x22, 0x22,
				// SAI
				0x21, 0xf3, 0x54, 0x11, 0x11, 0x33, 0x33,
				// RAI
				0x21, 0xf3, 0x54, 0x11, 0x11, 0x44, 0x44,
				// TAI
				0x21, 0xf3, 0x54, 0x55, 0x55,
				// ECGI
				0x21, 0xf3, 0x54, 0x00, 0x06, 0x66, 0x66,
				// RAI
				0x21, 0xf3, 0x54, 0x11, 0x11,
				// Macro eNB ID
				0x21, 0xf3, 0x54, 0x11, 0x11, 0x11,
				// Extended Macro eNB ID
				0x21, 0xf3, 0x54, 0x22, 0x22, 0x22,
			},
		}, {
			"FullyQualifiedTEID/v4",
			ies.NewFullyQualifiedTEID(v2.IFTypeS11MMEGTPC, 0xffffffff, "1.1.1.1", ""),
			[]byte{0x57, 0x00, 0x09, 0x00, 0x8a, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x01},
		}, {
			"FullyQualifiedTEID/v6",
			ies.NewFullyQualifiedTEID(v2.IFTypeS11MMEGTPC, 0xffffffff, "", "2001::1"),
			[]byte{0x57, 0x00, 0x15, 0x00, 0x4a, 0xff, 0xff, 0xff, 0xff, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		}, {
			"FullyQualifiedTEID/v4v6",
			ies.NewFullyQualifiedTEID(v2.IFTypeS11MMEGTPC, 0xffffffff, "1.1.1.1", "2001::1"),
			[]byte{0x57, 0x00, 0x19, 0x00, 0xca, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x01, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		}, {
			"TMSI",
			ies.NewTMSI(0xffffffff),
			[]byte{0x58, 0x00, 0x04, 0x00, 0xff, 0xff, 0xff, 0xff},
		}, {
			"GlobalCNID",
			ies.NewGlobalCNID("123", "45", 0xfff),
			[]byte{0x59, 0x00, 0x05, 0x00, 0x21, 0xf3, 0x54, 0x0f, 0xff},
		}, {
			"S103PDNDataForwardingInfo/v4",
			ies.NewS103PDNDataForwardingInfo("1.1.1.1", 0xdeadbeef, 5, 6, 7),
			[]byte{0x5a, 0x00, 0x0d, 0x00, 0x04, 0x01, 0x01, 0x01, 0x01, 0xde, 0xad, 0xbe, 0xef, 0x03, 0x05, 0x06, 0x07},
		}, {
			"S103PDNDataForwardingInfo/v6",
			ies.NewS103PDNDataForwardingInfo("2001::1", 0xdeadbeef, 5, 6, 7),
			[]byte{0x5a, 0x00, 0x19, 0x00, 0x10, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xde, 0xad, 0xbe, 0xef, 0x03, 0x05, 0x06, 0x07},
		}, {
			"S1UDataForwarding/v4",
			ies.NewS1UDataForwarding("1.1.1.1", 0xdeadbeef),
			[]byte{0x5b, 0x00, 0x09, 0x00, 0x04, 0x01, 0x01, 0x01, 0x01, 0xde, 0xad, 0xbe, 0xef},
		}, {
			"S1UDataForwarding/v6",
			ies.NewS1UDataForwarding("2001::1", 0xdeadbeef),
			[]byte{0x5b, 0x00, 0x15, 0x00, 0x10, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xde, 0xad, 0xbe, 0xef},
		}, {
			"DelayValue",
			ies.NewDelayValue(500 * time.Millisecond),
			[]byte{0x5c, 0x00, 0x01, 0x00, 0x0a},
		}, {
			"BearerContext",
			ies.NewBearerContext(ies.NewDelayValue(500*time.Millisecond), ies.NewDelayValue(100*time.Millisecond)),
			[]byte{0x5d, 0x00, 0x0a, 0x00, 0x5c, 0x00, 0x01, 0x00, 0x0a, 0x5c, 0x00, 0x01, 0x00, 0x02},
		}, {
			"ChargingID",
			ies.NewChargingID(0xffffffff),
			[]byte{0x5e, 0x00, 0x04, 0x00, 0xff, 0xff, 0xff, 0xff},
		}, {
			"ChargingCharacteristics",
			ies.NewChargingCharacteristics(0xffff),
			[]byte{0x5f, 0x00, 0x02, 0x00, 0xff, 0xff},
		}, {
			"BearerFlags",
			ies.NewBearerFlags(1, 1, 1, 1),
			[]byte{0x61, 0x00, 0x01, 0x00, 0x0f},
		}, {
			"PDNType",
			ies.NewPDNType(v2.PDNTypeIPv4),
			[]byte{0x63, 0x00, 0x01, 0x00, 0x01},
		}, {
			"ProcedureTransactionID",
			ies.NewProcedureTransactionID(1),
			[]byte{0x64, 0x00, 0x01, 0x00, 0x01},
		}, {
			"PacketTMSI",
			ies.NewPacketTMSI(0xdeadbeef),
			[]byte{0x6f, 0x00, 0x04, 0x00, 0xde, 0xad, 0xbe, 0xef},
		}, {
			"PTMSISignature",
			ies.NewPTMSISignature(0xbeebee),
			[]byte{0x70, 0x00, 0x03, 0x00, 0xbe, 0xeb, 0xee},
		}, {
			"HopCounter",
			ies.NewHopCounter(1),
			[]byte{0x71, 0x00, 0x01, 0x00, 0x01},
		}, {
			"UETimeZone",
			ies.NewUETimeZone(9*time.Hour, 0),
			[]byte{0x72, 0x00, 0x02, 0x00, 0x63, 0x00},
		}, {
			"TraceReference",
			ies.NewTraceReference("123", "45", 1),
			[]byte{0x73, 0x00, 0x06, 0x00, 0x21, 0xf3, 0x54, 0x00, 0x00, 0x01},
		}, {
			"GUTI",
			ies.NewGUTI("123", "45", 0x1111, 0x22, 0x33333333),
			[]byte{0x75, 0x00, 0x0a, 0x00, 0x21, 0xf3, 0x54, 0x11, 0x11, 0x22, 0x33, 0x33, 0x33, 0x33},
		}, {
			"PLMNID/2digits",
			ies.NewPLMNID("123", "45"),
			[]byte{0x78, 0x00, 0x03, 0x00, 0x21, 0xf3, 0x54},
		}, {
			"PLMNID/3digits",
			ies.NewPLMNID("123", "456"),
			[]byte{0x78, 0x00, 0x03, 0x00, 0x21, 0x63, 0x54},
		}, {
			"PortNumber",
			ies.NewPortNumber(2123),
			[]byte{0x7e, 0x00, 0x02, 0x00, 0x08, 0x4b},
		}, {
			"APNRestriction",
			ies.NewAPNRestriction(v2.APNRestrictionPublic1),
			[]byte{0x7f, 0x00, 0x01, 0x00, 0x01},
		}, {
			"SelectionMode",
			ies.NewSelectionMode(v2.SelectionModeMSProvidedAPNSubscriptionNotVerified),
			[]byte{0x80, 0x00, 0x01, 0x00, 0x01},
		}, {
			"FullyQualifiedCSID/v4",
			ies.NewFullyQualifiedCSID("1.1.1.1", 1),
			[]byte{0x84, 0x00, 0x07, 0x00, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01},
		}, {
			"FullyQualifiedCSID/v4/multiCSIDs",
			ies.NewFullyQualifiedCSID("1.1.1.1", 1, 2),
			[]byte{0x84, 0x00, 0x09, 0x00, 0x02, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01, 0x00, 0x02},
		}, {
			"FullyQualifiedCSID/v6",
			ies.NewFullyQualifiedCSID("2001::1", 1),
			[]byte{0x84, 0x00, 0x13, 0x00, 0x11, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01},
		}, {
			"FullyQualifiedCSID/other",
			ies.NewFullyQualifiedCSID("12304501", 1),
			[]byte{0x84, 0x00, 0x07, 0x00, 0x21, 0x12, 0x30, 0x45, 0x01, 0x00, 0x01},
		}, {
			"NodeType",
			ies.NewNodeType(v2.NodeTypeMME),
			[]byte{0x87, 0x00, 0x01, 0x00, 0x01},
		}, {
			"FullyQualifiedDomainName",
			ies.NewFullyQualifiedDomainName("some-fqdn.example"),
			[]byte{0x88, 0x00, 0x11, 0x00, 0x73, 0x6f, 0x6d, 0x65, 0x2d, 0x66, 0x71, 0x64, 0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65},
		}, {
			"RFSPIndex",
			ies.NewRFSPIndex(1),
			[]byte{0x90, 0x00, 0x01, 0x00, 0x01},
		}, {
			"UserCSGInformation",
			ies.NewUserCSGInformation("123", "45", 0x00ffffff, v2.AccessModeHybrid, 0, v2.CMICSG),
			[]byte{0x91, 0x00, 0x08, 0x00, 0x21, 0xf3, 0x54, 0x00, 0xff, 0xff, 0xff, 0x41},
		}, {
			"CSGID",
			ies.NewCSGID(0x00ffffff),
			[]byte{0x93, 0x00, 0x04, 0x00, 0x00, 0xff, 0xff, 0xff},
		}, {
			"CSGMembershipIndication",
			ies.NewCSGMembershipIndication(v2.CMICSG),
			[]byte{0x94, 0x00, 0x01, 0x00, 0x01},
		}, {
			"ServiceIndicator",
			ies.NewServiceIndicator(v2.ServiceIndCSCall),
			[]byte{0x95, 0x00, 0x01, 0x00, 0x01},
		}, {
			"DetachType",
			ies.NewDetachType(v2.DetachTypePS),
			[]byte{0x96, 0x00, 0x01, 0x00, 0x01},
		}, {
			"LocalDistinguishedName",
			ies.NewLocalDistinguishedName("some-name"),
			[]byte{0x97, 0x00, 0x09, 0x00, 0x73, 0x6f, 0x6d, 0x65, 0x2d, 0x6e, 0x61, 0x6d, 0x65},
		}, {
			"AllocationRetensionPriority",
			ies.NewAllocationRetensionPriority(1, 2, 1),
			[]byte{0x9b, 0x00, 0x01, 0x00, 0x49},
		}, {
			"ULITimestamp",
			ies.NewULITimestamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			[]byte{0xaa, 0x00, 0x04, 0x00, 0xdf, 0xd5, 0x2c, 0x00},
		}, {
			"MBMSFlags",
			ies.NewMBMSFlags(1, 1),
			[]byte{0xab, 0x00, 0x01, 0x00, 0x03},
		}, {
			"PrivateExtension",
			ies.NewPrivateExtension(10415, []byte{0xde, 0xad, 0xbe, 0xef}),
			[]byte{0xff, 0x00, 0x06, 0x00, 0x28, 0xaf, 0xde, 0xad, 0xbe, 0xef},
		},
	}

	for _, c := range cases {
		t.Run("serialize/"+c.description, func(t *testing.T) {
			got, err := c.structured.Serialize()
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.serialized); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("decode/"+c.description, func(t *testing.T) {
			got, err := ies.Decode(c.serialized)
			if err != nil {
				t.Fatal(err)
			}

			opt := cmp.AllowUnexported(*got, *c.structured)
			if diff := cmp.Diff(got, c.structured, opt); diff != "" {
				t.Error(diff)
			}
		})
	}
}