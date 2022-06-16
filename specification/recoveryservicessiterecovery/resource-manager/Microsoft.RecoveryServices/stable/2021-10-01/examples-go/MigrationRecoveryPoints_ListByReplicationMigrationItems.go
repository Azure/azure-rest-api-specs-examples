package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/MigrationRecoveryPoints_ListByReplicationMigrationItems.json
func ExampleMigrationRecoveryPointsClient_ListByReplicationMigrationItems() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewMigrationRecoveryPointsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	pager := client.ListByReplicationMigrationItems("<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("MigrationRecoveryPoint.ID: %s\n", *v.ID)
		}
	}
}
