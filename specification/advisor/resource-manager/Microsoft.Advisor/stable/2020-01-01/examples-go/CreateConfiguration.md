Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fadvisor%2Farmadvisor%2Fv0.2.0/sdk/resourcemanager/advisor/armadvisor/README.md) on how to add the SDK to your project and authenticate.

```go
package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
)

// x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/CreateConfiguration.json
func ExampleConfigurationsClient_CreateInResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armadvisor.NewConfigurationsClient("<subscription-id>", cred, nil)
	res, err := client.CreateInResourceGroup(ctx,
		armadvisor.ConfigurationName("default"),
		"<resource-group>",
		armadvisor.ConfigData{
			Properties: &armadvisor.ConfigDataProperties{
				Digests: []*armadvisor.DigestConfig{
					{
						Name:                  to.StringPtr("<name>"),
						ActionGroupResourceID: to.StringPtr("<action-group-resource-id>"),
						Categories: []*armadvisor.Category{
							armadvisor.Category("HighAvailability").ToPtr(),
							armadvisor.Category("Security").ToPtr(),
							armadvisor.Category("Performance").ToPtr(),
							armadvisor.Category("Cost").ToPtr(),
							armadvisor.Category("OperationalExcellence").ToPtr()},
						Frequency: to.Int32Ptr(30),
						State:     armadvisor.DigestConfigState("Active").ToPtr(),
						Language:  to.StringPtr("<language>"),
					}},
				Exclude:         to.BoolPtr(true),
				LowCPUThreshold: armadvisor.CPUThreshold("5").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ConfigurationsClientCreateInResourceGroupResult)
}
```
