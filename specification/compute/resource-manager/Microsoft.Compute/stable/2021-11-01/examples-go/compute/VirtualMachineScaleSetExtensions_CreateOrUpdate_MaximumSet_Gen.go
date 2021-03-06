package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetExtensions_CreateOrUpdate_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetExtensionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetExtensionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<vmss-extension-name>",
		armcompute.VirtualMachineScaleSetExtension{
			Name: to.Ptr("<name>"),
			Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
				Type:                    to.Ptr("<type>"),
				AutoUpgradeMinorVersion: to.Ptr(true),
				EnableAutomaticUpgrade:  to.Ptr(true),
				ForceUpdateTag:          to.Ptr("<force-update-tag>"),
				ProtectedSettings:       map[string]interface{}{},
				ProvisionAfterExtensions: []*string{
					to.Ptr("aa")},
				Publisher:          to.Ptr("<publisher>"),
				Settings:           map[string]interface{}{},
				SuppressFailures:   to.Ptr(true),
				TypeHandlerVersion: to.Ptr("<type-handler-version>"),
			},
		},
		&armcompute.VirtualMachineScaleSetExtensionsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
