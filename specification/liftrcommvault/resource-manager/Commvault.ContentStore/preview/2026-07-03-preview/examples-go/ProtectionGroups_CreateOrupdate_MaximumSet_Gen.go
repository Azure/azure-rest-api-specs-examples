package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-07-03-preview/ProtectionGroups_CreateOrupdate_MaximumSet_Gen.json
func ExampleProtectionGroupsClient_BeginCreateOrupdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("65D4E6D7-7063-4C4B-BAC5-13C45474009E", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProtectionGroupsClient().BeginCreateOrupdate(ctx, "rgcommvault", "sample-cloudAccountName", "sample-protectionGroupName", armcommvaultcontentstore.ProtectionGroup{
		Properties: &armcommvaultcontentstore.ProtectionGroupProperties{
			DataSourceType: to.Ptr("AzureVM"),
			Plan:           to.Ptr("ibcuuodwnnvgyhy"),
			Resources: &armcommvaultcontentstore.ProtectionGroupResources{
				Manual: []*string{
					to.Ptr("uljwtwhm"),
				},
				MatchRules: &armcommvaultcontentstore.ProtectionGroupResourcesMatchRules{
					Rules: []*armcommvaultcontentstore.Rule{
						{
							Property: to.Ptr(armcommvaultcontentstore.RulePropertyResourceGroup),
							Operator: to.Ptr(armcommvaultcontentstore.OperatorContains),
							Value:    to.Ptr("dgkmghsgmrbaatklarukbx"),
						},
					},
					MatchType: to.Ptr(armcommvaultcontentstore.MatchTypeAll),
				},
			},
			ProtectionStatus: to.Ptr(armcommvaultcontentstore.ProtectionStatusAll),
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
	// res = armcommvaultcontentstore.ProtectionGroupsClientCreateOrupdateResponse{
	// 	ProtectionGroup: armcommvaultcontentstore.ProtectionGroup{
	// 		Properties: &armcommvaultcontentstore.ProtectionGroupProperties{
	// 			DataSourceType: to.Ptr("AzureVM"),
	// 			Plan: to.Ptr("ibcuuodwnnvgyhy"),
	// 			Resources: &armcommvaultcontentstore.ProtectionGroupResources{
	// 				Manual: []*string{
	// 					to.Ptr("uljwtwhm"),
	// 				},
	// 				MatchRules: &armcommvaultcontentstore.ProtectionGroupResourcesMatchRules{
	// 					Rules: []*armcommvaultcontentstore.Rule{
	// 						{
	// 							Property: to.Ptr(armcommvaultcontentstore.RulePropertyResourceGroup),
	// 							Operator: to.Ptr(armcommvaultcontentstore.OperatorContains),
	// 							Value: to.Ptr("dgkmghsgmrbaatklarukbx"),
	// 						},
	// 					},
	// 					MatchType: to.Ptr(armcommvaultcontentstore.MatchTypeAll),
	// 				},
	// 			},
	// 			ProtectionStatus: to.Ptr(armcommvaultcontentstore.ProtectionStatusAll),
	// 			NumberOfProtectedItems: to.Ptr[int32](11),
	// 			LastBackUpTime: to.Ptr[int64](28),
	// 			BackupActivityStatus: to.Ptr("wqmqejszoyvspktbbo"),
	// 			ProvisioningState: to.Ptr(armcommvaultcontentstore.ResourceProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/65D4E6D7-7063-4C4B-BAC5-13C45474009E/resourceGroups/rgcommvault/providers/Commvault.ContentStore/cloudAccounts/myCloudAccount/protectionGroups/myProtectionGroup"),
	// 		Name: to.Ptr("iaqmjpznbsz"),
	// 		Type: to.Ptr("japjizreaoplcjpzyrdyi"),
	// 		SystemData: &armcommvaultcontentstore.SystemData{
	// 			CreatedBy: to.Ptr("wg"),
	// 			CreatedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-01T06:14:57.355Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("hbpzxzzwhqfy"),
	// 			LastModifiedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-01T06:14:57.355Z"); return t}()),
	// 		},
	// 	},
	// }
}
