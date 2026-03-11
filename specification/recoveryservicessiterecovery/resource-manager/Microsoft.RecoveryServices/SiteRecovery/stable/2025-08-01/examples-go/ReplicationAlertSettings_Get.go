package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v3"
)

// Generated from example definition: 2025-08-01/ReplicationAlertSettings_Get.json
func ExampleReplicationAlertSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationAlertSettingsClient().Get(ctx, "resourceGroupPS1", "vault1", "defaultAlertSetting", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicessiterecovery.ReplicationAlertSettingsClientGetResponse{
	// 	Alert: &armrecoveryservicessiterecovery.Alert{
	// 		Name: to.Ptr("defaultAlertSetting"),
	// 		Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationAlertSettings"),
	// 		ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationAlertSettings/defaultAlertSetting"),
	// 		Properties: &armrecoveryservicessiterecovery.AlertProperties{
	// 			CustomEmailAddresses: []*string{
	// 				to.Ptr("ronehr@microsoft.com"),
	// 			},
	// 			Locale: to.Ptr("en-US"),
	// 			SendToOwners: to.Ptr("false"),
	// 		},
	// 	},
	// }
}
