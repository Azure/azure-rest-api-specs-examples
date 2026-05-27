package armvirtualmachineimagebuilder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder/v3"
)

// Generated from example definition: 2025-10-01/UpdateImageTemplateVmProfile.json
func ExampleVirtualMachineImageTemplatesClient_BeginUpdate_updateParametersForVMProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvirtualmachineimagebuilder.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineImageTemplatesClient().BeginUpdate(ctx, "myResourceGroup", "myImageTemplate", armvirtualmachineimagebuilder.ImageTemplateUpdateParameters{
		Properties: &armvirtualmachineimagebuilder.ImageTemplateUpdateParametersProperties{
			VMProfile: &armvirtualmachineimagebuilder.ImageTemplateVMProfile{
				OSDiskSizeGB: to.Ptr[int32](127),
				VMSize:       to.Ptr("{updated_vmsize}"),
				VnetConfig: &armvirtualmachineimagebuilder.VirtualNetworkConfig{
					ContainerInstanceSubnetID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/subnetname"),
					SubnetID:                  to.Ptr("{updated_aci_subnet}"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientUpdateResponse{
	// 	ImageTemplate: armvirtualmachineimagebuilder.ImageTemplate{
	// 		Name: to.Ptr("myImageTemplate"),
	// 		Type: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.VirtualMachineImages/imageTemplates/myImageTemplate"),
	// 		Identity: &armvirtualmachineimagebuilder.ImageTemplateIdentity{
	// 			Type: to.Ptr(armvirtualmachineimagebuilder.ResourceIdentityTypeNone),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armvirtualmachineimagebuilder.ImageTemplateProperties{
	// 			Customize: []armvirtualmachineimagebuilder.ImageTemplateCustomizerClassification{
	// 				&armvirtualmachineimagebuilder.ImageTemplateShellCustomizer{
	// 					Name: to.Ptr("Shell customization example"),
	// 					Type: to.Ptr("Shell"),
	// 					ScriptURI: to.Ptr("https://example.com/path/to/script.sh"),
	// 				},
	// 			},
	// 			Distribute: []armvirtualmachineimagebuilder.ImageTemplateDistributorClassification{
	// 				&armvirtualmachineimagebuilder.ImageTemplateManagedImageDistributor{
	// 					Type: to.Ptr("ManagedImage"),
	// 					ArtifactTags: map[string]*string{
	// 						"tagName": to.Ptr("value"),
	// 					},
	// 					ImageID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Compute/images/image_it_1"),
	// 					Location: to.Ptr("1_location"),
	// 					RunOutputName: to.Ptr("image_it_pir_1"),
	// 				},
	// 			},
	// 			Source: &armvirtualmachineimagebuilder.ImageTemplateManagedImageSource{
	// 				Type: to.Ptr("ManagedImage"),
	// 				ImageID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/source_rg/providers/Microsoft.Compute/images/source_image"),
	// 			},
	// 			VMProfile: &armvirtualmachineimagebuilder.ImageTemplateVMProfile{
	// 				OSDiskSizeGB: to.Ptr[int32](127),
	// 				VMSize: to.Ptr("{updated_vmsize}"),
	// 				VnetConfig: &armvirtualmachineimagebuilder.VirtualNetworkConfig{
	// 					ContainerInstanceSubnetID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/subnetname"),
	// 					SubnetID: to.Ptr("{updated_aci_subnet}"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag": to.Ptr("tag-value"),
	// 		},
	// 	},
	// }
}
