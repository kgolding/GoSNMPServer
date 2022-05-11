package gosnmp

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestUnmarshalMsgV3UnknownPasswordEncryptedPDU(t *testing.T) {
	Default.Logger = log.New(ioutil.Discard, "", 0)

	vhandle := GoSNMP{}
	vhandle.Logger = Default.Logger
	testBytes := v3UnknownPasswordEncryptedPDU()
	res, _ := vhandle.SnmpDecodePacket(testBytes)
	if res.SecurityParameters.(*UsmSecurityParameters).UserName != "simulator" {
		t.Errorf("#v3UnknownPasswordEncryptedPDU: SnmpDecodePacket() err Username: %v [expect simulator]", res.SecurityParameters.(*UsmSecurityParameters).UserName)
	}
}

func TestUnmarshalMsgV3NotEncrypted(t *testing.T) {
	Default.Logger = log.New(ioutil.Discard, "", 0)

	vhandle := GoSNMP{}
	vhandle.Logger = Default.Logger
	testBytes := v3NotEncrypted()
	res, _ := vhandle.SnmpDecodePacket(testBytes)
	if res.SecurityParameters.(*UsmSecurityParameters).UserName != "testuser" {
		t.Errorf("#v3UnknownPasswordEncryptedPDU: SnmpDecodePacket() err Username: %v [expect simulator]", res.SecurityParameters.(*UsmSecurityParameters).UserName)
	}
}

func TestUnmarshalMsgV3Encrypted(t *testing.T) {
	Default.Logger = log.New(ioutil.Discard, "", 0)

	vhandle := GoSNMP{}
	vhandle.SecurityParameters = &UsmSecurityParameters{
		UserName:                 "testuser",
		AuthenticationProtocol:   MD5,
		PrivacyProtocol:          DES,
		AuthenticationPassphrase: "testauth",
		PrivacyPassphrase:        "testpriv",
		Logger:                   Default.Logger,
	}
	vhandle.Logger = Default.Logger
	testBytes := v3Encrypted()
	res, err := vhandle.SnmpDecodePacket(testBytes)
	if err != nil {
		t.Errorf("#v3UnknownPasswordEncryptedPDU: SnmpDecodePacket() meet err %v", err)
	}
	if res.SecurityParameters.(*UsmSecurityParameters).UserName != "testuser" {
		t.Errorf("#v3UnknownPasswordEncryptedPDU: SnmpDecodePacket() err Username: %v [expect simulator]", res.SecurityParameters.(*UsmSecurityParameters).UserName)
	}
}

// Simple Network Management Protocol
// msgVersion: snmpv3 (3)
// msgGlobalData
// 	msgID: 91040641
// 	msgMaxSize: 65507
// 	msgFlags: 07
// 		.... .1.. = Reportable: Set
// 		.... ..1. = Encrypted: Set
// 		.... ...1 = Authenticated: Set
// 	msgSecurityModel: USM (3)
// msgAuthoritativeEngineID: 80004fb8054445534b544f502d4a3732533245343ab63bc8
// 	1... .... = Engine ID Conformance: RFC3411 (SNMPv3)
// 	Engine Enterprise ID: pysnmp (20408)
// 	Engine ID Format: Octets, administratively assigned (5)
// 	Engine ID Data: 4445534b544f502d4a3732533245343ab63bc8
// msgAuthoritativeEngineBoots: 2
// msgAuthoritativeEngineTime: 50298
// msgUserName: simulator
// msgAuthenticationParameters: 1feffe5890da21ec984367c8
// msgPrivacyParameters: 000000010a8e365d
// msgData: encryptedPDU (1)
// 	encryptedPDU: 8dc7b6f6d0660dd2acd80db15af11d42db52423b3ddab01f…

func v3UnknownPasswordEncryptedPDU() []byte {
	return []byte{
		0x30, 0x81, 0xb9, 0x02, 0x01, 0x03, 0x30, 0x11,
		0x02, 0x04, 0x05, 0x6d, 0x2b, 0x81, 0x02, 0x03,
		0x00, 0xff, 0xe3, 0x04, 0x01, 0x07, 0x02, 0x01,
		0x03, 0x04, 0x47, 0x30, 0x45, 0x04, 0x18, 0x80,
		0x00, 0x4f, 0xb8, 0x05, 0x44, 0x45, 0x53, 0x4b,
		0x54, 0x4f, 0x50, 0x2d, 0x4a, 0x37, 0x32, 0x53,
		0x32, 0x45, 0x34, 0x3a, 0xb6, 0x3b, 0xc8, 0x02,
		0x01, 0x02, 0x02, 0x03, 0x00, 0xc4, 0x7a, 0x04,
		0x09, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
		0x6f, 0x72, 0x04, 0x0c, 0x1f, 0xef, 0xfe, 0x58,
		0x90, 0xda, 0x21, 0xec, 0x98, 0x43, 0x67, 0xc8,
		0x04, 0x08, 0x00, 0x00, 0x00, 0x01, 0x0a, 0x8e,
		0x36, 0x5d, 0x04, 0x58, 0x8d, 0xc7, 0xb6, 0xf6,
		0xd0, 0x66, 0x0d, 0xd2, 0xac, 0xd8, 0x0d, 0xb1,
		0x5a, 0xf1, 0x1d, 0x42, 0xdb, 0x52, 0x42, 0x3b,
		0x3d, 0xda, 0xb0, 0x1f, 0x19, 0x77, 0x27, 0xa2,
		0x0a, 0xe9, 0x2b, 0x73, 0x68, 0xc4, 0x3c, 0xec,
		0xd1, 0x2c, 0xac, 0x18, 0x1f, 0x10, 0x34, 0x05,
		0xe9, 0x45, 0xcc, 0x22, 0x79, 0xf7, 0x31, 0x28,
		0xcc, 0x41, 0x3b, 0x08, 0xc7, 0x4e, 0x1b, 0x22,
		0x8e, 0x32, 0x2d, 0x71, 0x3e, 0x6c, 0x92, 0xbb,
		0x00, 0x7c, 0x84, 0xfd, 0xa2, 0x05, 0x04, 0xff,
		0x79, 0xb5, 0x37, 0xd5, 0xad, 0xc1, 0x7e, 0x80,
		0xfa, 0x7e, 0x06, 0xa9,
	}
}

// Simple Network Management Protocol
//     msgVersion: snmpv3 (3)
//     msgGlobalData
//         msgID: 318530550
//         msgMaxSize: 65507
//         msgFlags: 04
//             .... .1.. = Reportable: Set
//             .... ..0. = Encrypted: Not set
//             .... ...0 = Authenticated: Not set
//         msgSecurityModel: USM (3)
//     msgAuthoritativeEngineID: 80004fb8054445534b544f502d4a3732533245343ab63bc8
//     msgAuthoritativeEngineBoots: 1
//     msgAuthoritativeEngineTime: 173735
//     msgUserName: testuser
//     msgAuthenticationParameters: <MISSING>
//     msgPrivacyParameters: <MISSING>
//     msgData: plaintext (0)
//         plaintext
//             contextEngineID: 80004fb8054445534b544f502d4a3732533245343ab63bc8
//             contextName: public
//             data: get-request (0)
//                 get-request
//                     request-id: 1981244075
//                     error-status: noError (0)
//                     error-index: 0
//                     variable-bindings: 1 item
//                         1.3.6.1.2.1.43.14.1.1.6.1.5: Value (Null)
//                             Object Name: 1.3.6.1.2.1.43.14.1.1.6.1.5 (iso.3.6.1.2.1.43.14.1.1.6.1.5)
//                             Value (Null)

func v3NotEncrypted() []byte {
	return []byte{
		0x30, 0x81, 0x90, 0x02, 0x01, 0x03, 0x30, 0x11,
		0x02, 0x04, 0x12, 0xfc, 0x63, 0xf6, 0x02, 0x03,
		0x00, 0xff, 0xe3, 0x04, 0x01, 0x04, 0x02, 0x01,
		0x03, 0x04, 0x32, 0x30, 0x30, 0x04, 0x18, 0x80,
		0x00, 0x4f, 0xb8, 0x05, 0x44, 0x45, 0x53, 0x4b,
		0x54, 0x4f, 0x50, 0x2d, 0x4a, 0x37, 0x32, 0x53,
		0x32, 0x45, 0x34, 0x3a, 0xb6, 0x3b, 0xc8, 0x02,
		0x01, 0x01, 0x02, 0x03, 0x02, 0xa6, 0xa7, 0x04,
		0x08, 0x74, 0x65, 0x73, 0x74, 0x75, 0x73, 0x65,
		0x72, 0x04, 0x00, 0x04, 0x00, 0x30, 0x44, 0x04,
		0x18, 0x80, 0x00, 0x4f, 0xb8, 0x05, 0x44, 0x45,
		0x53, 0x4b, 0x54, 0x4f, 0x50, 0x2d, 0x4a, 0x37,
		0x32, 0x53, 0x32, 0x45, 0x34, 0x3a, 0xb6, 0x3b,
		0xc8, 0x04, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69,
		0x63, 0xa0, 0x20, 0x02, 0x04, 0x76, 0x17, 0x62,
		0xab, 0x02, 0x01, 0x00, 0x02, 0x01, 0x00, 0x30,
		0x12, 0x30, 0x10, 0x06, 0x0c, 0x2b, 0x06, 0x01,
		0x02, 0x01, 0x2b, 0x0e, 0x01, 0x01, 0x06, 0x01,
		0x05, 0x05, 0x00,
	}
}

// Simple Network Management Protocol
//     msgVersion: snmpv3 (3)
//     msgGlobalData
//         msgID: 1246299848
//         msgMaxSize: 65507
//         msgFlags: 07
//             .... .1.. = Reportable: Set
//             .... ..1. = Encrypted: Set
//             .... ...1 = Authenticated: Set
//         msgSecurityModel: USM (3)
//     msgAuthoritativeEngineID: 80004fb8054445534b544f502d4a3732533245343ab63bc8
//         1... .... = Engine ID Conformance: RFC3411 (SNMPv3)
//         Engine Enterprise ID: pysnmp (20408)
//         Engine ID Format: Octets, administratively assigned (5)
//         Engine ID Data: 4445534b544f502d4a3732533245343ab63bc8
//     msgAuthoritativeEngineBoots: 1
//     msgAuthoritativeEngineTime: 514241
//     msgUserName: testuser
//     msgAuthenticationParameters: c72b2bec27bdf6bc3e8ea66a
//     msgPrivacyParameters: 00000001b0bd8ffd
//     msgData: encryptedPDU (1)
//         encryptedPDU: e0f055df57179446c914f4201f7892627738c87d52584e83…
//             Decrypted ScopedPDU: 3039041880004fb8054445534b544f502d4a373253324534…
//                 contextEngineID: 80004fb8054445534b544f502d4a3732533245343ab63bc8
//                     1... .... = Engine ID Conformance: RFC3411 (SNMPv3)
//                     Engine Enterprise ID: pysnmp (20408)
//                     Engine ID Format: Octets, administratively assigned (5)
//                     Engine ID Data: 4445534b544f502d4a3732533245343ab63bc8
//                 contextName: public
//                 data: get-next-request (1)
//                     get-next-request
//                         request-id: 1224110830
//                         error-status: noError (0)
//                         error-index: 0
//                         variable-bindings: 1 item
//                             itu-t.1 (0.1): Value (Null)
//                                 Object Name: 0.1 (itu-t.1)
//                                 Value (Null)
func v3Encrypted() []byte {
	return []byte{
		0x30, 0x81, 0xa0, 0x02, 0x01, 0x03, 0x30, 0x11,
		0x02, 0x04, 0x4a, 0x49, 0x06, 0xc8, 0x02, 0x03,
		0x00, 0xff, 0xe3, 0x04, 0x01, 0x07, 0x02, 0x01,
		0x03, 0x04, 0x46, 0x30, 0x44, 0x04, 0x18, 0x80,
		0x00, 0x4f, 0xb8, 0x05, 0x44, 0x45, 0x53, 0x4b,
		0x54, 0x4f, 0x50, 0x2d, 0x4a, 0x37, 0x32, 0x53,
		0x32, 0x45, 0x34, 0x3a, 0xb6, 0x3b, 0xc8, 0x02,
		0x01, 0x01, 0x02, 0x03, 0x07, 0xd8, 0xc1, 0x04,
		0x08, 0x74, 0x65, 0x73, 0x74, 0x75, 0x73, 0x65,
		0x72, 0x04, 0x0c, 0xc7, 0x2b, 0x2b, 0xec, 0x27,
		0xbd, 0xf6, 0xbc, 0x3e, 0x8e, 0xa6, 0x6a, 0x04,
		0x08, 0x00, 0x00, 0x00, 0x01, 0xb0, 0xbd, 0x8f,
		0xfd, 0x04, 0x40, 0xe0, 0xf0, 0x55, 0xdf, 0x57,
		0x17, 0x94, 0x46, 0xc9, 0x14, 0xf4, 0x20, 0x1f,
		0x78, 0x92, 0x62, 0x77, 0x38, 0xc8, 0x7d, 0x52,
		0x58, 0x4e, 0x83, 0xea, 0x35, 0x02, 0x0d, 0x6a,
		0x25, 0xc9, 0x74, 0xa6, 0xbd, 0x5f, 0xb7, 0x5f,
		0x6e, 0x3f, 0xee, 0x3c, 0xc6, 0x8b, 0x14, 0x98,
		0xe7, 0x18, 0x85, 0x2e, 0x5b, 0xe4, 0x95, 0x8a,
		0x1c, 0xa3, 0x7f, 0x77, 0x50, 0xfc, 0xe4, 0x04,
		0x3d, 0x93, 0x9e,
	}
}
