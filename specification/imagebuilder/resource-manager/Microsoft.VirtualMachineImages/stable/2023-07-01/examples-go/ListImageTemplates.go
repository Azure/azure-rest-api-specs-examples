package armvirtualmachineimagebuilder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2023-07-01/examples/ListImageTemplates.json
func ExampleVirtualMachineImageTemplatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvirtualmachineimagebuilder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineImageTemplatesClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ImageTemplateListResult = armvirtualmachineimagebuilder.ImageTemplateListResult{
		// 	Value: []*armvirtualmachineimagebuilder.ImageTemplate{
		// 		{
		// 			Name: to.Ptr("myImageTemplate"),
		// 			Type: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.VirtualMachineImages/imageTemplates/myImageTemplate"),
		// 			Location: to.Ptr("westus"),
		// 			Identity: &armvirtualmachineimagebuilder.ImageTemplateIdentity{
		// 				Type: to.Ptr(armvirtualmachineimagebuilder.ResourceIdentityTypeUserAssigned),
		// 				UserAssignedIdentities: map[string]*armvirtualmachineimagebuilder.UserAssignedIdentity{
		// 					"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity_1": &armvirtualmachineimagebuilder.UserAssignedIdentity{
		// 						ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armvirtualmachineimagebuilder.ImageTemplateProperties{
		// 				Customize: []armvirtualmachineimagebuilder.ImageTemplateCustomizerClassification{
		// 					&armvirtualmachineimagebuilder.ImageTemplateShellCustomizer{
		// 						Name: to.Ptr("Shell customization example"),
		// 						Type: to.Ptr("Shell"),
		// 						ScriptURI: to.Ptr("https://example.com/path/to/script.sh"),
		// 				}},
		// 				Distribute: []armvirtualmachineimagebuilder.ImageTemplateDistributorClassification{
		// 					&armvirtualmachineimagebuilder.ImageTemplateManagedImageDistributor{
		// 						Type: to.Ptr("ManagedImage"),
		// 						ArtifactTags: map[string]*string{
		// 							"tagName": to.Ptr("value"),
		// 						},
		// 						RunOutputName: to.Ptr("image_it_pir_1"),
		// 						ImageID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Compute/images/image_it_1"),
		// 						Location: to.Ptr("1_location"),
		// 				}},
		// 				Source: &armvirtualmachineimagebuilder.ImageTemplateManagedImageSource{
		// 					Type: to.Ptr("ManagedImage"),
		// 					ImageID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/source_rg/providers/Microsoft.Compute/images/source_image"),
		// 				},
		// 				VMProfile: &armvirtualmachineimagebuilder.ImageTemplateVMProfile{
		// 					OSDiskSizeGB: to.Ptr[int32](64),
		// 					VMSize: to.Ptr("Standard_D2s_v3"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mySecondImageTemplate"),
		// 			Type: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myOtherResourceGroup/providers/Microsoft.VirtualMachineImages/imageTemplates/mySecondImageTemplate"),
		// 			Location: to.Ptr("westus"),
		// 			Identity: &armvirtualmachineimagebuilder.ImageTemplateIdentity{
		// 				Type: to.Ptr(armvirtualmachineimagebuilder.ResourceIdentityTypeUserAssigned),
		// 				UserAssignedIdentities: map[string]*armvirtualmachineimagebuilder.UserAssignedIdentity{
		// 					"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity_1": &armvirtualmachineimagebuilder.UserAssignedIdentity{
		// 						ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armvirtualmachineimagebuilder.ImageTemplateProperties{
		// 				Customize: []armvirtualmachineimagebuilder.ImageTemplateCustomizerClassification{
		// 					&armvirtualmachineimagebuilder.ImageTemplateShellCustomizer{
		// 						Name: to.Ptr("Shell customization example"),
		// 						Type: to.Ptr("Shell"),
		// 						ScriptURI: to.Ptr("https://example.com/path/to/script.sh"),
		// 				}},
		// 				Distribute: []armvirtualmachineimagebuilder.ImageTemplateDistributorClassification{
		// 					&armvirtualmachineimagebuilder.ImageTemplateManagedImageDistributor{
		// 						Type: to.Ptr("ManagedImage"),
		// 						ArtifactTags: map[string]*string{
		// 							"stage": to.Ptr("development"),
		// 						},
		// 						RunOutputName: to.Ptr("eus"),
		// 						ImageID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Compute/images/mySecondImage-eus"),
		// 						Location: to.Ptr("eastus"),
		// 				}},
		// 				Source: &armvirtualmachineimagebuilder.ImageTemplatePlatformImageSource{
		// 					Type: to.Ptr("PlatformImage"),
		// 					Offer: to.Ptr("UbuntuServer"),
		// 					PlanInfo: &armvirtualmachineimagebuilder.PlatformImagePurchasePlan{
		// 						PlanName: to.Ptr("example_plan_name"),
		// 						PlanProduct: to.Ptr("example_plan_product"),
		// 						PlanPublisher: to.Ptr("example_plan_publisher"),
		// 					},
		// 					Publisher: to.Ptr("Canonical"),
		// 					SKU: to.Ptr("18.04-LTS"),
		// 					Version: to.Ptr("18.04.201902121"),
		// 				},
		// 				VMProfile: &armvirtualmachineimagebuilder.ImageTemplateVMProfile{
		// 					OSDiskSizeGB: to.Ptr[int32](32),
		// 					VMSize: to.Ptr("Standard_D8s_v3"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
