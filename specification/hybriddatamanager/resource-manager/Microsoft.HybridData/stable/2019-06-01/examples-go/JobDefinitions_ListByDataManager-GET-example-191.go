package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/JobDefinitions_ListByDataManager-GET-example-191.json
func ExampleJobDefinitionsClient_NewListByDataManagerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybriddatamanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobDefinitionsClient().NewListByDataManagerPager("ResourceGroupForSDKTest", "TestAzureSDKOperations", &armhybriddatamanager.JobDefinitionsClientListByDataManagerOptions{Filter: nil})
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
		// page.JobDefinitionList = armhybriddatamanager.JobDefinitionList{
		// 	Value: []*armhybriddatamanager.JobDefinition{
		// 		{
		// 			Name: to.Ptr("jobdeffromtestcode1"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers/dataServices/jobDefinitions"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/DataTransformation/jobDefinitions/jobdeffromtestcode1"),
		// 			Properties: &armhybriddatamanager.JobDefinitionProperties{
		// 				DataServiceInput: map[string]any{
		// 					"AzureStorageType": "Blob",
		// 					"BackupChoice": "UseExistingLatest",
		// 					"ContainerName": "containerfromtest",
		// 					"DeviceName": "8600-SHG0997877L71FC",
		// 					"FileNameFilter": "*",
		// 					"IsDirectoryMode": false,
		// 					"RootDirectories":[]any{
		// 						"\\",
		// 					},
		// 					"VolumeNames":[]any{
		// 						"TestAutomation",
		// 					},
		// 				},
		// 				DataSinkID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStores/TestAzureStorage1"),
		// 				DataSourceID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStores/TestStorSimpleSource1"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-05T08:51:43.366Z"); return t}()),
		// 				RunLocation: to.Ptr(armhybriddatamanager.RunLocationWestus),
		// 				Schedules: []*armhybriddatamanager.Schedule{
		// 				},
		// 				State: to.Ptr(armhybriddatamanager.StateEnabled),
		// 				UserConfirmation: to.Ptr(armhybriddatamanager.UserConfirmationRequired),
		// 			},
		// 	}},
		// }
	}
}
