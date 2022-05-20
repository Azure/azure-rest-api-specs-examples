Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappplatform%2Farmappplatform%2Fv1.1.0-beta.1/sdk/resourcemanager/appplatform/armappplatform/README.md) on how to add the SDK to your project and authenticate.

```go
package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-05-01-preview/examples/BuildpackBinding_CreateOrUpdate.json
func ExampleBuildpackBindingClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewBuildpackBindingClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myservice",
		"default",
		"default",
		"myBuildpackBinding",
		armappplatform.BuildpackBindingResource{
			Properties: &armappplatform.BuildpackBindingProperties{
				BindingType: to.Ptr(armappplatform.BindingTypeApplicationInsights),
				LaunchProperties: &armappplatform.BuildpackBindingLaunchProperties{
					Properties: map[string]*string{
						"abc":           to.Ptr("def"),
						"any-string":    to.Ptr("any-string"),
						"sampling-rate": to.Ptr("12.0"),
					},
					Secrets: map[string]*string{
						"connection-string": to.Ptr("XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXX;XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXXXXXXXX"),
					},
				},
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
