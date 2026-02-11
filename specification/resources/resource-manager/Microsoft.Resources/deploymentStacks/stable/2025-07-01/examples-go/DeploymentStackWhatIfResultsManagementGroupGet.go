package armdeploymentstacks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentstacks/v2"
)

// Generated from example definition: 2025-07-01/DeploymentStackWhatIfResultsManagementGroupGet.json
func ExampleWhatIfResultsAtManagementGroupClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentstacks.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWhatIfResultsAtManagementGroupClient().Get(ctx, "myMg", "simpleDeploymentStackWhatIfResult", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdeploymentstacks.WhatIfResultsAtManagementGroupClientGetResponse{
	// 	WhatIfResult: &armdeploymentstacks.WhatIfResult{
	// 		ID: to.Ptr("/providers/Microsoft.Management/managementGroups/myMg/providers/Microsoft.Resources/deploymentStacksWhatIfResults/simpleDeploymentStackWhatIfResult"),
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
	// 			DeploymentStackResourceID: to.Ptr("/providers/Microsoft.Management/managementGroups/myMg/providers/Microsoft.Resources/deploymentStacks/simpleDeploymentStack"),
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
	// 			ProvisioningState: to.Ptr(armdeploymentstacks.DeploymentStackProvisioningStateSucceeded),
	// 			Changes: &armdeploymentstacks.WhatIfChange{
	// 				ResourceChanges: []*armdeploymentstacks.WhatIfResourceChange{
	// 					{
	// 						ID: to.Ptr("/providers/Microsoft.Management/managementGroups/myMg/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment"),
	// 						ManagementStatusChange: &armdeploymentstacks.ChangeBaseDeploymentStacksManagementStatus{
	// 							Before: to.Ptr(armdeploymentstacks.DeploymentStacksManagementStatus("notManaged")),
	// 							After: to.Ptr(armdeploymentstacks.DeploymentStacksManagementStatusManaged),
	// 						},
	// 						ChangeType: to.Ptr(armdeploymentstacks.DeploymentStacksWhatIfChangeTypeCreate),
	// 						ChangeCertainty: to.Ptr(armdeploymentstacks.DeploymentStacksWhatIfChangeCertaintyDefinite),
	// 					},
	// 				},
	// 				DenySettingsChange: &armdeploymentstacks.ChangeDeltaDenySettings{
	// 					After: &armdeploymentstacks.DenySettings{
	// 						Mode: to.Ptr(armdeploymentstacks.DenySettingsModeNone),
	// 						ExcludedPrincipals: []*string{
	// 						},
	// 						ExcludedActions: []*string{
	// 						},
	// 						ApplyToChildScopes: to.Ptr(false),
	// 					},
	// 					Delta: []*armdeploymentstacks.WhatIfPropertyChange{
	// 						{
	// 							Path: to.Ptr("Mode"),
	// 							ChangeType: to.Ptr(armdeploymentstacks.DeploymentStacksWhatIfPropertyChangeTypeCreate),
	// 							After: "none",
	// 						},
	// 						{
	// 							Path: to.Ptr("ApplyToChildScopes"),
	// 							ChangeType: to.Ptr(armdeploymentstacks.DeploymentStacksWhatIfPropertyChangeTypeCreate),
	// 							After: false,
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
