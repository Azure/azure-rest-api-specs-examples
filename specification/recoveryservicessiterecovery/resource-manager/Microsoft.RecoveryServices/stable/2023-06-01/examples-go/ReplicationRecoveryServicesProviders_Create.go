package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationRecoveryServicesProviders_Create.json
func ExampleReplicationRecoveryServicesProvidersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationRecoveryServicesProvidersClient().BeginCreate(ctx, "migrationvault", "resourcegroup1", "vmwarefabric1", "vmwareprovider1", armrecoveryservicessiterecovery.AddRecoveryServicesProviderInput{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecoveryServicesProvider = armrecoveryservicessiterecovery.RecoveryServicesProvider{
	// 	Name: to.Ptr("vmwareprovider1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationRecoveryServicesProviders"),
	// 	ID: to.Ptr("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.RecoveryServices/vaults/migrationvault/replicationFabrics/vmwarefabric1/replicationRecoveryServicesProviders/vmwareprovider1"),
	// 	Properties: &armrecoveryservicessiterecovery.RecoveryServicesProviderProperties{
	// 		AllowedScenarios: []*string{
	// 			to.Ptr("Refresh")},
	// 			ConnectionStatus: to.Ptr("Connected"),
	// 			FabricFriendlyName: to.Ptr("vmwarefabric1"),
	// 			FabricType: to.Ptr("VMwareV2"),
	// 			FriendlyName: to.Ptr("vmwareprovider1"),
	// 			LastHeartBeat: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-27T09:16:04.9405768Z"); return t}()),
	// 			ProtectedItemCount: to.Ptr[int32](2),
	// 			ProviderVersion: to.Ptr("5.1.3688.0"),
	// 			ProviderVersionState: to.Ptr("Latest"),
	// 			ServerVersion: to.Ptr("3.2.7510.0"),
	// 		},
	// 	}
}
