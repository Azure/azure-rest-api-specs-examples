Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvirtualmachineimagebuilder%2Farmvirtualmachineimagebuilder%2Fv0.4.0/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder/README.md) on how to add the SDK to your project and authenticate.

```go
package armvirtualmachineimagebuilder_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/CreateImageTemplateLinux.json
func ExampleVirtualMachineImageTemplatesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<image-template-name>",
		armvirtualmachineimagebuilder.ImageTemplate{
			Location: to.Ptr("<location>"),
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
						Name:      to.Ptr("<name>"),
						Type:      to.Ptr("<type>"),
						ScriptURI: to.Ptr("<script-uri>"),
					}},
				Distribute: []armvirtualmachineimagebuilder.ImageTemplateDistributorClassification{
					&armvirtualmachineimagebuilder.ImageTemplateManagedImageDistributor{
						Type: to.Ptr("<type>"),
						ArtifactTags: map[string]*string{
							"tagName": to.Ptr("value"),
						},
						RunOutputName: to.Ptr("<run-output-name>"),
						ImageID:       to.Ptr("<image-id>"),
						Location:      to.Ptr("<location>"),
					}},
				Source: &armvirtualmachineimagebuilder.ImageTemplateManagedImageSource{
					Type:    to.Ptr("<type>"),
					ImageID: to.Ptr("<image-id>"),
				},
				VMProfile: &armvirtualmachineimagebuilder.ImageTemplateVMProfile{
					OSDiskSizeGB: to.Ptr[int32](64),
					VMSize:       to.Ptr("<vmsize>"),
					VnetConfig: &armvirtualmachineimagebuilder.VirtualNetworkConfig{
						SubnetID: to.Ptr("<subnet-id>"),
					},
				},
			},
		},
		&armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
