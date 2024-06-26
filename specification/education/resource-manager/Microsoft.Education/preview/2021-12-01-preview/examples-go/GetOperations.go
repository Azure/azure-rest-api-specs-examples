package armeducation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b9b91929c304f8fb44002267b6c98d9fb9dde014/specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/GetOperations.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeducation.NewClientFactory(cred, nil)
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
	// res.OperationListResult = armeducation.OperationListResult{
	// 	Value: []*armeducation.Operation{
	// 		{
	// 			Name: to.Ptr("Microsoft.Education/grants/read"),
	// 			Display: &armeducation.OperationDisplay{
	// 				Description: to.Ptr("Read all grants"),
	// 				Operation: to.Ptr("Get grants"),
	// 				Provider: to.Ptr("Microsoft Education"),
	// 				Resource: to.Ptr("Grants"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Education/labs/read"),
	// 			Display: &armeducation.OperationDisplay{
	// 				Description: to.Ptr("Read all labs"),
	// 				Operation: to.Ptr("Get labs"),
	// 				Provider: to.Ptr("Microsoft Edcucation"),
	// 				Resource: to.Ptr("Labs"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Education/labs/write"),
	// 			Display: &armeducation.OperationDisplay{
	// 				Description: to.Ptr("Create or update lab"),
	// 				Operation: to.Ptr("Create or update lab"),
	// 				Provider: to.Ptr("Microsoft Edcucation"),
	// 				Resource: to.Ptr("Labs"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Education/labs/delete"),
	// 			Display: &armeducation.OperationDisplay{
	// 				Description: to.Ptr("Delete a lab"),
	// 				Operation: to.Ptr("Delete lab"),
	// 				Provider: to.Ptr("Microsoft Edcucation"),
	// 				Resource: to.Ptr("Labs"),
	// 			},
	// 	}},
	// }
}
