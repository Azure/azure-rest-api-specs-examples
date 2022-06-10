```go
package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-05-02-preview/examples/ManagedClustersResetAADProfile.json
func ExampleManagedClustersClient_BeginResetAADProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewManagedClustersClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginResetAADProfile(ctx,
		"rg1",
		"clustername1",
		armcontainerservice.ManagedClusterAADProfile{
			ClientAppID:     to.Ptr("clientappid"),
			ServerAppID:     to.Ptr("serverappid"),
			ServerAppSecret: to.Ptr("serverappsecret"),
			TenantID:        to.Ptr("tenantid"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerservice%2Farmcontainerservice%2Fv2.0.0-beta.1/sdk/resourcemanager/containerservice/armcontainerservice/README.md) on how to add the SDK to your project and authenticate.
