package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-11-15-preview/CreateOrUpdateAttachedAksCluster.json
func ExampleAttachedAKSClustersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAttachedAKSClustersClient().CreateOrUpdate(ctx, "rgcognitiveservices", "myAccount", "myAttachment", armcognitiveservices.AttachedAksCluster{
		Properties: &armcognitiveservices.AttachedAksClusterProperties{
			ClusterResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/rgaks/providers/Microsoft.ContainerService/managedClusters/myAks"),
			Capabilities: []*armcognitiveservices.AttachedAksClusterCapability{
				to.Ptr(armcognitiveservices.AttachedAksClusterCapabilityFoundryNative),
				to.Ptr(armcognitiveservices.AttachedAksClusterCapabilityKuberay),
			},
			ExtensionProfile: &armcognitiveservices.AttachedAksClusterExtensionProfile{
				Version:      to.Ptr("1.0.0"),
				ReleaseTrain: to.Ptr("stable"),
				Configuration: &armcognitiveservices.AttachedAksClusterExtensionConfiguration{
					RayVersion: to.Ptr("2.52.0"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.AttachedAKSClustersClientCreateOrUpdateResponse{
	// 	AttachedAksCluster: armcognitiveservices.AttachedAksCluster{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/attachedAksClusters/myAttachment"),
	// 		Name: to.Ptr("myAttachment"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/attachedAksClusters"),
	// 		Properties: &armcognitiveservices.AttachedAksClusterProperties{
	// 			ClusterResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/rgaks/providers/Microsoft.ContainerService/managedClusters/myAks"),
	// 			Capabilities: []*armcognitiveservices.AttachedAksClusterCapability{
	// 				to.Ptr(armcognitiveservices.AttachedAksClusterCapabilityFoundryNative),
	// 				to.Ptr(armcognitiveservices.AttachedAksClusterCapabilityKuberay),
	// 			},
	// 			Status: &armcognitiveservices.AttachedAksClusterStatus{
	// 				Stage: to.Ptr(armcognitiveservices.AttachedAksClusterProvisioningStageReady),
	// 				LastUpdatedTimeUTC: to.Ptr(time.Date(2026, time.November, 15, 0, 0, 0, 0, time.UTC)),
	// 			},
	// 		},
	// 	},
	// }
}
