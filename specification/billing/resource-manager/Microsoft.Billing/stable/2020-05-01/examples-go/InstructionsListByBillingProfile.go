package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InstructionsListByBillingProfile.json
func ExampleInstructionsClient_NewListByBillingProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInstructionsClient().NewListByBillingProfilePager("{billingAccountName}", "{billingProfileName}", nil)
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
		// page.InstructionListResult = armbilling.InstructionListResult{
		// 	Value: []*armbilling.Instruction{
		// 		{
		// 			Name: to.Ptr("TO1:CLIN001"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/instructions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/instructions/TO1:CLIN001"),
		// 			Properties: &armbilling.InstructionProperties{
		// 				Amount: to.Ptr[float32](5000),
		// 				EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-30T21:26:47.997Z"); return t}()),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-30T21:26:47.997Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("TO1:CLIN002"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/instructions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/instructions/TO1:CLIN002"),
		// 			Properties: &armbilling.InstructionProperties{
		// 				Amount: to.Ptr[float32](2000),
		// 				EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-30T21:26:47.997Z"); return t}()),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-30T21:26:47.997Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
