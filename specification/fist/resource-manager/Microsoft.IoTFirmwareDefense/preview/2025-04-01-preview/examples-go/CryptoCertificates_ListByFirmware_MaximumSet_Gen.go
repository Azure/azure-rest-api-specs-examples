package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
)

// Generated from example definition: 2025-04-01-preview/CryptoCertificates_ListByFirmware_MaximumSet_Gen.json
func ExampleCryptoCertificatesClient_NewListByFirmwarePager_cryptoCertificatesListByFirmwareMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("5C707B5F-6130-4F71-819E-953A28942E88", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCryptoCertificatesClient().NewListByFirmwarePager("rgiotfirmwaredefense", "exampleWorkspaceName", "00000000-0000-0000-0000-000000000000", nil)
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
		// page = armiotfirmwaredefense.CryptoCertificatesClientListByFirmwareResponse{
		// 	CryptoCertificateResourceListResult: armiotfirmwaredefense.CryptoCertificateResourceListResult{
		// 		Value: []*armiotfirmwaredefense.CryptoCertificateResource{
		// 			{
		// 				Properties: &armiotfirmwaredefense.CryptoCertificate{
		// 					CryptoCertID: to.Ptr("iillnitoieyjjhhymaggp"),
		// 					Subject: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 						CommonName: to.Ptr("as"),
		// 						Organization: to.Ptr("uqaaxitzvimayheiixojfvtdkohwux"),
		// 						OrganizationalUnit: to.Ptr("tjppbejdpjwndktgb"),
		// 						State: to.Ptr("wptcsidcfwobyroygepitbmopuqyty"),
		// 						Country: to.Ptr("uhztskdgmxwzxqrucsmycqwema"),
		// 					},
		// 					Issuer: &armiotfirmwaredefense.CryptoCertificateEntity{
		// 						CommonName: to.Ptr("as"),
		// 						Organization: to.Ptr("uqaaxitzvimayheiixojfvtdkohwux"),
		// 						OrganizationalUnit: to.Ptr("tjppbejdpjwndktgb"),
		// 						State: to.Ptr("wptcsidcfwobyroygepitbmopuqyty"),
		// 						Country: to.Ptr("uhztskdgmxwzxqrucsmycqwema"),
		// 					},
		// 					IssuedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:47.274Z"); return t}()),
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:47.274Z"); return t}()),
		// 					SignatureAlgorithm: to.Ptr("yvnrjznhumg"),
		// 					Encoding: to.Ptr("n"),
		// 					SerialNumber: to.Ptr("tjvbzogfirohdwpckavvosyyevpar"),
		// 					Fingerprint: to.Ptr("bjidncbwvrxfkliddw"),
		// 					FilePaths: []*string{
		// 						to.Ptr("tbtijbpxccby"),
		// 					},
		// 					PairedKey: &armiotfirmwaredefense.PairedKey{
		// 						Type: to.Ptr("tvdhzzrieepfxmsfymzky"),
		// 						PairedKeyID: to.Ptr("krrhvprzbaccizagbobm"),
		// 					},
		// 					IsExpired: to.Ptr(true),
		// 					IsSelfSigned: to.Ptr(true),
		// 					IsWeakSignature: to.Ptr(true),
		// 					IsShortKeySize: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armiotfirmwaredefense.ProvisioningStateSucceeded),
		// 					CertificateName: to.Ptr("mxcalihubqunqspgskjt"),
		// 					CertificateRole: to.Ptr("bwjvtjizfmckztcedqmnzzaw"),
		// 					CertificateKeySize: to.Ptr[int64](4),
		// 					CertificateKeyAlgorithm: to.Ptr("xnsqvapvllvbwkckerpoef"),
		// 					CertificateUsage: []*armiotfirmwaredefense.CertificateUsage{
		// 						to.Ptr(armiotfirmwaredefense.CertificateUsageDigitalSignature),
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/summaries/firmware"),
		// 				Name: to.Ptr("cwzjayiki"),
		// 				Type: to.Ptr("vmrrtwzfdqtaaqyu"),
		// 				SystemData: &armiotfirmwaredefense.SystemData{
		// 					CreatedBy: to.Ptr("nqisshvdzqcxzbujvacin"),
		// 					CreatedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("of"),
		// 					LastModifiedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
