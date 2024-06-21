package armdeploymentstacks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentstacks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/DeploymentStackResourceGroupValidate.json
func ExampleClient_BeginValidateStackAtResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentstacks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginValidateStackAtResourceGroup(ctx, "deploymentStacksRG", "simpleDeploymentStack", armdeploymentstacks.DeploymentStack{
		Properties: &armdeploymentstacks.DeploymentStackProperties{
			ActionOnUnmanage: &armdeploymentstacks.ActionOnUnmanage{
				ManagementGroups: to.Ptr(armdeploymentstacks.DeploymentStacksDeleteDetachEnumDelete),
				ResourceGroups:   to.Ptr(armdeploymentstacks.DeploymentStacksDeleteDetachEnumDelete),
				Resources:        to.Ptr(armdeploymentstacks.DeploymentStacksDeleteDetachEnumDelete),
			},
			DenySettings: &armdeploymentstacks.DenySettings{
				ApplyToChildScopes: to.Ptr(false),
				ExcludedActions: []*string{
					to.Ptr("action")},
				ExcludedPrincipals: []*string{
					to.Ptr("principal")},
				Mode: to.Ptr(armdeploymentstacks.DenySettingsModeDenyDelete),
			},
			Parameters: map[string]*armdeploymentstacks.DeploymentParameter{
				"parameter1": {
					Value: "a string",
				},
			},
			TemplateLink: &armdeploymentstacks.TemplateLink{
				URI: to.Ptr("https://example.com/exampleTemplate.json"),
			},
		},
		Tags: map[string]*string{
			"tagkey": to.Ptr("tagVal"),
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
	// res.DeploymentStackValidateResult = armdeploymentstacks.DeploymentStackValidateResult{
	// 	Name: to.Ptr("simpleDeploymentStack"),
	// 	Type: to.Ptr("Microsoft.Resources/deploymentStacks"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/deploymentStacksRG/providers/Microsoft.Resources/deploymentStacks/simpleDeploymentStack"),
	// 	Properties: &armdeploymentstacks.DeploymentStackValidateProperties{
	// 		Description: to.Ptr("A validation description."),
	// 		CorrelationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		DenySettings: &armdeploymentstacks.DenySettings{
	// 			ApplyToChildScopes: to.Ptr(false),
	// 			ExcludedActions: []*string{
	// 				to.Ptr("action")},
	// 				ExcludedPrincipals: []*string{
	// 					to.Ptr("principal")},
	// 					Mode: to.Ptr(armdeploymentstacks.DenySettingsModeDenyDelete),
	// 				},
	// 				DeploymentScope: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/deploymentStacksRG"),
	// 				Parameters: map[string]*armdeploymentstacks.DeploymentParameter{
	// 					"parameter1": &armdeploymentstacks.DeploymentParameter{
	// 						Type: to.Ptr("string"),
	// 						Value: "a string",
	// 					},
	// 				},
	// 				ValidatedResources: []*armdeploymentstacks.ResourceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/deploymentStacksRG/providers/Microsoft.Sql/servers/server1"),
	// 					},
	// 					{
	// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/deploymentStacksRG/providers/Microsoft.Sql/servers/server2"),
	// 				}},
	// 			},
	// 		}
}
