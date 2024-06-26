package armmanagedservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2022-01-01-preview/examples/GetOperations.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedservices.NewClientFactory(cred, nil)
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
	// res.OperationList = armmanagedservices.OperationList{
	// 	Value: []*armmanagedservices.Operation{
	// 		{
	// 			Name: to.Ptr("Microsoft.ManagedServices/registrationDefinitions/read"),
	// 			Display: &armmanagedservices.OperationDisplay{
	// 				Description: to.Ptr("Retrieves a list of Managed Services registration definitions."),
	// 				Operation: to.Ptr("List Managed Services Registration Definitions"),
	// 				Provider: to.Ptr("Microsoft Managed Services"),
	// 				Resource: to.Ptr("Managed Services Registration Definition"),
	// 			},
	// 	}},
	// }
}
