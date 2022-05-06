Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybriddatamanager%2Farmhybriddatamanager%2Fv0.4.0/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybriddatamanager_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/DataStores_CreateOrUpdate_DataSink-PUT-example-162.json
func ExampleDataStoresClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhybriddatamanager.NewDataStoresClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<data-store-name>",
		"<resource-group-name>",
		"<data-manager-name>",
		armhybriddatamanager.DataStore{
			Properties: &armhybriddatamanager.DataStoreProperties{
				CustomerSecrets: []*armhybriddatamanager.CustomerSecret{
					{
						Algorithm:     to.Ptr(armhybriddatamanager.SupportedAlgorithmRSA15),
						KeyIdentifier: to.Ptr("<key-identifier>"),
						KeyValue:      to.Ptr("<key-value>"),
					},
					{
						Algorithm:     to.Ptr(armhybriddatamanager.SupportedAlgorithmRSA15),
						KeyIdentifier: to.Ptr("<key-identifier>"),
						KeyValue:      to.Ptr("<key-value>"),
					}},
				DataStoreTypeID: to.Ptr("<data-store-type-id>"),
				ExtendedProperties: map[string]interface{}{
					"extendedSaKey":              nil,
					"extendedSaName":             "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink",
					"storageAccountNameForQueue": "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink",
				},
				RepositoryID: to.Ptr("<repository-id>"),
				State:        to.Ptr(armhybriddatamanager.StateEnabled),
			},
		},
		&armhybriddatamanager.DataStoresClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
