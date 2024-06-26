package armazuredata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azuredata/resource-manager/Microsoft.AzureData/preview/2019-07-24-preview/examples/ListOperation.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuredata.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armazuredata.OperationListResult{
		// 	Value: []*armazuredata.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/sqlServerRegistrations/read"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Retrives details of SQL Server Registration"),
		// 				Operation: to.Ptr("Get SQL Server Registration details"),
		// 				Resource: to.Ptr("SQL Server Registration"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/sqlServerRegistrations/write"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Create a new or change properties of existing SQL Server Registration"),
		// 				Operation: to.Ptr("Create a new or update existing SQL Server Registration"),
		// 				Resource: to.Ptr("SQL Server Registration"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/sqlServerRegistrations/delete"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Delete existing SQL Server Registration"),
		// 				Operation: to.Ptr("Delete existing SQL Server Registration"),
		// 				Resource: to.Ptr("SQL Server Registration"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/operations/read"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Resource: to.Ptr("Available REST operations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/sqlServers/read"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Retrieves details of SQL Server"),
		// 				Operation: to.Ptr("Get SQL Server Instance details"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/sqlServers/write"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Create a new or change properties of existing SQL Server"),
		// 				Operation: to.Ptr("Create new or update existing SQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/sqlServers/delete"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Delete exisiting SQL Server"),
		// 				Operation: to.Ptr("Delete exisiting SQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/sqlInstance/read"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Retrieves details of SQL Instance"),
		// 				Operation: to.Ptr("Get SQL Instance details"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/sqlInstance/write"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Create a new or change properties of existing SQL Instance"),
		// 				Operation: to.Ptr("Create new or update existing SQL Instance"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/sqlInstance/delete"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Delete exisiting SQL Instance"),
		// 				Operation: to.Ptr("Delete exisiting SQL Instance"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/postgresInstances/read"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Retrieves details of Postgres Instances"),
		// 				Operation: to.Ptr("Get Postgres Instance details"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/postgresInstances/write"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Create a new or change properties of existing Postgres Instances"),
		// 				Operation: to.Ptr("Create new or update existing Postgres Instances"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/postgresInstances/delete"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Delete exisiting Postgres Instances"),
		// 				Operation: to.Ptr("Delete exisiting Postgres Instances"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/hybridDataManagers/read"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Retrieves details of HybridDataManagers"),
		// 				Operation: to.Ptr("Get HybridDataManagers details"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/hybridDataManagers/write"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Create a new or change properties of existing HybridDataManagers"),
		// 				Operation: to.Ptr("Create new or update existing HybridDataManagers"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/hybridDataManagers/delete"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Delete exisiting HybridDataManagers"),
		// 				Operation: to.Ptr("Delete exisiting HybridDataManagers"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/SqlBigDataClusters/read"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Retrieves details of SqlBigDataClusters"),
		// 				Operation: to.Ptr("Get SqlBigDataClusters details"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/SqlBigDataClusters/write"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Create a new or change properties of existing SqlBigDataClusters"),
		// 				Operation: to.Ptr("Create new or update existing SqlBigDataClusters"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureData/SqlBigDataClusters/delete"),
		// 			Display: &armazuredata.OperationDisplay{
		// 				Description: to.Ptr("Delete exisiting SqlBigDataClusters"),
		// 				Operation: to.Ptr("Delete exisiting SqlBigDataClusters"),
		// 			},
		// 	}},
		// }
	}
}
