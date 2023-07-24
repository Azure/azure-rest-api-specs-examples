package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_ListGenerateCryptoKeyList_MaximumSet_Gen.json
func ExampleFirmwareClient_NewListGenerateCryptoKeyListPager_firmwareListGenerateCryptoKeyListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirmwareClient().NewListGenerateCryptoKeyListPager("FirmwareAnalysisRG", "default", "DECAFBAD-0000-0000-0000-BADBADBADBAD", nil)
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
		// page.CryptoKeyList = armiotfirmwaredefense.CryptoKeyList{
		// 	Value: []*armiotfirmwaredefense.CryptoKey{
		// 		{
		// 			CryptoKeyID: to.Ptr("09c97cb7-0963-41ad-b9e2-0be7b10ac1bf"),
		// 			FilePaths: []*string{
		// 				to.Ptr("this/is/a/file/path"),
		// 				to.Ptr("this/is/another/file/path")},
		// 				IsShortKeySize: to.Ptr(armiotfirmwaredefense.IsShortKeySizeFalse),
		// 				KeyAlgorithm: to.Ptr("ECDSA"),
		// 				KeySize: to.Ptr[int64](256),
		// 				KeyType: to.Ptr("Public"),
		// 				PairedKey: &armiotfirmwaredefense.PairedKey{
		// 					Type: to.Ptr("Certificate"),
		// 					AdditionalProperties: map[string]any{
		// 						"displayName": "VeriSign Class 2",
		// 					},
		// 					ID: to.Ptr("09c97cb7-0963-41ad-b9e2-0be7b10ac1bf"),
		// 				},
		// 				Usage: []*string{
		// 				},
		// 			},
		// 			{
		// 				CryptoKeyID: to.Ptr("98ff3d6c-bb76-4037-b73e-89b3b47efd71"),
		// 				FilePaths: []*string{
		// 					to.Ptr("this/is/a/file/path_pub.pem")},
		// 					IsShortKeySize: to.Ptr(armiotfirmwaredefense.IsShortKeySizeFalse),
		// 					KeyAlgorithm: to.Ptr("RSA"),
		// 					KeySize: to.Ptr[int64](2048),
		// 					KeyType: to.Ptr("Public"),
		// 					PairedKey: &armiotfirmwaredefense.PairedKey{
		// 						Type: to.Ptr("Private Key"),
		// 						ID: to.Ptr("adfe0c4c-7d23-4f4f-a8bb-e65b8f455287"),
		// 					},
		// 					Usage: []*string{
		// 						to.Ptr("SSH")},
		// 					},
		// 					{
		// 						CryptoKeyID: to.Ptr("1beb6fac-1cfe-431c-9437-ea8cfa9e6aaa"),
		// 						FilePaths: []*string{
		// 						},
		// 						IsShortKeySize: to.Ptr(armiotfirmwaredefense.IsShortKeySizeTrue),
		// 						KeyAlgorithm: to.Ptr("ECDSA"),
		// 						KeySize: to.Ptr[int64](128),
		// 						KeyType: to.Ptr("Public"),
		// 						Usage: []*string{
		// 						},
		// 					},
		// 					{
		// 						CryptoKeyID: to.Ptr("adfe0c4c-7d23-4f4f-a8bb-e65b8f455287"),
		// 						FilePaths: []*string{
		// 							to.Ptr("this/is/a/file/path.pem")},
		// 							IsShortKeySize: to.Ptr(armiotfirmwaredefense.IsShortKeySizeFalse),
		// 							KeyAlgorithm: to.Ptr("RSA"),
		// 							KeySize: to.Ptr[int64](2048),
		// 							KeyType: to.Ptr("Private"),
		// 							PairedKey: &armiotfirmwaredefense.PairedKey{
		// 								Type: to.Ptr("Public Key"),
		// 								ID: to.Ptr("98ff3d6c-bb76-4037-b73e-89b3b47efd71"),
		// 							},
		// 							Usage: []*string{
		// 								to.Ptr("SSH")},
		// 							},
		// 							{
		// 								CryptoKeyID: to.Ptr("8c8d1003-efb2-4c54-83d6-16dc6a5bbfd5"),
		// 								FilePaths: []*string{
		// 								},
		// 								IsShortKeySize: to.Ptr(armiotfirmwaredefense.IsShortKeySizeTrue),
		// 								KeyAlgorithm: to.Ptr("RSA"),
		// 								KeySize: to.Ptr[int64](1024),
		// 								KeyType: to.Ptr("Private"),
		// 								Usage: []*string{
		// 								},
		// 							},
		// 							{
		// 								CryptoKeyID: to.Ptr("0b04fc4c-bfbd-491b-b4c8-e93b0b0b28e9"),
		// 								FilePaths: []*string{
		// 									to.Ptr("this/is/a/file/path"),
		// 									to.Ptr("this/is/another/file/path")},
		// 									IsShortKeySize: to.Ptr(armiotfirmwaredefense.IsShortKeySizeFalse),
		// 									KeyAlgorithm: to.Ptr("MD5"),
		// 									KeySize: to.Ptr[int64](256),
		// 									KeyType: to.Ptr("Private"),
		// 									Usage: []*string{
		// 									},
		// 							}},
		// 						}
	}
}
