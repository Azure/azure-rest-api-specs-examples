package armdeploymentstacks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentstacks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edacc3b43f9603efa119eabb6013d952d1dbe7d6/specification/resources/resource-manager/Microsoft.Resources/deploymentStacks/stable/2024-03-01/examples/DeploymentStackManagementGroupValidate.json
func ExampleClient_BeginValidateStackAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentstacks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginValidateStackAtManagementGroup(ctx, "myMg", "simpleDeploymentStack", armdeploymentstacks.DeploymentStack{
		Location: to.Ptr("eastus"),
		Properties: &armdeploymentstacks.DeploymentStackProperties{
			ActionOnUnmanage: &armdeploymentstacks.ActionOnUnmanage{
				ManagementGroups: to.Ptr(armdeploymentstacks.DeploymentStacksDeleteDetachEnumDetach),
				ResourceGroups:   to.Ptr(armdeploymentstacks.DeploymentStacksDeleteDetachEnumDetach),
				Resources:        to.Ptr(armdeploymentstacks.DeploymentStacksDeleteDetachEnumDetach),
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
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/myMg/providers/Microsoft.Resources/deploymentStacks/simpleDeploymentStack"),
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
	// 				DeploymentScope: to.Ptr("/providers/Microsoft.Management/managementGroups/myMg"),
	// 				Parameters: map[string]*armdeploymentstacks.DeploymentParameter{
	// 					"parameter1": &armdeploymentstacks.DeploymentParameter{
	// 						Type: to.Ptr("string"),
	// 						Value: "a string",
	// 					},
	// 				},
	// 				ValidatedResources: []*armdeploymentstacks.ResourceReference{
	// 					{
	// 						ID: to.Ptr("/providers/Microsoft.Management/managementGroups/myMg/providers/Microsoft.Authorization/policyDefinitions/Policy1"),
	// 				}},
	// 			},
	// 		}
}
