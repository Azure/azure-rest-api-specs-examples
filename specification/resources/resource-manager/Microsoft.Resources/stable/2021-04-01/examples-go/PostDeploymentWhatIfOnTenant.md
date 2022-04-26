Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmresources%2Fv0.5.0/sdk/resourcemanager/resources/armresources/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PostDeploymentWhatIfOnTenant.json
func ExampleDeploymentsClient_BeginWhatIfAtTenantScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginWhatIfAtTenantScope(ctx,
		"<deployment-name>",
		armresources.ScopedDeploymentWhatIf{
			Location: to.Ptr("<location>"),
			Properties: &armresources.DeploymentWhatIfProperties{
				Mode:         to.Ptr(armresources.DeploymentModeIncremental),
				Parameters:   map[string]interface{}{},
				TemplateLink: &armresources.TemplateLink{},
			},
		},
		&armresources.DeploymentsClientBeginWhatIfAtTenantScopeOptions{ResumeToken: ""})
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
