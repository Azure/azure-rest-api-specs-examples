package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: 2024-09-01/Vault_ListBySubscription.json
func ExampleVaultClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("930CEC23-4430-4513-B855-DBA237E2F3BF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVaultClient().NewListBySubscriptionPager(nil)
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
		// page = armrecoveryservicesdatareplication.VaultClientListBySubscriptionResponse{
		// 	VaultModelListResult: armrecoveryservicesdatareplication.VaultModelListResult{
		// 		Value: []*armrecoveryservicesdatareplication.VaultModel{
		// 			{
		// 				Location: to.Ptr("eck"),
		// 				Properties: &armrecoveryservicesdatareplication.VaultModelProperties{
		// 					ProvisioningState: to.Ptr(armrecoveryservicesdatareplication.ProvisioningStateCanceled),
		// 					ServiceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataReplication/replicationVaults/vault1"),
		// 					VaultType: to.Ptr(armrecoveryservicesdatareplication.ReplicationVaultTypeDisasterRecovery),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DataReplication/replicationVaults/vault1"),
		// 				Name: to.Ptr("bqgyqxmnlgwqxbmajddqwtao"),
		// 				Type: to.Ptr("xtcicpcmjvocohaznrk"),
		// 				SystemData: &armrecoveryservicesdatareplication.SystemData{
		// 					CreatedBy: to.Ptr("rm"),
		// 					CreatedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("uojlfokjrhzgqoodsvgz")),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:58.092Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("gkojzu"),
		// 					LastModifiedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("jua")),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:58.092Z"); return t}()),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key5359": to.Ptr("ljfilxolxzuxrauopwtyxghrp"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
