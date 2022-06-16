package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationProtectionContainerMappings_Update.json
func ExampleReplicationProtectionContainerMappingsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationProtectionContainerMappingsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<mapping-name>",
		armrecoveryservicessiterecovery.UpdateProtectionContainerMappingInput{
			Properties: &armrecoveryservicessiterecovery.UpdateProtectionContainerMappingInputProperties{
				ProviderSpecificInput: &armrecoveryservicessiterecovery.A2AUpdateContainerMappingInput{
					ReplicationProviderSpecificUpdateContainerMappingInput: armrecoveryservicessiterecovery.ReplicationProviderSpecificUpdateContainerMappingInput{
						InstanceType: to.StringPtr("<instance-type>"),
					},
					AgentAutoUpdateStatus:  armrecoveryservicessiterecovery.AgentAutoUpdateStatusEnabled.ToPtr(),
					AutomationAccountArmID: to.StringPtr("<automation-account-arm-id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ProtectionContainerMapping.ID: %s\n", *res.ID)
}
