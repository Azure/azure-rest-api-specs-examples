package armdeploymentstacks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentstacks/v2"
)

// Generated from example definition: 2025-07-01/DeploymentStackManagementGroupValidate.json
func ExampleClient_BeginValidateStackAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentstacks.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginValidateStackAtManagementGroup(ctx, "myMg", "simpleDeploymentStack", armdeploymentstacks.DeploymentStack{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"tagkey": to.Ptr("tagVal"),
		},
		Properties: &armdeploymentstacks.DeploymentStackProperties{
			ActionOnUnmanage: &armdeploymentstacks.ActionOnUnmanage{
				Resources:        to.Ptr(armdeploymentstacks.UnmanageActionResourceModeDetach),
				ResourceGroups:   to.Ptr(armdeploymentstacks.UnmanageActionResourceGroupModeDetach),
				ManagementGroups: to.Ptr(armdeploymentstacks.UnmanageActionManagementGroupModeDetach),
			},
			DenySettings: &armdeploymentstacks.DenySettings{
				Mode: to.Ptr(armdeploymentstacks.DenySettingsModeDenyDelete),
				ExcludedPrincipals: []*string{
					to.Ptr("principal"),
				},
				ExcludedActions: []*string{
					to.Ptr("action"),
				},
				ApplyToChildScopes: to.Ptr(false),
			},
			TemplateLink: &armdeploymentstacks.TemplateLink{
				URI: to.Ptr("https://example.com/exampleTemplate.json"),
			},
			Parameters: map[string]*armdeploymentstacks.DeploymentParameter{
				"parameter1": {
					Value: "a string",
				},
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
	// res = armdeploymentstacks.ClientValidateStackAtManagementGroupResponse{
	// 	DeploymentStackValidateResult: &armdeploymentstacks.DeploymentStackValidateResult{
	// 		ID: to.Ptr("/providers/Microsoft.Management/managementGroups/myMg/providers/Microsoft.Resources/deploymentStacks/simpleDeploymentStack"),
	// 		Name: to.Ptr("simpleDeploymentStack"),
	// 		Type: to.Ptr("Microsoft.Resources/deploymentStacks"),
	// 		Properties: &armdeploymentstacks.DeploymentStackValidateProperties{
	// 			CorrelationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			DenySettings: &armdeploymentstacks.DenySettings{
	// 				Mode: to.Ptr(armdeploymentstacks.DenySettingsModeDenyDelete),
	// 				ExcludedPrincipals: []*string{
	// 					to.Ptr("principal"),
	// 				},
	// 				ExcludedActions: []*string{
	// 					to.Ptr("action"),
	// 				},
	// 				ApplyToChildScopes: to.Ptr(false),
	// 			},
	// 			DeploymentScope: to.Ptr("/providers/Microsoft.Management/managementGroups/myMg"),
	// 			Description: to.Ptr("A validation description."),
	// 			Parameters: map[string]*armdeploymentstacks.DeploymentParameter{
	// 				"parameter1": &armdeploymentstacks.DeploymentParameter{
	// 					Type: to.Ptr("string"),
	// 					Value: "a string",
	// 				},
	// 			},
	// 			ValidatedResources: []*armdeploymentstacks.ResourceReference{
	// 				{
	// 					ID: to.Ptr("/providers/Microsoft.Management/managementGroups/myMg/providers/Microsoft.Authorization/policyDefinitions/Policy1"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
