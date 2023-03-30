package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollout_post_restart.json
func ExampleRolloutsClient_Restart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRolloutsClient().Restart(ctx, "myResourceGroup", "myRollout", &armdeploymentmanager.RolloutsClientRestartOptions{SkipSucceeded: to.Ptr(true)})
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
	// 				RetryAttempt: to.Ptr[int32](1),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-28T03:33:56.386Z"); return t}()),
	// 			},
	// 			Status: to.Ptr("Running"),
	// 			TotalRetryAttempts: to.Ptr[int32](1),
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
