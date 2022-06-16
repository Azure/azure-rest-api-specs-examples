package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationRecoveryServicesProviders_Create.json
func ExampleReplicationRecoveryServicesProvidersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationRecoveryServicesProvidersClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<fabric-name>",
		"<provider-name>",
		armrecoveryservicessiterecovery.AddRecoveryServicesProviderInput{
			Properties: &armrecoveryservicessiterecovery.AddRecoveryServicesProviderInputProperties{
				AuthenticationIdentityInput: &armrecoveryservicessiterecovery.IdentityProviderInput{
					AADAuthority:  to.StringPtr("<aadauthority>"),
					ApplicationID: to.StringPtr("<application-id>"),
					Audience:      to.StringPtr("<audience>"),
					ObjectID:      to.StringPtr("<object-id>"),
					TenantID:      to.StringPtr("<tenant-id>"),
				},
				MachineName: to.StringPtr("<machine-name>"),
				ResourceAccessIdentityInput: &armrecoveryservicessiterecovery.IdentityProviderInput{
					AADAuthority:  to.StringPtr("<aadauthority>"),
					ApplicationID: to.StringPtr("<application-id>"),
					Audience:      to.StringPtr("<audience>"),
					ObjectID:      to.StringPtr("<object-id>"),
					TenantID:      to.StringPtr("<tenant-id>"),
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
	log.Printf("Response result: %#v\n", res.ReplicationRecoveryServicesProvidersClientCreateResult)
}
