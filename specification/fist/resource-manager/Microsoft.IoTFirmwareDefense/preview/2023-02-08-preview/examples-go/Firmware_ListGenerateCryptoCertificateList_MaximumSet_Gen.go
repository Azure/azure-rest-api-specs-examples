package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_ListGenerateCryptoCertificateList_MaximumSet_Gen.json
func ExampleFirmwareClient_NewListGenerateCryptoCertificateListPager_firmwareListGenerateCryptoCertificateListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirmwareClient().NewListGenerateCryptoCertificateListPager("FirmwareAnalysisRG", "default", "DECAFBAD-0000-0000-0000-BADBADBADBAD", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.CryptoCertificateList = armiotfirmwaredefense.CryptoCertificateList{
		// 	Value: []*armiotfirmwaredefense.CryptoCertificate{
		// 		{
		// 			Name: to.Ptr("VeriSign Class 3"),
		// 			CryptoCertID: to.Ptr("fafcad44-1770-41f3-998e-875fa688c143"),
		// 			Encoding: to.Ptr("PEM"),
		// 			ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2012-02-12T20:50:16.753Z"); return t}()),
		// 			FilePaths: []*string{
		// 				to.Ptr("this/is/a/path"),
		// 				to.Ptr("/this/is/another/path")},
		// 				Fingerprint: to.Ptr("eb:04:cf:5e:b1:f3:9a:fa:76:2f:2b:b1:20:f2:96:cb:a5:20:c1:b9:7d:b1:58:95:65:b8:1c:b9:a1:7b:72:44"),
		// 				IsExpired: to.Ptr(armiotfirmwaredefense.IsExpiredTrue),
		// 				IsSelfSigned: to.Ptr(armiotfirmwaredefense.IsSelfSignedFalse),
		// 				IsShortKeySize: to.Ptr(armiotfirmwaredefense.IsShortKeySizeTrue),
		// 				IsWeakSignature: to.Ptr(armiotfirmwaredefense.IsWeakSignatureTrue),
		// 				IssuedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2010-02-12T20:50:16.753Z"); return t}()),
		// 				Issuer: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 					CommonName: to.Ptr("VeriSign Class 3"),
		// 					Country: to.Ptr("United States"),
		// 					Organization: to.Ptr("Verisign, Inc."),
		// 					OrganizationalUnit: to.Ptr("(c) 1999 Verisign inc."),
		// 					State: to.Ptr("NY"),
		// 				},
		// 				KeyAlgorithm: to.Ptr("RSA Encryption"),
		// 				KeySize: to.Ptr[int64](1024),
		// 				PairedKey: &armiotfirmwaredefense.PairedKey{
		// 					Type: to.Ptr("Private Key"),
		// 					AdditionalProperties: map[string]any{
		// 						"displayName": "keyfilename.cer",
		// 					},
		// 					ID: to.Ptr("fafcad44-1770-41f3-998e-875fa688c143"),
		// 				},
		// 				Role: to.Ptr("Root CA"),
		// 				SerialNumber: to.Ptr("206684696279472310254277870180966723415"),
		// 				SignatureAlgorithm: to.Ptr("MD5 with RSA Encryption"),
		// 				Subject: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 					CommonName: to.Ptr("VeriSign Class 3"),
		// 					Country: to.Ptr("United States"),
		// 					Organization: to.Ptr("Verisign, Inc."),
		// 					OrganizationalUnit: to.Ptr("(c) 1999 Verisign inc."),
		// 					State: to.Ptr("NY"),
		// 				},
		// 				Usage: []*string{
		// 					to.Ptr("Crypt"),
		// 					to.Ptr("sign"),
		// 					to.Ptr("auth")},
		// 				},
		// 				{
		// 					Name: to.Ptr("VeriSign Class 2"),
		// 					CryptoCertID: to.Ptr("dfc4a7e4-2037-482a-bdf8-81289335337e"),
		// 					Encoding: to.Ptr("PEM"),
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1999-12-31T15:59:59.999Z"); return t}()),
		// 					FilePaths: []*string{
		// 						to.Ptr("this/is/a/path"),
		// 						to.Ptr("/this/is/another/path")},
		// 						Fingerprint: to.Ptr("eb:04:cf:5e:b1:f3:9a:fa:76:2f:2b:b1:20:f2:96:cb:a5:20:c1:b9:7d:b1:58:95:65:b8:1c:b9:a1:7b:72:44"),
		// 						IsExpired: to.Ptr(armiotfirmwaredefense.IsExpiredTrue),
		// 						IsSelfSigned: to.Ptr(armiotfirmwaredefense.IsSelfSignedFalse),
		// 						IsShortKeySize: to.Ptr(armiotfirmwaredefense.IsShortKeySizeFalse),
		// 						IsWeakSignature: to.Ptr(armiotfirmwaredefense.IsWeakSignatureFalse),
		// 						IssuedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1970-01-01T00:00:00.000Z"); return t}()),
		// 						Issuer: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 							CommonName: to.Ptr("VeriSign Class 3"),
		// 							Country: to.Ptr("United States"),
		// 							Organization: to.Ptr("Verisign, Inc."),
		// 							OrganizationalUnit: to.Ptr("(c) 1999 Verisign inc."),
		// 							State: to.Ptr("NY"),
		// 						},
		// 						KeyAlgorithm: to.Ptr("RSA Encryption"),
		// 						KeySize: to.Ptr[int64](2048),
		// 						PairedKey: &armiotfirmwaredefense.PairedKey{
		// 							Type: to.Ptr("Private Key"),
		// 							ID: to.Ptr("fafcad44-1770-41f3-998e-875fa688c143"),
		// 						},
		// 						Role: to.Ptr("Intermediate CA"),
		// 						SerialNumber: to.Ptr("206684696279472310254277870180966723415"),
		// 						SignatureAlgorithm: to.Ptr("SHA-1 with RSA Encryption"),
		// 						Subject: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 							CommonName: to.Ptr("VeriSign Class 3"),
		// 							Country: to.Ptr("United States"),
		// 							Organization: to.Ptr("Verisign, Inc."),
		// 							OrganizationalUnit: to.Ptr("(c) 1999 Verisign inc."),
		// 							State: to.Ptr("NY"),
		// 						},
		// 						Usage: []*string{
		// 							to.Ptr("Crypt"),
		// 							to.Ptr("sign"),
		// 							to.Ptr("auth")},
		// 						},
		// 						{
		// 							Name: to.Ptr("JillianSign Class 1"),
		// 							CryptoCertID: to.Ptr("55d8329d-6f4b-4a36-a6a7-53f2ddf187b5"),
		// 							Encoding: to.Ptr("PEM"),
		// 							ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-16T20:50:16.753Z"); return t}()),
		// 							FilePaths: []*string{
		// 								to.Ptr("this/is/a/path/made/by/Jillian"),
		// 								to.Ptr("/this/is/another/path")},
		// 								Fingerprint: to.Ptr("eb:04:cf:5e:b1:f3:9a:fa:76:2f:2b:b1:20:f2:96:cb:a5:20:c1:b9:7d:b1:58:95:65:b8:1c:b9:a1:7b:72:44"),
		// 								IsExpired: to.Ptr(armiotfirmwaredefense.IsExpiredTrue),
		// 								IsSelfSigned: to.Ptr(armiotfirmwaredefense.IsSelfSignedTrue),
		// 								IsShortKeySize: to.Ptr(armiotfirmwaredefense.IsShortKeySizeFalse),
		// 								IsWeakSignature: to.Ptr(armiotfirmwaredefense.IsWeakSignatureFalse),
		// 								IssuedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-16T20:50:16.753Z"); return t}()),
		// 								Issuer: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 									CommonName: to.Ptr("VeriSign Class 3"),
		// 									Country: to.Ptr("United States"),
		// 									Organization: to.Ptr("Verisign, Inc."),
		// 									OrganizationalUnit: to.Ptr("(c) 1999 Verisign inc."),
		// 									State: to.Ptr("NY"),
		// 								},
		// 								KeyAlgorithm: to.Ptr("RSA Encryption"),
		// 								KeySize: to.Ptr[int64](2048),
		// 								Role: to.Ptr("Standard"),
		// 								SerialNumber: to.Ptr("206684696279472310254277870180966723415"),
		// 								SignatureAlgorithm: to.Ptr("SHA-1 with RSA Encryption"),
		// 								Subject: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 									CommonName: to.Ptr("JillianSign Class 1"),
		// 									Country: to.Ptr("United States"),
		// 									Organization: to.Ptr("Jilliansign, Inc."),
		// 									OrganizationalUnit: to.Ptr("(c) 2007 Jilliansign inc."),
		// 									State: to.Ptr("MA"),
		// 								},
		// 								Usage: []*string{
		// 									to.Ptr("Crypt"),
		// 									to.Ptr("sign"),
		// 									to.Ptr("auth")},
		// 								},
		// 								{
		// 									Name: to.Ptr("MikeSign Class 5"),
		// 									CryptoCertID: to.Ptr("c9c1466a-bef5-4883-872e-bc59375f5d89"),
		// 									Encoding: to.Ptr("PEM"),
		// 									ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-16T20:50:16.753Z"); return t}()),
		// 									FilePaths: []*string{
		// 										to.Ptr("this/is/a/path/by/Mike/but/isExpired/and/could/get/pretty/long"),
		// 										to.Ptr("/this/is/another/path")},
		// 										Fingerprint: to.Ptr("eb:04:cf:5e:b1:f3:9a:fa:76:2f:2b:b1:20:f2:96:cb:a5:20:c1:b9:7d:b1:58:95:65:b8:1c:b9:a1:7b:72:44"),
		// 										IsExpired: to.Ptr(armiotfirmwaredefense.IsExpiredFalse),
		// 										IsSelfSigned: to.Ptr(armiotfirmwaredefense.IsSelfSignedTrue),
		// 										IsShortKeySize: to.Ptr(armiotfirmwaredefense.IsShortKeySizeFalse),
		// 										IsWeakSignature: to.Ptr(armiotfirmwaredefense.IsWeakSignatureFalse),
		// 										IssuedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-16T20:50:16.753Z"); return t}()),
		// 										Issuer: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 											CommonName: to.Ptr("VeriSign Class 3"),
		// 											Country: to.Ptr("United States"),
		// 											Organization: to.Ptr("Verisign, Inc."),
		// 											OrganizationalUnit: to.Ptr("(c) 1999 Verisign inc."),
		// 											State: to.Ptr("NY"),
		// 										},
		// 										KeyAlgorithm: to.Ptr("RSA Encryption"),
		// 										KeySize: to.Ptr[int64](2048),
		// 										Role: to.Ptr("Root CA"),
		// 										SerialNumber: to.Ptr("206684696279472310254277870180966723415"),
		// 										SignatureAlgorithm: to.Ptr("SHA-1 with RSA Encryption"),
		// 										Subject: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 											CommonName: to.Ptr("MikeSign Class 5"),
		// 											Country: to.Ptr("United States"),
		// 											Organization: to.Ptr("Mikesign, Inc."),
		// 											OrganizationalUnit: to.Ptr("(c) 2021 Mikesign inc."),
		// 											State: to.Ptr("MD"),
		// 										},
		// 										Usage: []*string{
		// 											to.Ptr("Crypt"),
		// 											to.Ptr("sign"),
		// 											to.Ptr("auth")},
		// 										},
		// 										{
		// 											Name: to.Ptr("PeterSign Class 1000"),
		// 											CryptoCertID: to.Ptr("0b580a77-228b-4786-b623-518ec8cdd564"),
		// 											Encoding: to.Ptr("PEM"),
		// 											ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2030-03-16T20:50:16.753Z"); return t}()),
		// 											FilePaths: []*string{
		// 												to.Ptr("this/is/a/path/for/Peters/puzzlers"),
		// 												to.Ptr("/this/is/another/path")},
		// 												Fingerprint: to.Ptr("eb:04:cf:5e:b1:f3:9a:fa:76:2f:2b:b1:20:f2:96:cb:a5:20:c1:b9:7d:b1:58:95:65:b8:1c:b9:a1:7b:72:44"),
		// 												IsExpired: to.Ptr(armiotfirmwaredefense.IsExpiredFalse),
		// 												IsSelfSigned: to.Ptr(armiotfirmwaredefense.IsSelfSignedTrue),
		// 												IsShortKeySize: to.Ptr(armiotfirmwaredefense.IsShortKeySizeFalse),
		// 												IsWeakSignature: to.Ptr(armiotfirmwaredefense.IsWeakSignatureFalse),
		// 												IssuedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-16T20:50:16.753Z"); return t}()),
		// 												Issuer: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 													CommonName: to.Ptr("VeriSign Class 3"),
		// 													Country: to.Ptr("United States"),
		// 													Organization: to.Ptr("Verisign, Inc."),
		// 													OrganizationalUnit: to.Ptr("(c) 1999 Verisign inc."),
		// 													State: to.Ptr("NY"),
		// 												},
		// 												KeyAlgorithm: to.Ptr("RSA Encryption"),
		// 												KeySize: to.Ptr[int64](2048),
		// 												Role: to.Ptr("Standard"),
		// 												SerialNumber: to.Ptr("206684696279472310254277870180966723415"),
		// 												SignatureAlgorithm: to.Ptr("SHA-1 with RSA Encryption"),
		// 												Subject: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 													CommonName: to.Ptr("PeterSign Class 1000"),
		// 													Country: to.Ptr("United States"),
		// 													Organization: to.Ptr("Petersign, Inc."),
		// 													OrganizationalUnit: to.Ptr("(c) 1999 Petersign inc."),
		// 													State: to.Ptr("MD"),
		// 												},
		// 												Usage: []*string{
		// 													to.Ptr("Crypt"),
		// 													to.Ptr("sign"),
		// 													to.Ptr("auth")},
		// 											}},
		// 										}
	}
}
