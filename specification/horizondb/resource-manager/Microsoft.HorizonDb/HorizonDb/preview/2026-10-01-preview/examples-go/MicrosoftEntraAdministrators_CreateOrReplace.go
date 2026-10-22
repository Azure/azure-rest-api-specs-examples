package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/MicrosoftEntraAdministrators_CreateOrReplace.json
func ExampleMicrosoftEntraAdministratorsClient_BeginCreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMicrosoftEntraAdministratorsClient().BeginCreateOrReplace(ctx, "example-resource-group", "example-cluster", "11111111-1111-1111-1111-111111111111", armhorizondb.MicrosoftEntraAdministratorAdd{
		Properties: &armhorizondb.MicrosoftEntraAdministratorPropertiesForAdd{
			PrincipalName: to.Ptr("database.administrators@contoso.com"),
			PrincipalType: to.Ptr(armhorizondb.PrincipalTypesGroup),
			TenantID:      to.Ptr("22222222-2222-2222-2222-222222222222"),
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
	// res = armhorizondb.MicrosoftEntraAdministratorsClientCreateOrReplaceResponse{
	// 	MicrosoftEntraAdministrator: armhorizondb.MicrosoftEntraAdministrator{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/microsoftEntraAdministrators/11111111-1111-1111-1111-111111111111"),
	// 		Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters/microsoftEntraAdministrators"),
	// 		Properties: &armhorizondb.MicrosoftEntraAdministratorProperties{
	// 			PrincipalName: to.Ptr("database.administrators@contoso.com"),
	// 			PrincipalType: to.Ptr(armhorizondb.PrincipalTypesGroup),
	// 			ObjectID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 			TenantID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateProvisioning),
	// 		},
	// 	},
	// }
}
