package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/CryptoKeys_ListByFirmware_MaximumSet_Gen.json
func ExampleCryptoKeysClient_NewListByFirmwarePager_cryptoKeysListByFirmwareMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCryptoKeysClient().NewListByFirmwarePager("FirmwareAnalysisRG", "default", "109a9886-50bf-85a8-9d75-000000000000", nil)
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
		// page.CryptoKeyListResult = armiotfirmwaredefense.CryptoKeyListResult{
		// 	Value: []*armiotfirmwaredefense.CryptoKeyResource{
		// 		{
		// 			Name: to.Ptr("cd1bdece-0576-4ff5-8039-000000000000"),
		// 			Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares/cryptoKeys"),
		// 			ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/cryptoKeys/cd1bdece-0576-4ff5-8039-000000000000"),
		// 			Properties: &armiotfirmwaredefense.CryptoKey{
		// 				CryptoKeyID: to.Ptr("cd1bdece-0576-4ff5-8039-000000000000"),
		// 				FilePaths: []*string{
		// 					to.Ptr("/path/to/keys-certs/ntgr_ra_aws")},
		// 					IsShortKeySize: to.Ptr(false),
		// 					KeyAlgorithm: to.Ptr("ECDsa"),
		// 					KeySize: to.Ptr[int64](256),
		// 					KeyType: to.Ptr("Private"),
		// 					PairedKey: &armiotfirmwaredefense.PairedKey{
		// 						Type: to.Ptr("Public"),
		// 						ID: to.Ptr("95f41328-cc3b-41a1-bd70-000000000000"),
		// 					},
		// 					Usage: []*string{
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("5adeaa85-f30a-4295-9b4a-000000000000"),
		// 				Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares/cryptoKeys"),
		// 				ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/cryptoKeys/5adeaa85-f30a-4295-9b4a-000000000000"),
		// 				Properties: &armiotfirmwaredefense.CryptoKey{
		// 					CryptoKeyID: to.Ptr("5adeaa85-f30a-4295-9b4a-000000000000"),
		// 					FilePaths: []*string{
		// 						to.Ptr("/path/to/keys-certs/addon.key")},
		// 						IsShortKeySize: to.Ptr(true),
		// 						KeyAlgorithm: to.Ptr("RSA"),
		// 						KeySize: to.Ptr[int64](1024),
		// 						KeyType: to.Ptr("Public"),
		// 						PairedKey: &armiotfirmwaredefense.PairedKey{
		// 						},
		// 						Usage: []*string{
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
