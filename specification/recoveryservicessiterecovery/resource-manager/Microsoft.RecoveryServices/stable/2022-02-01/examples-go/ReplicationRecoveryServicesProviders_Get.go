package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationRecoveryServicesProviders_Get.json
func ExampleReplicationRecoveryServicesProvidersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationRecoveryServicesProvidersClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"cloud1",
		"241641e6-ee7b-4ee4-8141-821fadda43fa",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
