package armdatalakestore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Operations_List.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakestore.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().List(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationListResult = armdatalakestore.OperationListResult{
	// 	Value: []*armdatalakestore.Operation{
	// 		{
	// 			Name: to.Ptr("Microsoft.DataLakeStore/operations/read"),
	// 			Display: &armdatalakestore.OperationDisplay{
	// 				Description: to.Ptr("Get available operations of DataLakeStore."),
	// 				Operation: to.Ptr("Get Available Operations"),
	// 				Provider: to.Ptr("Microsoft DataLakeStore"),
	// 				Resource: to.Ptr("Available Operations"),
	// 			},
	// 	}},
	// }
}
