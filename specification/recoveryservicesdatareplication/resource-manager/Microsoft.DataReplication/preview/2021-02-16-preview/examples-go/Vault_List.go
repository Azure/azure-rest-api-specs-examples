package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Vault_List.json
func ExampleVaultClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVaultClient().NewListPager("rgrecoveryservicesdatareplication", &armrecoveryservicesdatareplication.VaultClientListOptions{ContinuationToken: to.Ptr("mwculdaqndp")})
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
		// page.VaultModelCollection = armrecoveryservicesdatareplication.VaultModelCollection{
		// 	Value: []*armrecoveryservicesdatareplication.VaultModel{
		// 		{
		// 			Name: to.Ptr("bqgyqxmnlgwqxbmajddqwtao"),
		// 			Type: to.Ptr("xtcicpcmjvocohaznrk"),
		// 			ID: to.Ptr("hbccseewgiagaxsjozxgwydjgy"),
		// 			Location: to.Ptr("eck"),
		// 			Properties: &armrecoveryservicesdatareplication.VaultModelProperties{
		// 				ProvisioningState: to.Ptr(armrecoveryservicesdatareplication.ProvisioningStateSucceeded),
		// 				ServiceResourceID: to.Ptr("mksumcmksgdsghojszxq"),
		// 				VaultType: to.Ptr(armrecoveryservicesdatareplication.ReplicationVaultTypeDisasterRecovery),
		// 			},
		// 			SystemData: &armrecoveryservicesdatareplication.VaultModelSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:58.092Z"); return t}()),
		// 				CreatedBy: to.Ptr("rm"),
		// 				CreatedByType: to.Ptr("uojlfokjrhzgqoodsvgz"),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:58.092Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("gkojzu"),
		// 				LastModifiedByType: to.Ptr("jua"),
		// 			},
		// 			Tags: map[string]*string{
		// 				"key5359": to.Ptr("ljfilxolxzuxrauopwtyxghrp"),
		// 			},
		// 	}},
		// }
	}
}
