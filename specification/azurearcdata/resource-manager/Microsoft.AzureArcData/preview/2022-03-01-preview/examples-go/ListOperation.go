package armazurearcdata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/ListOperation.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurearcdata.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armazurearcdata.OperationListResult{
		// 	Value: []*armazurearcdata.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/Locations/OperationStatuses/read"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("read OperationStatuses"),
		// 				Operation: to.Ptr("read_OperationStatuses"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("Locations/OperationStatuses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/Locations/OperationStatuses/write"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("write OperationStatuses"),
		// 				Operation: to.Ptr("write_OperationStatuses"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("Locations/OperationStatuses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/dataControllers/read"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("read dataControllers"),
		// 				Operation: to.Ptr("DataControllers_ListInSubscription"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("dataControllers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/dataControllers/read"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("read dataControllers"),
		// 				Operation: to.Ptr("DataControllers_ListInGroup"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("dataControllers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/dataControllers/read"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Retrieves a dataController resource"),
		// 				Operation: to.Ptr("DataControllers_GetDataController"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("dataControllers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/dataControllers/write"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Creates or replaces a dataController resource"),
		// 				Operation: to.Ptr("DataControllers_PutDataController"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("dataControllers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/dataControllers/delete"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Deletes a dataController resource"),
		// 				Operation: to.Ptr("DataControllers_DeleteDataController"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("dataControllers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/dataControllers/write"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Updates a dataController resource"),
		// 				Operation: to.Ptr("DataControllers_PatchDataController"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("dataControllers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/sqlManagedInstances/read"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("read sqlManagedInstances"),
		// 				Operation: to.Ptr("SqlManagedInstances_List"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("sqlManagedInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/sqlManagedInstances/read"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Gets all sqlManagedInstances in a resource group."),
		// 				Operation: to.Ptr("SqlManagedInstances_ListByResourceGroup"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("sqlManagedInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/sqlManagedInstances/read"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Retrieves a SQL Managed Instance resource"),
		// 				Operation: to.Ptr("SqlManagedInstances_Get"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("sqlManagedInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/sqlManagedInstances/write"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Creates or replaces a SQL Managed Instance resource"),
		// 				Operation: to.Ptr("SqlManagedInstances_Create"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("sqlManagedInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/sqlManagedInstances/delete"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Deletes a SQL Managed Instance resource"),
		// 				Operation: to.Ptr("SqlManagedInstances_Delete"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("sqlManagedInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/sqlManagedInstances/write"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Updates a SQL Managed Instance resource"),
		// 				Operation: to.Ptr("SqlManagedInstances_Update"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("sqlManagedInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/sqlServerInstances/read"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("read sqlServerInstances"),
		// 				Operation: to.Ptr("SqlServerInstances_List"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("sqlServerInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/sqlServerInstances/read"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Gets all sqlServerInstances in a resource group."),
		// 				Operation: to.Ptr("SqlServerInstances_ListByResourceGroup"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("sqlServerInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/sqlServerInstances/read"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Retrieves a SQL Server Instance resource"),
		// 				Operation: to.Ptr("SqlServerInstances_Get"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("sqlServerInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/sqlServerInstances/write"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Creates or replaces a SQL Server Instance resource"),
		// 				Operation: to.Ptr("SqlServerInstances_Create"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("sqlServerInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/sqlServerInstances/delete"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Deletes a SQL Server Instance resource"),
		// 				Operation: to.Ptr("SqlServerInstances_Delete"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("sqlServerInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/sqlServerInstances/write"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Updates a SQL Server Instance resource"),
		// 				Operation: to.Ptr("SqlServerInstances_Update"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("sqlServerInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/dataControllers/activeDirectoryConnectors/read"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("read activeDirectoryConnectors"),
		// 				Operation: to.Ptr("ActiveDirectoryConnectors_List"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("dataControllers/activeDirectoryConnectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/dataControllers/activeDirectoryConnectors/read"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Retrieves an active directory connector resource"),
		// 				Operation: to.Ptr("ActiveDirectoryConnectors_Get"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("dataControllers/activeDirectoryConnectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/dataControllers/activeDirectoryConnectors/write"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Creates or replaces an active directory connector resource."),
		// 				Operation: to.Ptr("ActiveDirectoryConnectors_Create"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("dataControllers/activeDirectoryConnectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/dataControllers/activeDirectoryConnectors/delete"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Deletes an active directory connector resource"),
		// 				Operation: to.Ptr("ActiveDirectoryConnectors_Delete"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("dataControllers/activeDirectoryConnectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/postgresInstances/read"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Retrieves details of Postgres Instances."),
		// 				Operation: to.Ptr("Get Postgres Instance details."),
		// 				Provider: to.Ptr("Azure Arc Data Resource Provider."),
		// 				Resource: to.Ptr("Microsoft.AzureArcData"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/postgresInstances/write"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Create a new or change properties of existing Postgres Instances."),
		// 				Operation: to.Ptr("Create new or update existing Postgres Instances."),
		// 				Provider: to.Ptr("Azure Arc Data Resource Provider."),
		// 				Resource: to.Ptr("Microsoft.AzureArcData"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/postgresInstances/delete"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Delete exisiting Postgres Instances."),
		// 				Operation: to.Ptr("Delete exisiting Postgres Instances."),
		// 				Provider: to.Ptr("Azure Arc Data Resource Provider."),
		// 				Resource: to.Ptr("Microsoft.AzureArcData"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/register/action"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Register the subscription for Microsoft.AzureArcData"),
		// 				Operation: to.Ptr("Register the Microsoft.AzureArcData"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("Microsoft.AzureArcData"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/unregister/action"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("Unregister the subscription for Microsoft.AzureArcData"),
		// 				Operation: to.Ptr("Unregister the Microsoft.AzureArcData"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("Microsoft.AzureArcData"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureArcData/Operations/read"),
		// 			Display: &armazurearcdata.OperationDisplay{
		// 				Description: to.Ptr("read Operations"),
		// 				Operation: to.Ptr("read_Operations"),
		// 				Provider: to.Ptr("Microsoft.AzureArcData"),
		// 				Resource: to.Ptr("Operations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
