package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/RecoveryPoints_Get.json
func ExampleRecoveryPointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecoveryPointsClient().Get(ctx, "resourceGroupPS1", "vault1", "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "f8491e4f-817a-40dd-a90c-af773978c75b", "b22134ea-620c-474b-9fa5-3c1cb47708e3", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecoveryPoint = armrecoveryservicessiterecovery.RecoveryPoint{
	// 	Name: to.Ptr("b22134ea-620c-474b-9fa5-3c1cb47708e3"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/recoveryPoints"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectedItems/f8491e4f-817a-40dd-a90c-af773978c75b/recoveryPoints/b22134ea-620c-474b-9fa5-3c1cb47708e3"),
	// 	Properties: &armrecoveryservicessiterecovery.RecoveryPointProperties{
	// 		RecoveryPointTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-26T06:37:50.808Z"); return t}()),
	// 		RecoveryPointType: to.Ptr("CrashConsistent"),
	// 	},
	// }
}
