package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_UpdateMobilityService.json
func ExampleReplicationProtectedItemsClient_BeginUpdateMobilityService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("WCUSVault",
		"wcusValidations",
		"b364ed8d-4279-4bf8-8fd1-56f8fa0ae05c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdateMobilityService(ctx,
		"WIN-JKKJ31QI8U2",
		"cloud_c6780228-83bd-4f3e-a70e-cb46b7da33a0",
		"79dd20ab-2b40-11e7-9791-0050568f387e",
		armrecoveryservicessiterecovery.UpdateMobilityServiceRequest{
			Properties: &armrecoveryservicessiterecovery.UpdateMobilityServiceRequestProperties{
				RunAsAccountID: to.Ptr("2"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
