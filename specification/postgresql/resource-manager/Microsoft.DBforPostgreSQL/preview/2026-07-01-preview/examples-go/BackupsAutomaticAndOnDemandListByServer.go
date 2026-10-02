package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-07-01-preview/BackupsAutomaticAndOnDemandListByServer.json
func ExampleBackupsAutomaticAndOnDemandClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupsAutomaticAndOnDemandClient().NewListByServerPager("exampleresourcegroup", "exampleserver", nil)
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
		// page = armpostgresqlflexibleservers.BackupsAutomaticAndOnDemandClientListByServerResponse{
		// 	BackupAutomaticAndOnDemandList: armpostgresqlflexibleservers.BackupAutomaticAndOnDemandList{
		// 		Value: []*armpostgresqlflexibleservers.BackupAutomaticAndOnDemand{
		// 			{
		// 				Name: to.Ptr("backup_638830782181266873"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/backups/backup_638830782181266873"),
		// 				Properties: &armpostgresqlflexibleservers.BackupAutomaticAndOnDemandProperties{
		// 					BackupType: to.Ptr(armpostgresqlflexibleservers.BackupTypeFull),
		// 					CompletedTime: to.Ptr(time.Date(2025, time.June, 1, 14, 30, 22, 123456000, time.UTC)),
		// 					Source: to.Ptr("Automatic"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("ondemandbackup-20250601T183022"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/backups/ondemandbackup-20250601T183022"),
		// 				Properties: &armpostgresqlflexibleservers.BackupAutomaticAndOnDemandProperties{
		// 					BackupType: to.Ptr(armpostgresqlflexibleservers.BackupTypeCustomerOnDemand),
		// 					CompletedTime: to.Ptr(time.Date(2025, time.June, 1, 18, 30, 22, 123456000, time.UTC)),
		// 					Source: to.Ptr("Customer Initiated"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
