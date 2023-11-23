package armdeploymentstacks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentstacks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ca162308f1010bfb85b9c85021e863e7bd397a1f/specification/resources/resource-manager/Microsoft.Resources/preview/2022-08-01-preview/examples/DeploymentStackSubscriptionCreate.json
func ExampleClient_BeginCreateOrUpdateAtSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentstacks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginCreateOrUpdateAtSubscription(ctx, "simpleDeploymentStack", armdeploymentstacks.DeploymentStack{
		Location: to.Ptr("eastus"),
		Properties: &armdeploymentstacks.DeploymentStackProperties{
			ActionOnUnmanage: &armdeploymentstacks.DeploymentStackPropertiesActionOnUnmanage{
				ManagementGroups: to.Ptr(armdeploymentstacks.DeploymentStacksDeleteDetachEnumDetach),
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
			Parameters: map[string]any{
				"parameter1": map[string]any{
					"value": "a string",
				},
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
	// res.DeploymentStack = armdeploymentstacks.DeploymentStack{
	// 	Name: to.Ptr("simpleDeploymentStack"),
	// 	Type: to.Ptr("Microsoft.Resources/deploymentStacks"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/deploymentStacksRG/providers/Microsoft.Resources/deploymentStacks/simpleDeploymentStack"),
	// 	SystemData: &armdeploymentstacks.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armdeploymentstacks.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-02T02:03:01.197Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armdeploymentstacks.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armdeploymentstacks.DeploymentStackProperties{
	// 		Description: to.Ptr("my Description"),
	// 		ActionOnUnmanage: &armdeploymentstacks.DeploymentStackPropertiesActionOnUnmanage{
	// 			ManagementGroups: to.Ptr(armdeploymentstacks.DeploymentStacksDeleteDetachEnumDetach),
	// 			ResourceGroups: to.Ptr(armdeploymentstacks.DeploymentStacksDeleteDetachEnumDelete),
	// 			Resources: to.Ptr(armdeploymentstacks.DeploymentStacksDeleteDetachEnumDelete),
	// 		},
	// 		DenySettings: &armdeploymentstacks.DenySettings{
	// 			ApplyToChildScopes: to.Ptr(false),
	// 			ExcludedActions: []*string{
	// 				to.Ptr("action")},
	// 				ExcludedPrincipals: []*string{
	// 					to.Ptr("principal")},
	// 					Mode: to.Ptr(armdeploymentstacks.DenySettingsModeDenyDelete),
	// 				},
	// 				Parameters: map[string]any{
	// 					"parameter1":map[string]any{
	// 						"value": "a string",
	// 					},
	// 				},
	// 				ProvisioningState: to.Ptr(armdeploymentstacks.DeploymentStackProvisioningStateSucceeded),
	// 			},
	// 			Tags: map[string]*string{
	// 				"tagkey": to.Ptr("tagVal"),
	// 			},
	// 		}
}
