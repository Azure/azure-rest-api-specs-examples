Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdeploymentmanager%2Farmdeploymentmanager%2Fv0.4.0/sdk/resourcemanager/deploymentmanager/armdeploymentmanager/README.md) on how to add the SDK to your project and authenticate.

```go
package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollout_createorupdate.json
func ExampleRolloutsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdeploymentmanager.NewRolloutsClient("caac1590-e859-444f-a9e0-62091c0f5929", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myRollout",
		&armdeploymentmanager.RolloutsClientBeginCreateOrUpdateOptions{RolloutRequest: &armdeploymentmanager.RolloutRequest{
			Location: to.Ptr("centralus"),
			Tags:     map[string]*string{},
			Identity: &armdeploymentmanager.Identity{
				Type: to.Ptr("userAssigned"),
				IdentityIDs: []*string{
					to.Ptr("/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userassignedidentities/myuseridentity")},
			},
			Properties: &armdeploymentmanager.RolloutRequestProperties{
				ArtifactSourceID: to.Ptr("/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/Microsoft.DeploymentManager/artifactSources/myArtifactSource"),
				BuildVersion:     to.Ptr("1.0.0.1"),
				StepGroups: []*armdeploymentmanager.StepGroup{
					{
						Name:               to.Ptr("FirstRegion"),
						DeploymentTargetID: to.Ptr("Microsoft.DeploymentManager/serviceTopologies/myTopology/services/myService/serviceUnits/myServiceUnit1'"),
						PostDeploymentSteps: []*armdeploymentmanager.PrePostStep{
							{
								StepID: to.Ptr("Microsoft.DeploymentManager/steps/postDeployStep1"),
							}},
						PreDeploymentSteps: []*armdeploymentmanager.PrePostStep{
							{
								StepID: to.Ptr("Microsoft.DeploymentManager/steps/preDeployStep1"),
							},
							{
								StepID: to.Ptr("Microsoft.DeploymentManager/steps/preDeployStep2"),
							}},
					},
					{
						Name: to.Ptr("SecondRegion"),
						DependsOnStepGroups: []*string{
							to.Ptr("FirstRegion")},
						DeploymentTargetID: to.Ptr("Microsoft.DeploymentManager/serviceTopologies/myTopology/services/myService/serviceUnits/myServiceUnit2'"),
						PostDeploymentSteps: []*armdeploymentmanager.PrePostStep{
							{
								StepID: to.Ptr("Microsoft.DeploymentManager/steps/postDeployStep5"),
							}},
						PreDeploymentSteps: []*armdeploymentmanager.PrePostStep{
							{
								StepID: to.Ptr("Microsoft.DeploymentManager/steps/preDeployStep3"),
							},
							{
								StepID: to.Ptr("Microsoft.DeploymentManager/steps/preDeployStep4"),
							}},
					}},
				TargetServiceTopologyID: to.Ptr("/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/Microsoft.DeploymentManager/serviceTopologies/myTopology"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
```
