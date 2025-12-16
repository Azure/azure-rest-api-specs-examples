package armdeployments_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeployments"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/18609d68cf243ee3ce35d7c005ff3c7dd2cd9477/specification/resources/resource-manager/Microsoft.Resources/deployments/stable/2025-04-01/examples/PostDeploymentWhatIfOnSubscription.json
func ExampleDeploymentsClient_BeginWhatIfAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeployments.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentsClient().BeginWhatIfAtSubscriptionScope(ctx, "my-deployment", armdeployments.DeploymentWhatIf{
		Location: to.Ptr("westus"),
		Properties: &armdeployments.DeploymentWhatIfProperties{
			Mode:       to.Ptr(armdeployments.DeploymentModeIncremental),
			Parameters: map[string]*armdeployments.DeploymentParameter{},
			TemplateLink: &armdeployments.TemplateLink{
				URI: to.Ptr("https://example.com/exampleTemplate.json"),
			},
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
	// res.WhatIfOperationResult = armdeployments.WhatIfOperationResult{
	// 	Properties: &armdeployments.WhatIfOperationProperties{
	// 		Changes: []*armdeployments.WhatIfChange{
	// 			{
	// 				After: map[string]any{
	// 					"name": "myExistingIdentity",
	// 					"type": "Microsoft.ManagedIdentity/userAssignedIdentities",
	// 					"apiVersion": "2018-11-30",
	// 					"id": "/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myExistingIdentity",
	// 					"location": "westus2",
	// 					"tags":map[string]any{
	// 						"myNewTag": "my tag value",
	// 					},
	// 				},
	// 				Before: map[string]any{
	// 					"name": "myExistingIdentity",
	// 					"type": "Microsoft.ManagedIdentity/userAssignedIdentities",
	// 					"apiVersion": "2018-11-30",
	// 					"id": "/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myExistingIdentity",
	// 					"location": "westus2",
	// 				},
	// 				ChangeType: to.Ptr(armdeployments.ChangeTypeModify),
	// 				Delta: []*armdeployments.WhatIfPropertyChange{
	// 					{
	// 						Path: to.Ptr("tags.myNewTag"),
	// 						After: "my tag value",
	// 						PropertyChangeType: to.Ptr(armdeployments.PropertyChangeTypeCreate),
	// 				}},
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myExistingIdentity"),
	// 			},
	// 			{
	// 				After: map[string]any{
	// 					"name": "myNewIdentity",
	// 					"type": "Microsoft.ManagedIdentity/userAssignedIdentities",
	// 					"apiVersion": "2018-11-30",
	// 					"id": "/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myNewIdentity",
	// 					"location": "eastus",
	// 					"tags":map[string]any{
	// 						"myOtherNewTag": "another new tag value",
	// 					},
	// 				},
	// 				ChangeType: to.Ptr(armdeployments.ChangeTypeCreate),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myNewIdentity"),
	// 			},
	// 			{
	// 				After: map[string]any{
	// 					"name": "my-resource-group2",
	// 					"type": "Microsoft.Resources/resourceGroups",
	// 					"apiVersion": "2019-03-01",
	// 					"id": "/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group2",
	// 					"location": "{location3}",
	// 				},
	// 				ChangeType: to.Ptr(armdeployments.ChangeTypeCreate),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group2"),
	// 		}},
	// 	},
	// 	Status: to.Ptr("Succeeded"),
	// }
}
