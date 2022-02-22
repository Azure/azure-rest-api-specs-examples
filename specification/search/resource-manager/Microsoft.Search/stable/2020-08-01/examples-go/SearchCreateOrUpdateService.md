Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsearch%2Farmsearch%2Fv0.3.1/sdk/resourcemanager/search/armsearch/README.md) on how to add the SDK to your project and authenticate.

```go
package armsearch_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchCreateOrUpdateService.json
func ExampleServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsearch.NewServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<search-service-name>",
		armsearch.Service{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"app-name": to.StringPtr("My e-commerce app"),
			},
			Properties: &armsearch.ServiceProperties{
				HostingMode:    armsearch.HostingModeDefault.ToPtr(),
				PartitionCount: to.Int32Ptr(1),
				ReplicaCount:   to.Int32Ptr(3),
			},
			SKU: &armsearch.SKU{
				Name: armsearch.SKUNameStandard.ToPtr(),
			},
		},
		&armsearch.SearchManagementRequestOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServicesClientCreateOrUpdateResult)
}
```
