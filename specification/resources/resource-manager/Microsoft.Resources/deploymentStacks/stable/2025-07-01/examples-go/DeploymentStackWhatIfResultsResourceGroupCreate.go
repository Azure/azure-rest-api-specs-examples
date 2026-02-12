package armdeploymentstacks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentstacks/v2"
)

// Generated from example definition: 2025-07-01/DeploymentStackWhatIfResultsResourceGroupCreate.json
func ExampleWhatIfResultsAtResourceGroupClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentstacks.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWhatIfResultsAtResourceGroupClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "simpleDeploymentStackWhatIfResult", armdeploymentstacks.WhatIfResult{
		Location: to.Ptr("eastus"),
		Properties: &armdeploymentstacks.WhatIfResultProperties{
			DeploymentStackResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Resources/deploymentStacks/simpleDeploymentStack"),
			RetentionInterval:         to.Ptr("P7D"),
			TemplateLink: &armdeploymentstacks.TemplateLink{
				URI: to.Ptr("https://example.com/exampleTemplate.json"),
			},
			Parameters: map[string]*armdeploymentstacks.DeploymentParameter{},
			ActionOnUnmanage: &armdeploymentstacks.ActionOnUnmanage{
				Resources:        to.Ptr(armdeploymentstacks.UnmanageActionResourceModeDelete),
				ResourceGroups:   to.Ptr(armdeploymentstacks.UnmanageActionResourceGroupModeDelete),
				ManagementGroups: to.Ptr(armdeploymentstacks.UnmanageActionManagementGroupModeDetach),
			},
			DenySettings: &armdeploymentstacks.DenySettings{
				Mode:               to.Ptr(armdeploymentstacks.DenySettingsModeNone),
				ApplyToChildScopes: to.Ptr(false),
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
	// res = armdeploymentstacks.WhatIfResultsAtResourceGroupClientCreateOrUpdateResponse{
	// 	WhatIfResult: &armdeploymentstacks.WhatIfResult{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Resources/deploymentStacksWhatIfResults/simpleDeploymentStackWhatIfResult"),
	// 		Type: to.Ptr("Microsoft.Resources/deploymentStacksWhatIfResults"),
	// 		Name: to.Ptr("simpleDeploymentStackWhatIfResult"),
	// 		Location: to.Ptr("eastus"),
	// 		SystemData: &armdeploymentstacks.SystemData{
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armdeploymentstacks.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-01T01:01:01.1075056Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armdeploymentstacks.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-02T02:03:01.1974346Z"); return t}()),
	// 		},
	// 		Properties: &armdeploymentstacks.WhatIfResultProperties{
	// 			DeploymentStackResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Resources/deploymentStacks/simpleDeploymentStack"),
	// 			DeploymentStackLastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-01T01:01:01.1075056Z"); return t}()),
	// 			RetentionInterval: to.Ptr("P7D"),
	// 			CorrelationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TemplateLink: &armdeploymentstacks.TemplateLink{
	// 				URI: to.Ptr("https://example.com/exampleTemplate.json"),
	// 			},
	// 			Parameters: map[string]*armdeploymentstacks.DeploymentParameter{
	// 			},
	// 			ActionOnUnmanage: &armdeploymentstacks.ActionOnUnmanage{
	// 				Resources: to.Ptr(armdeploymentstacks.UnmanageActionResourceModeDelete),
	// 				ResourceGroups: to.Ptr(armdeploymentstacks.UnmanageActionResourceGroupModeDelete),
	// 				ManagementGroups: to.Ptr(armdeploymentstacks.UnmanageActionManagementGroupModeDetach),
	// 			},
	// 			DenySettings: &armdeploymentstacks.DenySettings{
	// 				Mode: to.Ptr(armdeploymentstacks.DenySettingsModeNone),
	// 				ApplyToChildScopes: to.Ptr(false),
	// 			},
	// 			ProvisioningState: to.Ptr(armdeploymentstacks.DeploymentStackProvisioningStateInitializing),
	// 		},
	// 	},
	// }
}
