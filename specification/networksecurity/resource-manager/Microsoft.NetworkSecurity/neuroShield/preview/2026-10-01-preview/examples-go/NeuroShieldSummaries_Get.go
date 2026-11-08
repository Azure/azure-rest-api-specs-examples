package armnetworksecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networksecurity/armnetworksecurity"
)

// Generated from example definition: 2026-10-01-preview/NeuroShieldSummaries_Get.json
func ExampleNeuroShieldSummariesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworksecurity.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNeuroShieldSummariesClient().Get(ctx, "myResourceGroup", "myNeuroShield", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetworksecurity.NeuroShieldSummariesClientGetResponse{
	// 	NeuroShieldSummary: armnetworksecurity.NeuroShieldSummary{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.NetworkSecurity/neuroShields/myNeuroShield/summary/default"),
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.NetworkSecurity/neuroShields/summary"),
	// 		Properties: &armnetworksecurity.NeuroShieldSummaryProperties{
	// 			ProvisioningState: to.Ptr(armnetworksecurity.ResourceProvisioningStateSucceeded),
	// 			ProtectedCount: to.Ptr[int32](12),
	// 			ScannedProtectedCount: to.Ptr[int32](10),
	// 			ActiveFindings: to.Ptr[int32](8),
	// 			CriticalFindings: to.Ptr[int32](1),
	// 			HighFindings: to.Ptr[int32](2),
	// 			MediumFindings: to.Ptr[int32](3),
	// 			LowFindings: to.Ptr[int32](1),
	// 			UnreachableFindings: to.Ptr[int32](1),
	// 			AwaitingApprovalMitigations: to.Ptr[int32](3),
	// 			QueuedMitigations: to.Ptr[int32](2),
	// 			InProgressMitigations: to.Ptr[int32](1),
	// 			SucceededMitigations: to.Ptr[int32](14),
	// 			FailedMitigations: to.Ptr[int32](2),
	// 			RollbackInProgressMitigations: to.Ptr[int32](1),
	// 		},
	// 	},
	// }
}
