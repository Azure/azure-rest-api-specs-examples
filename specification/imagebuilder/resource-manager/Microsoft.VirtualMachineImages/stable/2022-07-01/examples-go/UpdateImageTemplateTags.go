package armvirtualmachineimagebuilder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-07-01/examples/UpdateImageTemplateTags.json
func ExampleVirtualMachineImageTemplatesClient_BeginUpdate_updateTheTagsForAnImageTemplate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvirtualmachineimagebuilder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineImageTemplatesClient().BeginUpdate(ctx, "myResourceGroup", "myImageTemplate", armvirtualmachineimagebuilder.ImageTemplateUpdateParameters{
		Tags: map[string]*string{
			"new-tag": to.Ptr("new-value"),
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
	// res.ImageTemplate = armvirtualmachineimagebuilder.ImageTemplate{
	// 	Name: to.Ptr("myImageTemplate"),
	// 	Type: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.VirtualMachineImages/imageTemplates/myImageTemplate"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"new-tag": to.Ptr("new-value"),
	// 	},
	// 	Identity: &armvirtualmachineimagebuilder.ImageTemplateIdentity{
	// 		Type: to.Ptr(armvirtualmachineimagebuilder.ResourceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armvirtualmachineimagebuilder.UserAssignedIdentity{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity_1": &armvirtualmachineimagebuilder.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armvirtualmachineimagebuilder.ImageTemplateProperties{
	// 		Customize: []armvirtualmachineimagebuilder.ImageTemplateCustomizerClassification{
	// 			&armvirtualmachineimagebuilder.ImageTemplateShellCustomizer{
	// 				Name: to.Ptr("Shell customization example"),
	// 				Type: to.Ptr("Shell"),
	// 				ScriptURI: to.Ptr("https://example.com/path/to/script.sh"),
	// 		}},
	// 		Distribute: []armvirtualmachineimagebuilder.ImageTemplateDistributorClassification{
	// 			&armvirtualmachineimagebuilder.ImageTemplateManagedImageDistributor{
	// 				Type: to.Ptr("ManagedImage"),
	// 				ArtifactTags: map[string]*string{
	// 					"tagName": to.Ptr("value"),
	// 				},
	// 				RunOutputName: to.Ptr("image_it_pir_1"),
	// 				ImageID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Compute/images/image_it_1"),
	// 				Location: to.Ptr("1_location"),
	// 		}},
	// 		Source: &armvirtualmachineimagebuilder.ImageTemplateManagedImageSource{
	// 			Type: to.Ptr("ManagedImage"),
	// 			ImageID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/source_rg/providers/Microsoft.Compute/images/source_image"),
	// 		},
	// 		VMProfile: &armvirtualmachineimagebuilder.ImageTemplateVMProfile{
	// 			OSDiskSizeGB: to.Ptr[int32](64),
	// 			VMSize: to.Ptr("Standard_D2s_v3"),
	// 		},
	// 	},
	// }
}
