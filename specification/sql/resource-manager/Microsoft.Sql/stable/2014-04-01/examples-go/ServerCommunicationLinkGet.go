package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8358c7473dfe057d84a6b6a921225063c040b31a/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ServerCommunicationLinkGet.json
func ExampleServerCommunicationLinksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerCommunicationLinksClient().Get(ctx, "sqlcrudtest-7398", "sqlcrudtest-4645", "link1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerCommunicationLink = armsql.ServerCommunicationLink{
	// 	Name: to.Ptr("link1"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/communicationLinks"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645/communicationLinks/link1"),
	// 	Location: to.Ptr("Central US"),
	// 	Properties: &armsql.ServerCommunicationLinkProperties{
	// 		PartnerServer: to.Ptr("sqlcrudtest-test"),
	// 		State: to.Ptr("Ready"),
	// 	},
	// }
}
