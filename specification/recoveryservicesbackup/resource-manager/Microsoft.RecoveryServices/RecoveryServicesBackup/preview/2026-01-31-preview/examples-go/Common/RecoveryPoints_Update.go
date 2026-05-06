package armrecoveryservicesbackup_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
)

// Generated from example definition: 2026-01-31-preview/Common/RecoveryPoints_Update.json
func ExampleRecoveryPointsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecoveryPointsClient().Update(ctx, "hanasnapshottesting", "HanaSnapshotTest", "Azure", "VMAppContainer;compute;hanasnapshottesting;hana-eacan-2", "SAPHanaDatabase;hye;hye", "2265668074516978193", armrecoveryservicesbackup.UpdateRecoveryPointRequest{
		Properties: &armrecoveryservicesbackup.PatchRecoveryPointInput{
			RecoveryPointProperties: &armrecoveryservicesbackup.PatchRecoveryPointPropertiesInput{
				ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-02T00:00:00.0000000Z"); return t }()),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicesbackup.RecoveryPointsClientUpdateResponse{
	// 	RecoveryPointResource: armrecoveryservicesbackup.RecoveryPointResource{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hanasnapshottesting/providers/Microsoft.RecoveryServices/vaults/HanaSnapshotTest/backupFabrics/Azure/protectionContainers/VMAppContainer;compute;hanasnapshottesting;hana-eacan-2/protectedItems/SAPHanaDatabase;hye;hye/recoveryPoints/2265668074516978193"),
	// 		Name: to.Ptr("2265668074516978193"),
	// 		Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints"),
	// 		Properties: &armrecoveryservicesbackup.AzureWorkloadSAPHanaRecoveryPoint{
	// 			ObjectType: to.Ptr("AzureWorkloadSAPHanaRecoveryPoint"),
	// 			RecoveryPointProperties: &armrecoveryservicesbackup.RecoveryPointProperties{
	// 				ExpiryTime: to.Ptr("2025-01-02T00:00:00.0000000Z"),
	// 				IsSoftDeleted: to.Ptr(false),
	// 				RuleName: to.Ptr("On-Demand"),
	// 			},
	// 			RecoveryPointTierDetails: []*armrecoveryservicesbackup.RecoveryPointTierInformationV2{
	// 				{
	// 					Status: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierStatusValid),
	// 					Type: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeHardenedRP),
	// 				},
	// 			},
	// 			RecoveryPointTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-18T06:26:36.922Z"); return t}()),
	// 			Type: to.Ptr(armrecoveryservicesbackup.RestorePointTypeFull),
	// 		},
	// 	},
	// }
}
