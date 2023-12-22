package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ManagedServerDnsAliasGet.json
func ExampleManagedServerDNSAliasesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedServerDNSAliasesClient().Get(ctx, "Default", "dns-mi", "dns-alias-mi", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedServerDNSAlias = armsql.ManagedServerDNSAlias{
	// 	Name: to.Ptr("dns-alias-mi"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/dnsAliases"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/dns-mi/dnsAliases/dns-alias-mi"),
	// 	Properties: &armsql.ManagedServerDNSAliasProperties{
	// 		AzureDNSRecord: to.Ptr("dns-alias-mi.abcd1234.database.windows.net"),
	// 		PublicAzureDNSRecord: to.Ptr("dns-alias-mi.public.abcd1234.database.windows.net"),
	// 	},
	// }
}
