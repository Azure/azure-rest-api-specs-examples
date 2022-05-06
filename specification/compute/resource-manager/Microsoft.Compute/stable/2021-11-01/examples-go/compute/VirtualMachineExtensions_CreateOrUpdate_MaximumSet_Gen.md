Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.7.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineExtensions_CreateOrUpdate_MaximumSet_Gen.json
func ExampleVirtualMachineExtensionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineExtensionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vm-name>",
		"<vm-extension-name>",
		armcompute.VirtualMachineExtension{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key9183": to.Ptr("aa"),
			},
			Properties: &armcompute.VirtualMachineExtensionProperties{
				Type:                    to.Ptr("<type>"),
				AutoUpgradeMinorVersion: to.Ptr(true),
				EnableAutomaticUpgrade:  to.Ptr(true),
				ForceUpdateTag:          to.Ptr("<force-update-tag>"),
				InstanceView: &armcompute.VirtualMachineExtensionInstanceView{
					Name: to.Ptr("<name>"),
					Type: to.Ptr("<type>"),
					Statuses: []*armcompute.InstanceViewStatus{
						{
							Code:          to.Ptr("<code>"),
							DisplayStatus: to.Ptr("<display-status>"),
							Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
							Message:       to.Ptr("<message>"),
							Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
						}},
					Substatuses: []*armcompute.InstanceViewStatus{
						{
							Code:          to.Ptr("<code>"),
							DisplayStatus: to.Ptr("<display-status>"),
							Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
							Message:       to.Ptr("<message>"),
							Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
						}},
					TypeHandlerVersion: to.Ptr("<type-handler-version>"),
				},
				ProtectedSettings:  map[string]interface{}{},
				Publisher:          to.Ptr("<publisher>"),
				Settings:           map[string]interface{}{},
				SuppressFailures:   to.Ptr(true),
				TypeHandlerVersion: to.Ptr("<type-handler-version>"),
			},
		},
		&armcompute.VirtualMachineExtensionsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
