package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/JobDefinitions_CreateOrUpdate-PUT-example-83.json
func ExampleJobDefinitionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhybriddatamanager.NewJobDefinitionsClient("6e0219f5-327a-4365-904f-05eed4227ad7", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"DataTransformation",
		"jobdeffromtestcode1",
		"ResourceGroupForSDKTest",
		"TestAzureSDKOperations",
		armhybriddatamanager.JobDefinition{
			Properties: &armhybriddatamanager.JobDefinitionProperties{
				DataServiceInput: map[string]interface{}{
					"AzureStorageType": "Blob",
					"BackupChoice":     "UseExistingLatest",
					"ContainerName":    "containerfromtest",
					"DeviceName":       "8600-SHG0997877L71FC",
					"FileNameFilter":   "*",
					"IsDirectoryMode":  false,
					"RootDirectories": []interface{}{
						"\\",
					},
					"VolumeNames": []interface{}{
						"TestAutomation",
					},
				},
				DataSinkID:       to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStores/TestAzureStorage1"),
				DataSourceID:     to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStores/TestStorSimpleSource1"),
				RunLocation:      to.Ptr(armhybriddatamanager.RunLocationWestus),
				State:            to.Ptr(armhybriddatamanager.StateEnabled),
				UserConfirmation: to.Ptr(armhybriddatamanager.UserConfirmationRequired),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
