package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_CreateOrUpdate_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetExtensionsClient_BeginCreateOrUpdate_virtualMachineScaleSetExtensionCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetExtensionsClient().BeginCreateOrUpdate(ctx, "rgcompute", "aaaaaaa", "aaaaaaaaaaaaaaaaaaaaa", armcompute.VirtualMachineScaleSetExtension{
		Name: to.Ptr("{extension-name}"),
		Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
			Type:                    to.Ptr("{extension-Type}"),
			AutoUpgradeMinorVersion: to.Ptr(true),
			EnableAutomaticUpgrade:  to.Ptr(true),
			ForceUpdateTag:          to.Ptr("aaaaaaaaa"),
			ProtectedSettings:       map[string]any{},
			ProvisionAfterExtensions: []*string{
				to.Ptr("aa")},
			Publisher:          to.Ptr("{extension-Publisher}"),
			Settings:           map[string]any{},
			SuppressFailures:   to.Ptr(true),
			TypeHandlerVersion: to.Ptr("{handler-version}"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineScaleSetExtension = armcompute.VirtualMachineScaleSetExtension{
	// 	ID: to.Ptr("aaaaaaaa"),
	// 	Name: to.Ptr("{extension-name}"),
	// 	Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
	// 		Type: to.Ptr("{extension-Type}"),
	// 		AutoUpgradeMinorVersion: to.Ptr(true),
	// 		EnableAutomaticUpgrade: to.Ptr(true),
	// 		ForceUpdateTag: to.Ptr("aaaaaaaaa"),
	// 		ProtectedSettings: map[string]any{
	// 		},
	// 		ProvisionAfterExtensions: []*string{
	// 			to.Ptr("aa")},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			Publisher: to.Ptr("{extension-Publisher}"),
	// 			Settings: map[string]any{
	// 			},
	// 			SuppressFailures: to.Ptr(true),
	// 			TypeHandlerVersion: to.Ptr("{handler-version}"),
	// 		},
	// 	}
}
