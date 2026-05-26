package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2025-09-01-preview/DefenderForStorage/ListDefenderForStorageSettings_example.json
func ExampleDefenderForStorageClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDefenderForStorageClient().NewListPager("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount", nil)
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
		// page = armsecurity.DefenderForStorageClientListResponse{
		// 	DefenderForStorageSettingList: armsecurity.DefenderForStorageSettingList{
		// 		Value: []*armsecurity.DefenderForStorageSetting{
		// 			{
		// 				Name: to.Ptr("current"),
		// 				Type: to.Ptr("Microsoft.Security/defenderForStorageSettings"),
		// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount/providers/Microsoft.Security/defenderForStorageSettings/current"),
		// 				Properties: &armsecurity.DefenderForStorageSettingProperties{
		// 					IsEnabled: to.Ptr(true),
		// 					MalwareScanning: &armsecurity.MalwareScanningProperties{
		// 						AutomatedResponse: to.Ptr(armsecurity.AutomatedResponseTypeBlobSoftDelete),
		// 						BlobScanResultsOptions: to.Ptr(armsecurity.BlobScanResultsOptionsBlobIndexTags),
		// 						OnUpload: &armsecurity.OnUploadProperties{
		// 							CapGBPerMonth: to.Ptr[int32](10000),
		// 							Filters: &armsecurity.OnUploadFilters{
		// 								ExcludeBlobsLargerThan: 1024,
		// 								ExcludeBlobsWithPrefix: []*string{
		// 									to.Ptr("sample-container/logs"),
		// 									to.Ptr("single-excluded-container/"),
		// 									to.Ptr("excluded-containers"),
		// 								},
		// 								ExcludeBlobsWithSuffix: []*string{
		// 									to.Ptr(".log"),
		// 									to.Ptr(".jpg"),
		// 								},
		// 							},
		// 							IsEnabled: to.Ptr(true),
		// 						},
		// 						ScanResultsEventGridTopicResourceID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.EventGrid/topics/sampletopic"),
		// 					},
		// 					OverrideSubscriptionLevelSettings: to.Ptr(true),
		// 					SensitiveDataDiscovery: &armsecurity.SensitiveDataDiscoveryProperties{
		// 						IsEnabled: to.Ptr(true),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
