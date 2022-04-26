Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicessiterecovery%2Fv0.4.0/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationRecoveryServicesProviders_Create.json
func ExampleReplicationRecoveryServicesProvidersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationRecoveryServicesProvidersClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<fabric-name>",
		"<provider-name>",
		armrecoveryservicessiterecovery.AddRecoveryServicesProviderInput{
			Properties: &armrecoveryservicessiterecovery.AddRecoveryServicesProviderInputProperties{
				AuthenticationIdentityInput: &armrecoveryservicessiterecovery.IdentityProviderInput{
					AADAuthority:  to.Ptr("<aadauthority>"),
					ApplicationID: to.Ptr("<application-id>"),
					Audience:      to.Ptr("<audience>"),
					ObjectID:      to.Ptr("<object-id>"),
					TenantID:      to.Ptr("<tenant-id>"),
				},
				MachineName: to.Ptr("<machine-name>"),
				ResourceAccessIdentityInput: &armrecoveryservicessiterecovery.IdentityProviderInput{
					AADAuthority:  to.Ptr("<aadauthority>"),
					ApplicationID: to.Ptr("<application-id>"),
					Audience:      to.Ptr("<audience>"),
					ObjectID:      to.Ptr("<object-id>"),
					TenantID:      to.Ptr("<tenant-id>"),
				},
			},
		},
		&armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
