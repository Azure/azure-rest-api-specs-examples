package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationAlertSettings_Create.json
func ExampleReplicationAlertSettingsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationAlertSettingsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<alert-setting-name>",
		armrecoveryservicessiterecovery.ConfigureAlertRequest{
			Properties: &armrecoveryservicessiterecovery.ConfigureAlertRequestProperties{
				CustomEmailAddresses: []*string{
					to.StringPtr("ronehr@microsoft.com")},
				Locale:       to.StringPtr("<locale>"),
				SendToOwners: to.StringPtr("<send-to-owners>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ReplicationAlertSettingsClientCreateResult)
}
