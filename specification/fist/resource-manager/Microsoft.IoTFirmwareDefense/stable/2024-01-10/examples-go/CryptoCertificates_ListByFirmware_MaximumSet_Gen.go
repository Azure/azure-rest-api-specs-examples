package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/CryptoCertificates_ListByFirmware_MaximumSet_Gen.json
func ExampleCryptoCertificatesClient_NewListByFirmwarePager_cryptoCertificatesListByFirmwareMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCryptoCertificatesClient().NewListByFirmwarePager("FirmwareAnalysisRG", "default", "109a9886-50bf-85a8-9d75-000000000000", nil)
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
		// page.CryptoCertificateListResult = armiotfirmwaredefense.CryptoCertificateListResult{
		// 	Value: []*armiotfirmwaredefense.CryptoCertificateResource{
		// 		{
		// 			Name: to.Ptr("d402993a-eca1-4ef4-a13d-000000000000"),
		// 			Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares/cryptoCertificates"),
		// 			ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/cryptoCertificates/d402993a-eca1-4ef4-a13d-000000000000"),
		// 			Properties: &armiotfirmwaredefense.CryptoCertificate{
		// 				Name: to.Ptr("PolarSSL Test Client 2"),
		// 				CryptoCertID: to.Ptr("d402993a-eca1-4ef4-a13d-000000000000"),
		// 				Encoding: to.Ptr("DER"),
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2035-09-22T15:52:04.000Z"); return t}()),
		// 				FilePaths: []*string{
		// 					to.Ptr("/path/to/keys-certs/ntgr_ra_aws")},
		// 					Fingerprint: to.Ptr("EA:70:F9:B4:45:F3:41:BC:33:22:34:A4:98:DD:69:04:1D:E8:82:05"),
		// 					IsExpired: to.Ptr(false),
		// 					IsSelfSigned: to.Ptr(false),
		// 					IsShortKeySize: to.Ptr(false),
		// 					IssuedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-24T15:52:04.000Z"); return t}()),
		// 					Issuer: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 						CommonName: to.Ptr("Polarssl Test EC CA"),
		// 						Country: to.Ptr("NL"),
		// 						Organization: to.Ptr("PolarSSL"),
		// 						OrganizationalUnit: to.Ptr(""),
		// 						State: to.Ptr(""),
		// 					},
		// 					KeyAlgorithm: to.Ptr("ECC"),
		// 					KeySize: to.Ptr[int64](256),
		// 					PairedKey: &armiotfirmwaredefense.PairedKey{
		// 						Type: to.Ptr("Private"),
		// 						ID: to.Ptr("cd1bdece-0576-4ff5-8039-000000000000"),
		// 					},
		// 					Role: to.Ptr("Non-CA"),
		// 					SerialNumber: to.Ptr("0D"),
		// 					SignatureAlgorithm: to.Ptr("sha256ECDSA"),
		// 					Subject: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 						CommonName: to.Ptr("PolarSSL Test Client 2"),
		// 						Country: to.Ptr("NL"),
		// 						Organization: to.Ptr("PolarSSL"),
		// 						OrganizationalUnit: to.Ptr(""),
		// 						State: to.Ptr(""),
		// 					},
		// 					Usage: []*string{
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("78fd7929-fac1-45cd-bfc6-000000000000"),
		// 				Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares/cryptoCertificates"),
		// 				ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/cryptoCertificates/78fd7929-fac1-45cd-bfc6-000000000000"),
		// 				Properties: &armiotfirmwaredefense.CryptoCertificate{
		// 					Name: to.Ptr("DigiCert SHA2 Secure Server CA"),
		// 					CryptoCertID: to.Ptr("78fd7929-fac1-45cd-bfc6-000000000000"),
		// 					Encoding: to.Ptr("DER"),
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-08T12:00:00.000Z"); return t}()),
		// 					FilePaths: []*string{
		// 						to.Ptr("/path/to/keys-certs/revoke.pem"),
		// 						to.Ptr("/path/to/keys-certs/sample.crt")},
		// 						Fingerprint: to.Ptr("1F:B8:6B:11:68:EC:74:31:54:06:2E:8C:9C:C5:B1:71:A4:B7:CC:B4"),
		// 						IsExpired: to.Ptr(true),
		// 						IsSelfSigned: to.Ptr(false),
		// 						IsShortKeySize: to.Ptr(false),
		// 						IssuedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2013-03-08T12:00:00.000Z"); return t}()),
		// 						Issuer: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 							CommonName: to.Ptr("DigiCert Global Root CA"),
		// 							Country: to.Ptr("US"),
		// 							Organization: to.Ptr("DigiCert Inc"),
		// 							OrganizationalUnit: to.Ptr("www.digicert.com"),
		// 							State: to.Ptr(""),
		// 						},
		// 						KeyAlgorithm: to.Ptr("RSA"),
		// 						KeySize: to.Ptr[int64](2048),
		// 						PairedKey: &armiotfirmwaredefense.PairedKey{
		// 						},
		// 						Role: to.Ptr("Intermediate"),
		// 						SerialNumber: to.Ptr("01:FD:A3:EB:6E:CA:75:C8:88:43:8B:72:4B:CF:BC:91"),
		// 						SignatureAlgorithm: to.Ptr("sha256RSA"),
		// 						Subject: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 							CommonName: to.Ptr("DigiCert SHA2 Secure Server CA"),
		// 							Country: to.Ptr("US"),
		// 							Organization: to.Ptr("DigiCert Inc"),
		// 							OrganizationalUnit: to.Ptr(""),
		// 							State: to.Ptr(""),
		// 						},
		// 						Usage: []*string{
		// 							to.Ptr("CrlSign"),
		// 							to.Ptr("KeyCertSign"),
		// 							to.Ptr("DigitalSignature")},
		// 						},
		// 				}},
		// 			}
	}
}
