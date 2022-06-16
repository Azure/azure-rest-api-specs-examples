package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationProtectionIntents_Get.json
func ExampleReplicationProtectionIntentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationProtectionIntentsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<intent-object-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ReplicationProtectionIntentsClientGetResult)
}
