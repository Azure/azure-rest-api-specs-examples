package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/recipientTransfersValidate.json
func ExampleRecipientTransfersClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecipientTransfersClient().Validate(ctx, "aabb123", armbilling.AcceptTransferRequest{
		Properties: &armbilling.AcceptTransferProperties{
			ProductDetails: []*armbilling.ProductDetails{
				{
					ProductID:   to.Ptr("subscriptionId"),
					ProductType: to.Ptr(armbilling.ProductTypeAzureSubscription),
				},
				{
					ProductID:   to.Ptr("reservedInstanceId"),
					ProductType: to.Ptr(armbilling.ProductTypeAzureReservation),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ValidateTransferListResponse = armbilling.ValidateTransferListResponse{
	// 	Value: []*armbilling.ValidateTransferResponse{
	// 		{
	// 			Properties: &armbilling.ValidateTransferResponseProperties{
	// 				ProductID: to.Ptr("subscriptionId"),
	// 				Results: []*armbilling.ValidationResultProperties{
	// 					{
	// 						Code: to.Ptr("NotIntendedRecipient"),
	// 						Level: to.Ptr("Error"),
	// 						Message: to.Ptr("Intended recipient is different."),
	// 				}},
	// 				Status: to.Ptr("Failed"),
	// 			},
	// 	}},
	// }
}
