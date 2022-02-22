Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmresources%2Fv0.3.1/sdk/resourcemanager/resources/armresources/README.md) on how to add the SDK to your project and authenticate.

```go
package armresources_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PutDeploymentAtTenant.json
func ExampleDeploymentsClient_BeginCreateOrUpdateAtTenantScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdateAtTenantScope(ctx,
		"<deployment-name>",
		armresources.ScopedDeployment{
			Location: to.StringPtr("<location>"),
			Properties: &armresources.DeploymentProperties{
				Mode:       armresources.DeploymentModeIncremental.ToPtr(),
				Parameters: map[string]interface{}{},
				TemplateLink: &armresources.TemplateLink{
					URI: to.StringPtr("<uri>"),
				},
			},
			Tags: map[string]*string{
				"tagKey1": to.StringPtr("tag-value-1"),
				"tagKey2": to.StringPtr("tag-value-2"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DeploymentsClientCreateOrUpdateAtTenantScopeResult)
}
```
