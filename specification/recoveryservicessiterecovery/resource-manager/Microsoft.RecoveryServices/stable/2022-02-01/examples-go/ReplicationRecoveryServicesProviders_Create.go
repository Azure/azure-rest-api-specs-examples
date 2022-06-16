package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationRecoveryServicesProviders_Create.json
func ExampleReplicationRecoveryServicesProvidersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationRecoveryServicesProvidersClient("migrationvault",
		"resourcegroup1",
		"cb53d0c3-bd59-4721-89bc-06916a9147ef", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"vmwarefabric1",
		"vmwareprovider1",
		armrecoveryservicessiterecovery.AddRecoveryServicesProviderInput{
			Properties: &armrecoveryservicessiterecovery.AddRecoveryServicesProviderInputProperties{
				AuthenticationIdentityInput: &armrecoveryservicessiterecovery.IdentityProviderInput{
					AADAuthority:  to.Ptr("https://login.microsoftonline.com"),
					ApplicationID: to.Ptr("f66fce08-c0c6-47a1-beeb-0ede5ea94f90"),
					Audience:      to.Ptr("https://microsoft.onmicrosoft.com/cf19e349-644c-4c6a-bcae-9c8f35357874"),
					ObjectID:      to.Ptr("141360b8-5686-4240-a027-5e24e6affeba"),
					TenantID:      to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
				},
				MachineName: to.Ptr("vmwareprovider1"),
				ResourceAccessIdentityInput: &armrecoveryservicessiterecovery.IdentityProviderInput{
					AADAuthority:  to.Ptr("https://login.microsoftonline.com"),
					ApplicationID: to.Ptr("f66fce08-c0c6-47a1-beeb-0ede5ea94f90"),
					Audience:      to.Ptr("https://microsoft.onmicrosoft.com/cf19e349-644c-4c6a-bcae-9c8f35357874"),
					ObjectID:      to.Ptr("141360b8-5686-4240-a027-5e24e6affeba"),
					TenantID:      to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
				},
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
