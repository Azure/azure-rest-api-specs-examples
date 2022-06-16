package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationVaultHealth_Get.json
func ExampleReplicationVaultHealthClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationVaultHealthClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("VaultHealthDetails.ID: %s\n", *res.ID)
}
