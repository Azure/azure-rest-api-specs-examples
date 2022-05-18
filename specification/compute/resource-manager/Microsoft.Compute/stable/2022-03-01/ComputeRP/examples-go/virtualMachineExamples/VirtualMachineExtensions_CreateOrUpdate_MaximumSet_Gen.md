Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv1.0.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachineExtensions_CreateOrUpdate_MaximumSet_Gen.json
func ExampleVirtualMachineExtensionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineExtensionsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rgcompute",
		"aaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaaaaaa",
		armcompute.VirtualMachineExtension{
			Location: to.Ptr("westus"),
			Tags: map[string]*string{
				"key9183": to.Ptr("aa"),
			},
			Properties: &armcompute.VirtualMachineExtensionProperties{
				Type:                    to.Ptr("extType"),
				AutoUpgradeMinorVersion: to.Ptr(true),
				EnableAutomaticUpgrade:  to.Ptr(true),
				ForceUpdateTag:          to.Ptr("a"),
				InstanceView: &armcompute.VirtualMachineExtensionInstanceView{
					Name: to.Ptr("aaaaaaaaaaaaaaaaa"),
					Type: to.Ptr("aaaaaaaaa"),
					Statuses: []*armcompute.InstanceViewStatus{
						{
							Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
							DisplayStatus: to.Ptr("aaaaaa"),
							Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
							Message:       to.Ptr("a"),
							Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
						}},
					Substatuses: []*armcompute.InstanceViewStatus{
						{
							Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
							DisplayStatus: to.Ptr("aaaaaa"),
							Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
							Message:       to.Ptr("a"),
							Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
						}},
					TypeHandlerVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
				},
				ProtectedSettings:  map[string]interface{}{},
				Publisher:          to.Ptr("extPublisher"),
				Settings:           map[string]interface{}{},
				SuppressFailures:   to.Ptr(true),
				TypeHandlerVersion: to.Ptr("1.2"),
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
