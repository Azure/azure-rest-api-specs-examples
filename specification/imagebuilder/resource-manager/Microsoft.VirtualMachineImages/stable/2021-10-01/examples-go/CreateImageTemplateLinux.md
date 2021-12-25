Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvirtualmachineimagebuilder%2Farmvirtualmachineimagebuilder%2Fv0.1.0/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/CreateImageTemplateLinux.json
func ExampleVirtualMachineImageTemplatesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<image-template-name>",
		armvirtualmachineimagebuilder.ImageTemplate{
			TrackedResource: armvirtualmachineimagebuilder.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"imagetemplate_tag1": to.StringPtr("IT_T1"),
					"imagetemplate_tag2": to.StringPtr("IT_T2"),
				},
			},
			Identity: &armvirtualmachineimagebuilder.ImageTemplateIdentity{
				Type: armvirtualmachineimagebuilder.ResourceIdentityTypeUserAssigned.ToPtr(),
				UserAssignedIdentities: map[string]*armvirtualmachineimagebuilder.ComponentsVrq145SchemasImagetemplateidentityPropertiesUserassignedidentitiesAdditionalproperties{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity_1": {},
				},
			},
			Properties: &armvirtualmachineimagebuilder.ImageTemplateProperties{
				Customize: []armvirtualmachineimagebuilder.ImageTemplateCustomizerClassification{
					&armvirtualmachineimagebuilder.ImageTemplateShellCustomizer{
						ImageTemplateCustomizer: armvirtualmachineimagebuilder.ImageTemplateCustomizer{
							Name: to.StringPtr("<name>"),
							Type: to.StringPtr("<type>"),
						},
						ScriptURI: to.StringPtr("<script-uri>"),
					}},
				Distribute: []armvirtualmachineimagebuilder.ImageTemplateDistributorClassification{
					&armvirtualmachineimagebuilder.ImageTemplateManagedImageDistributor{
						ImageTemplateDistributor: armvirtualmachineimagebuilder.ImageTemplateDistributor{
							Type: to.StringPtr("<type>"),
							ArtifactTags: map[string]*string{
								"tagName": to.StringPtr("value"),
							},
							RunOutputName: to.StringPtr("<run-output-name>"),
						},
						ImageID:  to.StringPtr("<image-id>"),
						Location: to.StringPtr("<location>"),
					}},
				Source: &armvirtualmachineimagebuilder.ImageTemplateManagedImageSource{
					ImageTemplateSource: armvirtualmachineimagebuilder.ImageTemplateSource{
						Type: to.StringPtr("<type>"),
					},
					ImageID: to.StringPtr("<image-id>"),
				},
				VMProfile: &armvirtualmachineimagebuilder.ImageTemplateVMProfile{
					OSDiskSizeGB: to.Int32Ptr(64),
					VMSize:       to.StringPtr("<vmsize>"),
					VnetConfig: &armvirtualmachineimagebuilder.VirtualNetworkConfig{
						SubnetID: to.StringPtr("<subnet-id>"),
					},
				},
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
	log.Printf("ImageTemplate.ID: %s\n", *res.ID)
}
```
