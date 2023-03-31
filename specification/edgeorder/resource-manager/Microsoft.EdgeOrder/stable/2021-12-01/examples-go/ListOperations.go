package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListOperations.json
func ExampleManagementClient_NewListOperationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeorder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagementClient().NewListOperationsPager(nil)
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
		// page.OperationListResult = armedgeorder.OperationListResult{
		// 	Value: []*armedgeorder.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/addresses/read"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("List or get the Addresses"),
		// 				Operation: to.Ptr("List or Get Addresses"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("Addresses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/addresses/delete"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("Delete the Addresses"),
		// 				Operation: to.Ptr("Delete Addresses"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("Addresses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/addresses/write"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("Create or update the Addresses"),
		// 				Operation: to.Ptr("Create or Update Addresses"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("Addresses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/locations/operationResults/read"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("List or get the Operation Results"),
		// 				Operation: to.Ptr("List or Get Operation Results"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("Operation Results"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/operations/read"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("List or get the Operations"),
		// 				Operation: to.Ptr("List or Get Operations"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("Operations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/locations/orders/read"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("List or get the Order"),
		// 				Operation: to.Ptr("List or Get Order"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("Order"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/orders/read"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("List or get the Order"),
		// 				Operation: to.Ptr("List or Get Order"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("Order"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/orderItems/cancel/action"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("Cancels an OrderItem in progress."),
		// 				Operation: to.Ptr("Cancel OrderItem"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("OrderItem"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/orderItems/return/action"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("Return an OrderItem."),
		// 				Operation: to.Ptr("Return OrderItem"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("OrderItem"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/orderItems/read"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("List or get the OrderItem"),
		// 				Operation: to.Ptr("List or Get OrderItem"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("OrderItem"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/orderItems/delete"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("Delete the OrderItem"),
		// 				Operation: to.Ptr("Delete OrderItem"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("OrderItem"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/orderItems/write"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("Create or update the OrderItem"),
		// 				Operation: to.Ptr("Create or Update OrderItem"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("OrderItem"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/productFamiliesMetadata/action"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("This method lists or gets the product families metadata."),
		// 				Operation: to.Ptr("List or Get product families metadata"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("ArmApiRes_Microsoft.EdgeOrder"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/listProductFamilies/read"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("This method returns list of product families."),
		// 				Operation: to.Ptr("List Product Families"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("ArmApiRes_Microsoft.EdgeOrder"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrder/listConfigurations/action"),
		// 			Display: &armedgeorder.OperationDisplay{
		// 				Description: to.Ptr("This method returns list of product configurations."),
		// 				Operation: to.Ptr("List Product Configurations"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("ArmApiRes_Microsoft.EdgeOrder"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorder.OriginUser),
		// 	}},
		// }
	}
}
