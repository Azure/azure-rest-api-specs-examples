package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingProfilesValidateDeleteEligibilityFailure.json
func ExampleProfilesClient_ValidateDeleteEligibility_billingProfilesValidateDeleteEligibilityFailure() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProfilesClient().ValidateDeleteEligibility(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeleteBillingProfileEligibilityResult = armbilling.DeleteBillingProfileEligibilityResult{
	// 	EligibilityDetails: []*armbilling.DeleteBillingProfileEligibilityDetail{
	// 		{
	// 			Code: to.Ptr(armbilling.DeleteBillingProfileEligibilityCodeActiveBillingSubscriptions),
	// 			Message: to.Ptr("There are active or disabled subscriptions assigned to the invoice section. Either move the subscription to another invoice section or delete the subscriptions and then try deleting the invoice section."),
	// 		},
	// 		{
	// 			Code: to.Ptr(armbilling.DeleteBillingProfileEligibilityCodeLastBillingProfile),
	// 			Message: to.Ptr("Billing profile cannot be deleted as this is the only billing profile in this billing account."),
	// 		},
	// 		{
	// 			Code: to.Ptr(armbilling.DeleteBillingProfileEligibilityCodeOutstandingCharges),
	// 			Message: to.Ptr("Billing Profile cannot be deleted as there are outstanding charges on this billing profile."),
	// 		},
	// 		{
	// 			Code: to.Ptr(armbilling.DeleteBillingProfileEligibilityCodePendingCharges),
	// 			Message: to.Ptr("Billing Profile cannot be deleted as there are pending charges accumulating on this billing profile."),
	// 		},
	// 		{
	// 			Code: to.Ptr(armbilling.DeleteBillingProfileEligibilityCodeReservedInstances),
	// 			Message: to.Ptr("Billing Profile cannot be deleted as there are reserved assets with a billing plan."),
	// 	}},
	// 	EligibilityStatus: to.Ptr(armbilling.DeleteBillingProfileEligibilityStatusNotAllowed),
	// }
}
