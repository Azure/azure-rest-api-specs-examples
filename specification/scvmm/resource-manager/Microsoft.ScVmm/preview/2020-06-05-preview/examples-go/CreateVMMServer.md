Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fscvmm%2Farmscvmm%2Fv0.2.0/sdk/resourcemanager/scvmm/armscvmm/README.md) on how to add the SDK to your project and authenticate.

```go
package armscvmm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateVMMServer.json
func ExampleVmmServersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armscvmm.NewVmmServersClient("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testrg",
		"ContosoVMMServer",
		armscvmm.VMMServer{
			ExtendedLocation: &armscvmm.ExtendedLocation{
				Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
				Type: to.Ptr("customLocation"),
			},
			Location: to.Ptr("East US"),
			Properties: &armscvmm.VMMServerProperties{
				Credentials: &armscvmm.VMMServerPropertiesCredentials{
					Password: to.Ptr("password"),
					Username: to.Ptr("testuser"),
				},
				Fqdn: to.Ptr("VMM.contoso.com"),
				Port: to.Ptr[int32](1234),
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
```
