package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ValidateProductMoveFailure.json
func ExampleProductsClient_ValidateMove_subscriptionMoveValidateFailure() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProductsClient().ValidateMove(ctx, "{billingAccountName}", "{productName}", armbilling.TransferProductRequestProperties{
		DestinationInvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{newInvoiceSectionName}"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ValidateProductTransferEligibilityResult = armbilling.ValidateProductTransferEligibilityResult{
	// 	ErrorDetails: &armbilling.ValidateProductTransferEligibilityError{
	// 		Code: to.Ptr(armbilling.ProductTransferValidationErrorCodeProductTypeNotSupported),
	// 		Message: to.Ptr("Product '{productName}' is not allowed to be transferred."),
	// 	},
	// 	IsMoveEligible: to.Ptr(false),
	// }
}
