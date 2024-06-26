package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7e29dd59eef13ef347d09e41a63f2585be77b3ca/specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/ListSupportedPrivateLinkResources.json
func ExamplePrivateLinkResourcesClient_NewListSupportedPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListSupportedPager("rg1", "mysearchservice", &armsearch.SearchManagementRequestOptions{ClientRequestID: nil}, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PrivateLinkResourcesResult = armsearch.PrivateLinkResourcesResult{
		// 	Value: []*armsearch.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("searchService"),
		// 			Type: to.Ptr("Microsoft.Search/searchServices/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Search/searchServices/mysearchservice/privateLinkResources/searchService"),
		// 			Properties: &armsearch.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("searchService"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("searchService")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.search.windows.net")},
		// 						ShareablePrivateLinkResourceTypes: []*armsearch.ShareablePrivateLinkResourceType{
		// 							{
		// 								Name: to.Ptr("blob"),
		// 								Properties: &armsearch.ShareablePrivateLinkResourceProperties{
		// 									Type: to.Ptr("Microsoft.Storage/storageAccounts"),
		// 									Description: to.Ptr("Azure Cognitive Search indexers can connect to blobs in Azure Storage for reading data (data source), for writing intermediate results of indexer execution (annotation cache, preview) or for storing any knowledge store projections (preview)"),
		// 									GroupID: to.Ptr("blob"),
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("table"),
		// 								Properties: &armsearch.ShareablePrivateLinkResourceProperties{
		// 									Type: to.Ptr("Microsoft.Storage/storageAccounts"),
		// 									Description: to.Ptr("Azure Cognitive Search indexers can connect to tables in Azure Storage for reading data (data source), for writing book-keeping information about intermediate results of indexer execution (annotation cache, preview) or for storing any knowledge store projections (preview)"),
		// 									GroupID: to.Ptr("table"),
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("Sql"),
		// 								Properties: &armsearch.ShareablePrivateLinkResourceProperties{
		// 									Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts"),
		// 									Description: to.Ptr("Azure Cognitive Search indexers can connect to CosmosDB using the SQL head for reading data (data source)."),
		// 									GroupID: to.Ptr("Sql"),
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("plr"),
		// 								Properties: &armsearch.ShareablePrivateLinkResourceProperties{
		// 									Type: to.Ptr("Microsoft.Sql/servers"),
		// 									Description: to.Ptr("Azure Cognitive Search indexers can connect to AzureSQL databases in a SQL server for reading data (data source)."),
		// 									GroupID: to.Ptr("sqlServer"),
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("vault"),
		// 								Properties: &armsearch.ShareablePrivateLinkResourceProperties{
		// 									Type: to.Ptr("Microsoft.KeyVault/vaults"),
		// 									Description: to.Ptr("Azure Cognitive Search can access keys in Azure Key Vault to encrypt search index and synonym map data"),
		// 									GroupID: to.Ptr("vault"),
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("plr"),
		// 								Properties: &armsearch.ShareablePrivateLinkResourceProperties{
		// 									Type: to.Ptr("Microsoft.DBforMySQL/servers"),
		// 									Description: to.Ptr("Azure Cognitive Search indexers can connect to MySQL databases for reading data (data source, preview)."),
		// 									GroupID: to.Ptr("mysqlServer"),
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("site"),
		// 								Properties: &armsearch.ShareablePrivateLinkResourceProperties{
		// 									Type: to.Ptr("Microsoft.Web/sites"),
		// 									Description: to.Ptr("Azure Cognitive Search indexers can connect to App Services when executing custom web api skills that can be present in a skillset (optional) attached to the indexer."),
		// 									GroupID: to.Ptr("sites"),
		// 								},
		// 						}},
		// 					},
		// 			}},
		// 		}
	}
}
