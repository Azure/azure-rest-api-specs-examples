package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/PasswordAuthentications_CreateOrReplace.json
func ExamplePasswordAuthenticationsClient_BeginCreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPasswordAuthenticationsClient().BeginCreateOrReplace(ctx, "example-resource-group", "example-cluster", armhorizondb.PasswordAuthentication{
		Properties: &armhorizondb.PasswordAuthenticationProperties{
			State:                      to.Ptr(armhorizondb.AuthenticationStateEnabled),
			AdministratorLogin:         to.Ptr("adminlogin1"),
			AdministratorLoginPassword: to.Ptr("********"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhorizondb.PasswordAuthenticationsClientCreateOrReplaceResponse{
	// 	PasswordAuthentication: armhorizondb.PasswordAuthentication{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/authentications/passwordMethod"),
	// 		Name: to.Ptr("passwordMethod"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters/authentications"),
	// 		Properties: &armhorizondb.PasswordAuthenticationProperties{
	// 			State: to.Ptr(armhorizondb.AuthenticationStateEnabled),
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
