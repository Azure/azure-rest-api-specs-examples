Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fautomation%2Farmautomation%2Fv0.5.0/sdk/resourcemanager/automation/armautomation/README.md) on how to add the SDK to your project and authenticate.

```go
package armautomation_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateDscNodeConfiguration.json
func ExampleDscNodeConfigurationClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armautomation.NewDscNodeConfigurationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<node-configuration-name>",
		armautomation.DscNodeConfigurationCreateOrUpdateParameters{
			Name: to.Ptr("<name>"),
			Properties: &armautomation.DscNodeConfigurationCreateOrUpdateParametersProperties{
				Configuration: &armautomation.DscConfigurationAssociationProperty{
					Name: to.Ptr("<name>"),
				},
				IncrementNodeConfigurationBuild: to.Ptr(true),
				Source: &armautomation.ContentSource{
					Type: to.Ptr(armautomation.ContentSourceTypeEmbeddedContent),
					Hash: &armautomation.ContentHash{
						Algorithm: to.Ptr("<algorithm>"),
						Value:     to.Ptr("<value>"),
					},
					Value:   to.Ptr("<value>"),
					Version: to.Ptr("<version>"),
				},
			},
		},
		&armautomation.DscNodeConfigurationClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```
