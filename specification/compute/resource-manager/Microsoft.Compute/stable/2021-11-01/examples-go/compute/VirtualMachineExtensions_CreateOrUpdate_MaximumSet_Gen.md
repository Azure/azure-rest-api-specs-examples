Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.4.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineExtensions_CreateOrUpdate_MaximumSet_Gen.json
func ExampleVirtualMachineExtensionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineExtensionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vm-name>",
		"<vm-extension-name>",
		armcompute.VirtualMachineExtension{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key9183": to.StringPtr("aa"),
			},
			Properties: &armcompute.VirtualMachineExtensionProperties{
				Type:                    to.StringPtr("<type>"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				EnableAutomaticUpgrade:  to.BoolPtr(true),
				ForceUpdateTag:          to.StringPtr("<force-update-tag>"),
				InstanceView: &armcompute.VirtualMachineExtensionInstanceView{
					Name: to.StringPtr("<name>"),
					Type: to.StringPtr("<type>"),
					Statuses: []*armcompute.InstanceViewStatus{
						{
							Code:          to.StringPtr("<code>"),
							DisplayStatus: to.StringPtr("<display-status>"),
							Level:         armcompute.StatusLevelTypesInfo.ToPtr(),
							Message:       to.StringPtr("<message>"),
							Time:          to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
						}},
					Substatuses: []*armcompute.InstanceViewStatus{
						{
							Code:          to.StringPtr("<code>"),
							DisplayStatus: to.StringPtr("<display-status>"),
							Level:         armcompute.StatusLevelTypesInfo.ToPtr(),
							Message:       to.StringPtr("<message>"),
							Time:          to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
						}},
					TypeHandlerVersion: to.StringPtr("<type-handler-version>"),
				},
				ProtectedSettings:  map[string]interface{}{},
				Publisher:          to.StringPtr("<publisher>"),
				Settings:           map[string]interface{}{},
				SuppressFailures:   to.BoolPtr(true),
				TypeHandlerVersion: to.StringPtr("<type-handler-version>"),
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
	log.Printf("Response result: %#v\n", res.VirtualMachineExtensionsClientCreateOrUpdateResult)
}
```
