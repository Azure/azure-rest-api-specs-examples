package armvirtualmachineimagebuilder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/UpdateImageTemplateToRemoveIdentities.json
func ExampleVirtualMachineImageTemplatesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myImageTemplate",
		armvirtualmachineimagebuilder.ImageTemplateUpdateParameters{
			Identity: &armvirtualmachineimagebuilder.ImageTemplateIdentity{
				Type: to.Ptr(armvirtualmachineimagebuilder.ResourceIdentityTypeNone),
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
