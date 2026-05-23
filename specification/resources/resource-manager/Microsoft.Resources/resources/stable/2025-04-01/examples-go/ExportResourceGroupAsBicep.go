package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources/v4"
)

// Generated from example definition: 2025-04-01/ExportResourceGroupAsBicep.json
func ExampleResourceGroupsClient_BeginExportTemplate_exportAResourceGroupAsBicep() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresources.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewResourceGroupsClient().BeginExportTemplate(ctx, "my-resource-group", armresources.ExportTemplateRequest{
		Options:      to.Ptr("IncludeParameterDefaultValue,IncludeComments"),
		OutputFormat: to.Ptr(armresources.ExportTemplateOutputFormatBicep),
		Resources: []*string{
			to.Ptr("*"),
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
	// res = armresources.ResourceGroupsClientExportTemplateResponse{
	// 	ResourceGroupExportResult: armresources.ResourceGroupExportResult{
	// 		Error: &armresources.ErrorResponse{
	// 			Code: to.Ptr("ExportTemplateCompletedWithErrors"),
	// 			Message: to.Ptr("Export template operation completed with errors. Some resources were not exported. Please see details for more information."),
	// 			Details: []*armresources.ErrorResponse{
	// 			},
	// 		},
	// 		Output: to.Ptr("\nparam myResourceType_myFirstResource_name string = 'myFirstResource'\nparam myResourceType_mySecondResource_name string = 'mySecondResource'\n\n@secure()\nparam myResourceType_myFirstResource_secret string = null\n\nresource myResourceType_myFirstResource_name_resource 'My.RP/myResourceType@2019-01-01' = {\n  name: myResourceType_myFirstResource_name\n  location: 'West US'\n  properties: {\n    secret: myResourceType_myFirstResource_secret\n  }\n}\n\nresource myResourceType_mySecondResource_name_resource 'My.RP/myResourceType@2019-01-01' = {\n  name: myResourceType_mySecondResource_name\n  location: 'West US'\n  properties: {\n    customProperty: 'hello!'\n  }\n}\n"),
	// 	},
	// }
}
