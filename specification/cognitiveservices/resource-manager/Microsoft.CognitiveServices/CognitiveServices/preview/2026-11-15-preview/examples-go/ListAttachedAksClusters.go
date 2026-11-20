package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-11-15-preview/ListAttachedAksClusters.json
func ExampleAttachedAKSClustersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAttachedAKSClustersClient().NewListPager("rgcognitiveservices", "myAccount", &armcognitiveservices.AttachedAKSClustersClientListOptions{
		Filter:  to.Ptr("stage eq 'Ready'"),
		Orderby: to.Ptr("createdTime desc"),
		Top:     to.Ptr[int32](1),
		Skip:    to.Ptr[int32](0)})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcognitiveservices.AttachedAKSClustersClientListResponse{
		// 	AttachedAksClusterListResult: armcognitiveservices.AttachedAksClusterListResult{
		// 		Value: []*armcognitiveservices.AttachedAksClusterListItem{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/attachedAksClusters/myAttachment"),
		// 				Name: to.Ptr("myAttachment"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/accounts/attachedAksClusters"),
		// 				Properties: &armcognitiveservices.AttachedAksClusterListProperties{
		// 					ClusterResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/rgaks/providers/Microsoft.ContainerService/managedClusters/myAks"),
		// 					Capabilities: []*armcognitiveservices.AttachedAksClusterCapability{
		// 						to.Ptr(armcognitiveservices.AttachedAksClusterCapabilityFoundryNative),
		// 						to.Ptr(armcognitiveservices.AttachedAksClusterCapabilityKuberay),
		// 					},
		// 					Stage: to.Ptr(armcognitiveservices.AttachedAksClusterProvisioningStageReady),
		// 					Health: &armcognitiveservices.AttachedAksClusterHealthStatus{
		// 						State: to.Ptr(armcognitiveservices.AttachedAksClusterHealthStateHealthy),
		// 						Summary: to.Ptr("The attachment is ready for workloads."),
		// 						NewWorkAdmissionAllowed: to.Ptr(true),
		// 						Components: []*armcognitiveservices.AttachedAksClusterHealthComponent{
		// 						},
		// 						Reasons: []*armcognitiveservices.AttachedAksClusterHealthReason{
		// 						},
		// 					},
		// 					ProjectLocation: to.Ptr("eastus2"),
		// 					CreatedTime: to.Ptr(time.Date(2026, time.November, 15, 0, 0, 0, 0, time.UTC)),
		// 					ModifiedTime: to.Ptr(time.Date(2026, time.November, 15, 0, 0, 0, 0, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/attachedAksClusters?api-version=2026-11-15-preview&$filter=stage%20eq%20%27Ready%27&$orderby=createdTime%20desc&$top=1&$skip=1"),
		// 	},
		// }
	}
}
