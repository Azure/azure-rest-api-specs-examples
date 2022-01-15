Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmresources%2Fv0.3.0/sdk/resourcemanager/resources/armresources/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PostDeploymentWhatIfOnSubscription.json
func ExampleDeploymentsClient_BeginWhatIfAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginWhatIfAtSubscriptionScope(ctx,
		"<deployment-name>",
		armresources.DeploymentWhatIf{
			Location: to.StringPtr("<location>"),
			Properties: &armresources.DeploymentWhatIfProperties{
				Mode:         armresources.DeploymentModeIncremental.ToPtr(),
				Parameters:   map[string]interface{}{},
				TemplateLink: &armresources.TemplateLink{},
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
	log.Printf("Response result: %#v\n", res.DeploymentsClientWhatIfAtSubscriptionScopeResult)
}
```
