package armmysql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerAdminCreateUpdate.json
func ExampleServerAdministratorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerAdministratorsClient().BeginCreateOrUpdate(ctx, "testrg", "mysqltestsvc4", armmysql.ServerAdministratorResource{
		Properties: &armmysql.ServerAdministratorProperties{
			AdministratorType: to.Ptr("ActiveDirectory"),
			Login:             to.Ptr("bob@contoso.com"),
			Sid:               to.Ptr("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
			TenantID:          to.Ptr("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
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
	// res.ServerAdministratorResource = armmysql.ServerAdministratorResource{
	// 	Name: to.Ptr("activeDirectory"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers/administrators"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc4/administrators/activeDirectory"),
	// 	Properties: &armmysql.ServerAdministratorProperties{
	// 		AdministratorType: to.Ptr("ActiveDirectory"),
	// 		Login: to.Ptr("bob@contoso.com"),
	// 		Sid: to.Ptr("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
	// 		TenantID: to.Ptr("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
	// 	},
	// }
}
