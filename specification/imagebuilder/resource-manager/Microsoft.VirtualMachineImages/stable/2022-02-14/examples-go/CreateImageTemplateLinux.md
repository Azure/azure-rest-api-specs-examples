```go
package armvirtualmachineimagebuilder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/CreateImageTemplateLinux.json
func ExampleVirtualMachineImageTemplatesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myImageTemplate",
		armvirtualmachineimagebuilder.ImageTemplate{
			Location: to.Ptr("westus"),
			Tags: map[string]*string{
				"imagetemplate_tag1": to.Ptr("IT_T1"),
				"imagetemplate_tag2": to.Ptr("IT_T2"),
			},
			Identity: &armvirtualmachineimagebuilder.ImageTemplateIdentity{
				Type: to.Ptr(armvirtualmachineimagebuilder.ResourceIdentityTypeUserAssigned),
				UserAssignedIdentities: map[string]*armvirtualmachineimagebuilder.ComponentsVrq145SchemasImagetemplateidentityPropertiesUserassignedidentitiesAdditionalproperties{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity_1": {},
				},
			},
			Properties: &armvirtualmachineimagebuilder.ImageTemplateProperties{
				Customize: []armvirtualmachineimagebuilder.ImageTemplateCustomizerClassification{
					&armvirtualmachineimagebuilder.ImageTemplateShellCustomizer{
						Name:      to.Ptr("Shell Customizer Example"),
						Type:      to.Ptr("Shell"),
						ScriptURI: to.Ptr("https://example.com/path/to/script.sh"),
					}},
				Distribute: []armvirtualmachineimagebuilder.ImageTemplateDistributorClassification{
					&armvirtualmachineimagebuilder.ImageTemplateManagedImageDistributor{
						Type: to.Ptr("ManagedImage"),
						ArtifactTags: map[string]*string{
							"tagName": to.Ptr("value"),
						},
						RunOutputName: to.Ptr("image_it_pir_1"),
						ImageID:       to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Compute/images/image_it_1"),
						Location:      to.Ptr("1_location"),
					}},
				Source: &armvirtualmachineimagebuilder.ImageTemplateManagedImageSource{
					Type:    to.Ptr("ManagedImage"),
					ImageID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/images/source_image"),
				},
				VMProfile: &armvirtualmachineimagebuilder.ImageTemplateVMProfile{
					OSDiskSizeGB: to.Ptr[int32](64),
					VMSize:       to.Ptr("Standard_D2s_v3"),
					VnetConfig: &armvirtualmachineimagebuilder.VirtualNetworkConfig{
						SubnetID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet_name/subnets/subnet_name"),
					},
				},
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvirtualmachineimagebuilder%2Farmvirtualmachineimagebuilder%2Fv1.1.0/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder/README.md) on how to add the SDK to your project and authenticate.
