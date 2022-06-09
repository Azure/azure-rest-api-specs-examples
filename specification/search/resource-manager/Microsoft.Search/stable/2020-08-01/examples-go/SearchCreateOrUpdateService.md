```go
package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchCreateOrUpdateService.json
func ExampleServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsearch.NewServicesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"mysearchservice",
		armsearch.Service{
			Location: to.Ptr("westus"),
			Tags: map[string]*string{
				"app-name": to.Ptr("My e-commerce app"),
			},
			Properties: &armsearch.ServiceProperties{
				HostingMode:    to.Ptr(armsearch.HostingModeDefault),
				PartitionCount: to.Ptr[int32](1),
				ReplicaCount:   to.Ptr[int32](3),
			},
			SKU: &armsearch.SKU{
				Name: to.Ptr(armsearch.SKUNameStandard),
			},
		},
		&armsearch.SearchManagementRequestOptions{ClientRequestID: nil},
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsearch%2Farmsearch%2Fv1.0.0/sdk/resourcemanager/search/armsearch/README.md) on how to add the SDK to your project and authenticate.
