package armdeploymentstacks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentstacks/v2"
)

// Generated from example definition: 2025-07-01/DeploymentStackManagementGroupCreate.json
func ExampleClient_BeginCreateOrUpdateAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentstacks.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginCreateOrUpdateAtManagementGroup(ctx, "myMg", "simpleDeploymentStack", armdeploymentstacks.DeploymentStack{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"tagkey": to.Ptr("tagVal"),
		},
		Properties: &armdeploymentstacks.DeploymentStackProperties{
			ActionOnUnmanage: &armdeploymentstacks.ActionOnUnmanage{
				Resources:        to.Ptr(armdeploymentstacks.UnmanageActionResourceModeDelete),
				ResourceGroups:   to.Ptr(armdeploymentstacks.UnmanageActionResourceGroupModeDelete),
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
			Parameters: map[string]*armdeploymentstacks.DeploymentParameter{
				"parameter1": {
					Value: "a string",
				},
			},
			ExtensionConfigs: map[string]*armdeploymentstacks.DeploymentExtensionConfig{
				"contoso": {
					AdditionalProperties: map[string]*armdeploymentstacks.DeploymentExtensionConfigItem{
						"configOne": {
							Value: "config1Value",
						},
						"configTwo": {
							Value: true,
						},
					},
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
	// res = armdeploymentstacks.ClientCreateOrUpdateAtManagementGroupResponse{
	// 	DeploymentStack: &armdeploymentstacks.DeploymentStack{
	// 		ID: to.Ptr("/providers/Microsoft.Management/managementGroups/myMg/providers/Microsoft.Resources/deploymentStacks/simpleDeploymentStack"),
	// 		Type: to.Ptr("Microsoft.Resources/deploymentStacks"),
	// 		Name: to.Ptr("simpleDeploymentStack"),
	// 		Location: to.Ptr("eastus"),
	// 		Tags: map[string]*string{
	// 			"tagkey": to.Ptr("tagVal"),
	// 		},
	// 		SystemData: &armdeploymentstacks.SystemData{
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armdeploymentstacks.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armdeploymentstacks.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-02T02:03:01.1974346Z"); return t}()),
	// 		},
	// 		Properties: &armdeploymentstacks.DeploymentStackProperties{
	// 			Description: to.Ptr("my Description"),
	// 			ActionOnUnmanage: &armdeploymentstacks.ActionOnUnmanage{
	// 				Resources: to.Ptr(armdeploymentstacks.UnmanageActionResourceModeDelete),
	// 				ResourceGroups: to.Ptr(armdeploymentstacks.UnmanageActionResourceGroupModeDelete),
	// 				ManagementGroups: to.Ptr(armdeploymentstacks.UnmanageActionManagementGroupModeDetach),
	// 			},
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
	// 			Parameters: map[string]*armdeploymentstacks.DeploymentParameter{
	// 				"parameter1": &armdeploymentstacks.DeploymentParameter{
	// 					Value: "a string",
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armdeploymentstacks.DeploymentStackProvisioningStateCreating),
	// 		},
	// 	},
	// }
}
