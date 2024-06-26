package armdatacatalog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datacatalog/armdatacatalog"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datacatalog/resource-manager/Microsoft.DataCatalog/stable/2016-03-30/examples/GetOperations.json
func ExampleADCOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatacatalog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewADCOperationsClient().List(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationEntityListResult = armdatacatalog.OperationEntityListResult{
	// 	Value: []*armdatacatalog.OperationEntity{
	// 		{
	// 			Name: to.Ptr("Microsoft.DataCatalog/catalogs/read"),
	// 			Display: &armdatacatalog.OperationDisplayInfo{
	// 				Description: to.Ptr("Get properties for catalog or catalogs under subscription or resource group."),
	// 				Operation: to.Ptr("Catalog Read Or List"),
	// 				Provider: to.Ptr("Microsoft Data Catalog"),
	// 				Resource: to.Ptr("Catalog"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.DataCatalog/catalogs/write"),
	// 			Display: &armdatacatalog.OperationDisplayInfo{
	// 				Description: to.Ptr("Creates catalog or updates the tags and properties for the catalog."),
	// 				Operation: to.Ptr("Create Or Update Catalog"),
	// 				Provider: to.Ptr("Microsoft Data Catalog"),
	// 				Resource: to.Ptr("Catalog"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.DataCatalog/catalogs/delete"),
	// 			Display: &armdatacatalog.OperationDisplayInfo{
	// 				Description: to.Ptr("Deletes the catalog."),
	// 				Operation: to.Ptr("Delete Catalog"),
	// 				Provider: to.Ptr("Microsoft Data Catalog"),
	// 				Resource: to.Ptr("Catalog"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.DataCatalog/operations/read"),
	// 			Display: &armdatacatalog.OperationDisplayInfo{
	// 				Description: to.Ptr("Lists operations available on Microsoft.DataCatalog resource provider."),
	// 				Operation: to.Ptr("List Available Catalog Operations"),
	// 				Provider: to.Ptr("Microsoft Data Catalog"),
	// 				Resource: to.Ptr("Available Catalog Operations"),
	// 			},
	// 	}},
	// }
}
