package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fd842fb73656039ec94ce367bcedee25a57bd18/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PostDeploymentWhatIfOnTenant.json
func ExampleDeploymentsClient_BeginWhatIfAtTenantScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresources.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentsClient().BeginWhatIfAtTenantScope(ctx, "exampleDeploymentName", armresources.ScopedDeploymentWhatIf{
		Location: to.Ptr("eastus"),
		Properties: &armresources.DeploymentWhatIfProperties{
			Mode:         to.Ptr(armresources.DeploymentModeIncremental),
			Parameters:   map[string]any{},
			TemplateLink: &armresources.TemplateLink{},
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
	// res.WhatIfOperationResult = armresources.WhatIfOperationResult{
	// 	Properties: &armresources.WhatIfOperationProperties{
	// 		Changes: []*armresources.WhatIfChange{
	// 			{
	// 				After: map[string]any{
	// 					"name": "myManagementGroup",
	// 					"type": "Microsoft.Management/managementGroups",
	// 					"apiVersion": "2019-11-01",
	// 					"id": "/providers/Microsoft.Management/managementGroups/myManagementGroup",
	// 				},
	// 				ChangeType: to.Ptr(armresources.ChangeTypeCreate),
	// 				ResourceID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup"),
	// 		}},
	// 	},
	// 	Status: to.Ptr("Succeeded"),
	// }
}
