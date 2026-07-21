package armenclave_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/enclave/armenclave"
)

// Generated from example definition: 2026-03-01-preview/Approvals_Update.json
func ExampleApprovalClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armenclave.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApprovalClient().BeginUpdate(ctx, "subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/Microsoft.Mission/enclaveconnections/TestMyEnclaveConnection", "TestApprovals", armenclave.ApprovalPatchModel{
		Properties: &armenclave.ApprovalPatchProperties{
			ParentResourceID:      to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/microsoft.mission/virtualenclaves/TestMyEnclave"),
			GrandparentResourceID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/testrg/providers/Microsoft.Mission/communities/TestMyCommunity"),
			RequestMetadata: &armenclave.RequestMetadataUpdatableProperties{
				ResourceAction:          to.Ptr("string"),
				ApprovalStatus:          to.Ptr(armenclave.ApprovalStatusApproved),
				ApprovalCallbackRoute:   to.Ptr("approvalCallback"),
				ApprovalCallbackPayload: to.Ptr("{\n  \"key1\": \"value1\",\n  \"key2\": \"value2\"\n}"),
			},
			Approvers: []*armenclave.Approver{
				{
					ApproverEntraID: to.Ptr("00000000-0000-0000-0000-000000000000"),
					ActionPerformed: to.Ptr(armenclave.ActionPerformedApproved),
					LastUpdatedAt:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t }()),
				},
			},
			TicketID:       to.Ptr("string"),
			CreatedAt:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t }()),
			StateChangedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t }()),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armenclave.ApprovalClientUpdateResponse{
	// 	ApprovalResource: armenclave.ApprovalResource{
	// 		Properties: &armenclave.ApprovalProperties{
	// 			ParentResourceID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/microsoft.mission/virtualenclaves/TestMyEnclave"),
	// 			GrandparentResourceID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/testrg/providers/Microsoft.Mission/communities/TestMyCommunity"),
	// 			RequestMetadata: &armenclave.RequestMetadata{
	// 				ResourceAction: to.Ptr("string"),
	// 				ApprovalStatus: to.Ptr(armenclave.ApprovalStatusApproved),
	// 				ApprovalCallbackRoute: to.Ptr("approvalCallback"),
	// 				ApprovalCallbackPayload: to.Ptr("{\n  \"key1\": \"value1\",\n  \"key2\": \"value2\"\n}"),
	// 			},
	// 			Approvers: []*armenclave.Approver{
	// 				{
	// 					ApproverEntraID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					ActionPerformed: to.Ptr(armenclave.ActionPerformedApproved),
	// 					LastUpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 					MandatoryApprovalGroupMembershipIDs: []*string{
	// 						to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 						to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 					},
	// 				},
	// 			},
	// 			MandatoryApprovers: []*armenclave.MandatoryApprover{
	// 				{
	// 					ApproverEntraID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 				},
	// 				{
	// 					ApproverEntraID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 				},
	// 			},
	// 			MinimumApproversRequired: to.Ptr[int64](2),
	// 			ApproversApprovedCount: to.Ptr[int64](1),
	// 			MandatoryApproversApprovedCount: to.Ptr[int64](2),
	// 			ApprovedByEntraIDs: []*string{
	// 				to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 			TicketID: to.Ptr("string"),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 			StateChangedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 		},
	// 		ID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/Microsoft.Mission/approvals/TestApprovals"),
	// 		Name: to.Ptr("TestApprovals"),
	// 		Type: to.Ptr("Microsoft.Mission/approvals"),
	// 		SystemData: &armenclave.SystemData{
	// 			CreatedBy: to.Ptr("myAlias"),
	// 			CreatedByType: to.Ptr(armenclave.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("myAlias"),
	// 			LastModifiedByType: to.Ptr(armenclave.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 		},
	// 	},
	// }
}
