package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/DataStores_CreateOrUpdate_DataSink-PUT-example-162.json
func ExampleDataStoresClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhybriddatamanager.NewDataStoresClient("6e0219f5-327a-4365-904f-05eed4227ad7", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"TestAzureStorage1",
		"ResourceGroupForSDKTest",
		"TestAzureSDKOperations",
		armhybriddatamanager.DataStore{
			Properties: &armhybriddatamanager.DataStoreProperties{
				CustomerSecrets: []*armhybriddatamanager.CustomerSecret{
					{
						Algorithm:     to.Ptr(armhybriddatamanager.SupportedAlgorithmRSA15),
						KeyIdentifier: to.Ptr("StorageAccountAccessKey"),
						KeyValue:      to.Ptr("Of4H9eF03G8QuxvkZQEbFWv3YdN3U//WugzuqReQekbXXQyg+QSicVKrwSOOKVi1zWMYGbKg7d5/ES2gdz+O5ZEw89bvE4mJD/wQmkIsqhPnbN0gyVK6nZePXVUU1A+UzjLfvhSA6KyUQfzNAZ5/TLt6fo1JyQrKTtkvnkLFyfv1AqBZ+dw8JK3RZi/rEN8HD3R3qsBwUYfyEuGLGiujy2CGrr/1uPiUVMR6QuFWRsjm39eMSHa4maLg4tQ0IY/jIy8rMlx3KjF3CcCbPzAqEq5vXy37wkjZbus771te1gLSrzcpVKIMg4DrmgaoJ02jAu+izBjNgLXAFPSUneQ8yw==:ezMkh4PMhCnjJtYkpTaP0SdblP5VAeRe4glW2PgIzICHw3T8ZyGDoaTrCv4/m5wtcEhWdtxhta+j1MQWlK5MIA+hvf8QjIDIjQv696ov5y+pcFe/upd2ekGOei7FCwB2u7I8WnkAtIKTUkf6eDQBZXm26DjfG1Dlc+Mjjq+AerukEa6WpOyqrD7Qub26Pgmj4AsuBx19X1EAmTZacubkoiNagXM8V+IDanHOhLMvfgQ7rw8oZhWfofxi4m+eJqjOXXaqSyorNK8UEcqP6P9pDP8AN8ulXEx6rZy2B5Oi0vSV+wlRLbUuQslga4ItOGxctW/ZX8uWozt+5A3k4URt6A=="),
					},
					{
						Algorithm:     to.Ptr(armhybriddatamanager.SupportedAlgorithmRSA15),
						KeyIdentifier: to.Ptr("StorageAccountAccessKeyForQueue"),
						KeyValue:      to.Ptr("Of4H9eF03G8QuxvkZQEbFWv3YdN3U//WugzuqReQekbXXQyg+QSicVKrwSOOKVi1zWMYGbKg7d5/ES2gdz+O5ZEw89bvE4mJD/wQmkIsqhPnbN0gyVK6nZePXVUU1A+UzjLfvhSA6KyUQfzNAZ5/TLt6fo1JyQrKTtkvnkLFyfv1AqBZ+dw8JK3RZi/rEN8HD3R3qsBwUYfyEuGLGiujy2CGrr/1uPiUVMR6QuFWRsjm39eMSHa4maLg4tQ0IY/jIy8rMlx3KjF3CcCbPzAqEq5vXy37wkjZbus771te1gLSrzcpVKIMg4DrmgaoJ02jAu+izBjNgLXAFPSUneQ8yw==:ezMkh4PMhCnjJtYkpTaP0SdblP5VAeRe4glW2PgIzICHw3T8ZyGDoaTrCv4/m5wtcEhWdtxhta+j1MQWlK5MIA+hvf8QjIDIjQv696ov5y+pcFe/upd2ekGOei7FCwB2u7I8WnkAtIKTUkf6eDQBZXm26DjfG1Dlc+Mjjq+AerukEa6WpOyqrD7Qub26Pgmj4AsuBx19X1EAmTZacubkoiNagXM8V+IDanHOhLMvfgQ7rw8oZhWfofxi4m+eJqjOXXaqSyorNK8UEcqP6P9pDP8AN8ulXEx6rZy2B5Oi0vSV+wlRLbUuQslga4ItOGxctW/ZX8uWozt+5A3k4URt6A=="),
					}},
				DataStoreTypeID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/AzureStorageAccount"),
				ExtendedProperties: map[string]interface{}{
					"extendedSaKey":              nil,
					"extendedSaName":             "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink",
					"storageAccountNameForQueue": "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink",
				},
				RepositoryID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink"),
				State:        to.Ptr(armhybriddatamanager.StateEnabled),
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
