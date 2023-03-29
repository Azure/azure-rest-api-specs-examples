package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollout_post_cancel.json
func ExampleRolloutsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRolloutsClient().Cancel(ctx, "myResourceGroup", "myRollout", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Rollout = armdeploymentmanager.Rollout{
	// 	Name: to.Ptr("myRollout"),
	// 	Type: to.Ptr("Microsoft.DeploymentManager/rollouts"),
	// 	Location: to.Ptr("centralus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armdeploymentmanager.Identity{
	// 		Type: to.Ptr("userAssigned"),
	// 		IdentityIDs: []*string{
	// 			to.Ptr("/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userassignedidentities/myuseridentity")},
	// 		},
	// 		Properties: &armdeploymentmanager.RolloutProperties{
	// 			OperationInfo: &armdeploymentmanager.RolloutOperationInfo{
	// 				RetryAttempt: to.Ptr[int32](0),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-28T03:33:56.386Z"); return t}()),
	// 			},
	// 			Services: []*armdeploymentmanager.Service{
	// 				{
	// 					TargetLocation: to.Ptr("centralus"),
	// 					TargetSubscriptionID: to.Ptr("600c95c5-3ee5-44fe-b190-ca38a19adcd7"),
	// 					Name: to.Ptr("myService1"),
	// 					ServiceUnits: []*armdeploymentmanager.ServiceUnit{
	// 						{
	// 							DeploymentMode: to.Ptr(armdeploymentmanager.DeploymentModeIncremental),
	// 							TargetResourceGroup: to.Ptr("myDeploymentResourceGroup"),
	// 							Name: to.Ptr("myTopologyUni1"),
	// 							Steps: []*armdeploymentmanager.RolloutStep{
	// 								{
	// 									Name: to.Ptr("preDeploymentStep1"),
	// 									OperationInfo: &armdeploymentmanager.StepOperationInfo{
	// 										EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-28T03:35:28.556Z"); return t}()),
	// 										StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-28T03:33:56.386Z"); return t}()),
	// 									},
	// 									Status: to.Ptr("succeeded"),
	// 								},
	// 								{
	// 									Name: to.Ptr("preDeploymentStep2"),
	// 									OperationInfo: &armdeploymentmanager.StepOperationInfo{
	// 										EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-28T03:37:28.556Z"); return t}()),
	// 										StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-28T03:36:56.386Z"); return t}()),
	// 									},
	// 									Status: to.Ptr("succeeded"),
	// 								},
	// 								{
	// 									Name: to.Ptr("deploy"),
	// 									OperationInfo: &armdeploymentmanager.StepOperationInfo{
	// 										LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-28T03:39:28.556Z"); return t}()),
	// 										StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-28T03:38:56.386Z"); return t}()),
	// 									},
	// 									ResourceOperations: []*armdeploymentmanager.ResourceOperation{
	// 										{
	// 											OperationID: to.Ptr("20FC5A21382DA306"),
	// 											ProvisioningState: to.Ptr("Succeeded"),
	// 											ResourceName: to.Ptr("keyVaultcentralus"),
	// 											ResourceType: to.Ptr("Microsoft.KeyVault/vaults"),
	// 											StatusCode: to.Ptr("OK"),
	// 											StatusMessage: to.Ptr(""),
	// 									}},
	// 									Status: to.Ptr("running"),
	// 							}},
	// 					}},
	// 			}},
	// 			Status: to.Ptr("Canceling"),
	// 			TotalRetryAttempts: to.Ptr[int32](0),
	// 			ArtifactSourceID: to.Ptr("/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/Microsoft.DeploymentManager/artifactSources/myArtifactSource"),
	// 			BuildVersion: to.Ptr("1.0.0.1"),
	// 			StepGroups: []*armdeploymentmanager.StepGroup{
	// 				{
	// 					Name: to.Ptr("FirstRegion"),
	// 					DeploymentTargetID: to.Ptr("Microsoft.DeploymentManager/serviceTopologies/myTopology/services/myService/serviceUnits/myServiceUnit1'"),
	// 					PostDeploymentSteps: []*armdeploymentmanager.PrePostStep{
	// 						{
	// 							StepID: to.Ptr("Microsoft.DeploymentManager/steps/postDeployStep1"),
	// 					}},
	// 					PreDeploymentSteps: []*armdeploymentmanager.PrePostStep{
	// 						{
	// 							StepID: to.Ptr("Microsoft.DeploymentManager/steps/preDeployStep1"),
	// 						},
	// 						{
	// 							StepID: to.Ptr("Microsoft.DeploymentManager/steps/preDeployStep2"),
	// 					}},
	// 				},
	// 				{
	// 					Name: to.Ptr("SecondRegion"),
	// 					DependsOnStepGroups: []*string{
	// 						to.Ptr("FirstRegion")},
	// 						DeploymentTargetID: to.Ptr("Microsoft.DeploymentManager/serviceTopologies/myTopology/services/myService/serviceUnits/myServiceUnit2'"),
	// 						PostDeploymentSteps: []*armdeploymentmanager.PrePostStep{
	// 							{
	// 								StepID: to.Ptr("Microsoft.DeploymentManager/steps/postDeployStep5"),
	// 						}},
	// 						PreDeploymentSteps: []*armdeploymentmanager.PrePostStep{
	// 							{
	// 								StepID: to.Ptr("Microsoft.DeploymentManager/steps/preDeployStep3"),
	// 							},
	// 							{
	// 								StepID: to.Ptr("Microsoft.DeploymentManager/steps/preDeployStep4"),
	// 						}},
	// 				}},
	// 				TargetServiceTopologyID: to.Ptr("/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/Microsoft.DeploymentManager/serviceTopologies/myTopology"),
	// 			},
	// 		}
}
