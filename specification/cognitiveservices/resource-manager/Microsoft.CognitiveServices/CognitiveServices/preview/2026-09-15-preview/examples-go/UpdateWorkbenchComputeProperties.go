package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-09-15-preview/UpdateWorkbenchComputeProperties.json
func ExampleWorkbenchesClient_Update_updateWorkbenchComputeProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkbenchesClient().Update(ctx, "rgcognitiveservices", "myAccount", "myProject", "myWorkbench", armcognitiveservices.WorkbenchUpdate{
		Properties: &armcognitiveservices.WorkbenchUpdateProperties{
			TargetClusterID:        to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/vc-rg/providers/Microsoft.MachineLearningServices/virtualClusters/test-vc"),
			IdleTimeBeforeShutdown: to.Ptr("PT1H"),
			InstanceType:           to.Ptr("Singularity.ND12_H100_v5-n1"),
			GpuCount:               to.Ptr[int32](1),
		},
		Identity: &armcognitiveservices.Identity{
			Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armcognitiveservices.UserAssignedIdentity{
				"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity": {},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.WorkbenchesClientUpdateResponse{
	// 	Workbench: armcognitiveservices.Workbench{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/projects/myProject/workbenches/myWorkbench"),
	// 		Name: to.Ptr("myWorkbench"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects/workbenches"),
	// 		Properties: &armcognitiveservices.WorkbenchProperties{
	// 			TargetClusterID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/vc-rg/providers/Microsoft.MachineLearningServices/virtualClusters/test-vc"),
	// 			ImageLink: to.Ptr("mcr.microsoft.com/azureml/curated/pytorch-gpu:latest"),
	// 			IdleTimeBeforeShutdown: to.Ptr("PT1H"),
	// 			InstanceType: to.Ptr("Singularity.ND12_H100_v5-n1"),
	// 			GpuCount: to.Ptr[int32](1),
	// 			ProvisioningState: to.Ptr(armcognitiveservices.WorkbenchProvisioningStateSucceeded),
	// 			Status: to.Ptr(armcognitiveservices.WorkbenchStatusStopped),
	// 		},
	// 		Identity: &armcognitiveservices.Identity{
	// 			Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armcognitiveservices.UserAssignedIdentity{
	// 				"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity": &armcognitiveservices.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000002"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
