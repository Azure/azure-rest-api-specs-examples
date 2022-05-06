Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fadvisor%2Farmadvisor%2Fv0.4.0/sdk/resourcemanager/advisor/armadvisor/README.md) on how to add the SDK to your project and authenticate.

```go
package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/CreateConfiguration.json
func ExampleConfigurationsClient_CreateInResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armadvisor.NewConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateInResourceGroup(ctx,
		armadvisor.ConfigurationNameDefault,
		"<resource-group>",
		armadvisor.ConfigData{
			Properties: &armadvisor.ConfigDataProperties{
				Digests: []*armadvisor.DigestConfig{
					{
						Name:                  to.Ptr("<name>"),
						ActionGroupResourceID: to.Ptr("<action-group-resource-id>"),
						Categories: []*armadvisor.Category{
							to.Ptr(armadvisor.CategoryHighAvailability),
							to.Ptr(armadvisor.CategorySecurity),
							to.Ptr(armadvisor.CategoryPerformance),
							to.Ptr(armadvisor.CategoryCost),
							to.Ptr(armadvisor.CategoryOperationalExcellence)},
						Frequency: to.Ptr[int32](30),
						State:     to.Ptr(armadvisor.DigestConfigStateActive),
						Language:  to.Ptr("<language>"),
					}},
				Exclude:         to.Ptr(true),
				LowCPUThreshold: to.Ptr(armadvisor.CPUThresholdFive),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
