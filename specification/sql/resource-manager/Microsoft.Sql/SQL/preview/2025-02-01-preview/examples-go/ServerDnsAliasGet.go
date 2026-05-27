package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/ServerDnsAliasGet.json
func ExampleServerDNSAliasesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerDNSAliasesClient().Get(ctx, "Default", "dns-alias-server", "dns-alias-name-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.ServerDNSAliasesClientGetResponse{
	// 	ServerDNSAlias: armsql.ServerDNSAlias{
	// 		Name: to.Ptr("dns-alias-name-1"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/dnsAliases"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/dns-alias-server/dnsAliases/dns-alias-name-1"),
	// 		Properties: &armsql.ServerDNSAliasProperties{
	// 			AzureDNSRecord: to.Ptr("dns-alias-name-1.database.windows.net"),
	// 		},
	// 	},
	// }
}
