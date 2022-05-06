Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flogic%2Farmlogic%2Fv0.5.0/sdk/resourcemanager/logic/armlogic/README.md) on how to add the SDK to your project and authenticate.

```go
package armlogic_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Put.json
func ExampleIntegrationServiceEnvironmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armlogic.NewIntegrationServiceEnvironmentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group>",
		"<integration-service-environment-name>",
		armlogic.IntegrationServiceEnvironment{
			Location: to.Ptr("<location>"),
			Properties: &armlogic.IntegrationServiceEnvironmentProperties{
				EncryptionConfiguration: &armlogic.IntegrationServiceEnvironmenEncryptionConfiguration{
					EncryptionKeyReference: &armlogic.IntegrationServiceEnvironmenEncryptionKeyReference{
						KeyName: to.Ptr("<key-name>"),
						KeyVault: &armlogic.ResourceReference{
							ID: to.Ptr("<id>"),
						},
						KeyVersion: to.Ptr("<key-version>"),
					},
				},
				NetworkConfiguration: &armlogic.NetworkConfiguration{
					AccessEndpoint: &armlogic.IntegrationServiceEnvironmentAccessEndpoint{
						Type: to.Ptr(armlogic.IntegrationServiceEnvironmentAccessEndpointTypeInternal),
					},
					Subnets: []*armlogic.ResourceReference{
						{
							ID: to.Ptr("<id>"),
						},
						{
							ID: to.Ptr("<id>"),
						},
						{
							ID: to.Ptr("<id>"),
						},
						{
							ID: to.Ptr("<id>"),
						}},
				},
			},
			SKU: &armlogic.IntegrationServiceEnvironmentSKU{
				Name:     to.Ptr(armlogic.IntegrationServiceEnvironmentSKUNamePremium),
				Capacity: to.Ptr[int32](2),
			},
		},
		&armlogic.IntegrationServiceEnvironmentsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
