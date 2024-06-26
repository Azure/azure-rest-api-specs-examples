package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/OperationsGet.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
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
		// page.OperationListResult = armbillingbenefits.OperationListResult{
		// 	Value: []*armbillingbenefits.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrderAliases/read"),
		// 			Display: &armbillingbenefits.OperationDisplay{
		// 				Description: to.Ptr("Read all savings plan order aliases"),
		// 				Operation: to.Ptr("Get Savings plan order alias"),
		// 				Provider: to.Ptr("Microsoft Benefits"),
		// 				Resource: to.Ptr("SavingsPlanOrderAliases"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrderAliases/write"),
		// 			Display: &armbillingbenefits.OperationDisplay{
		// 				Description: to.Ptr("Create a Savings plan order alias"),
		// 				Operation: to.Ptr("Create SavingsPlanOrderAliases"),
		// 				Provider: to.Ptr("Microsoft Benefits"),
		// 				Resource: to.Ptr("SavingsPlanOrderAliases"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrders/read"),
		// 			Display: &armbillingbenefits.OperationDisplay{
		// 				Description: to.Ptr("Read all savings plan orders"),
		// 				Operation: to.Ptr("Get Savings plan orders"),
		// 				Provider: to.Ptr("Microsoft Benefits"),
		// 				Resource: to.Ptr("SavingsPlanOrders"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrders/write"),
		// 			Display: &armbillingbenefits.OperationDisplay{
		// 				Description: to.Ptr("Patch a Savings plan order"),
		// 				Operation: to.Ptr("Patch SavingsPlanOrders"),
		// 				Provider: to.Ptr("Microsoft Benefits"),
		// 				Resource: to.Ptr("SavingsPlanOrders"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrders/action"),
		// 			Display: &armbillingbenefits.OperationDisplay{
		// 				Description: to.Ptr("Update a Savings plan order"),
		// 				Operation: to.Ptr("Update SavingsPlanOrders"),
		// 				Provider: to.Ptr("Microsoft Benefits"),
		// 				Resource: to.Ptr("SavingsPlanOrders"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrders/savingsPlans/read"),
		// 			Display: &armbillingbenefits.OperationDisplay{
		// 				Description: to.Ptr("Read All SavingsPlans"),
		// 				Operation: to.Ptr("Get SavingsPlans"),
		// 				Provider: to.Ptr("Microsoft Benefits"),
		// 				Resource: to.Ptr("SavingsPlans"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrders/savingsPlans/write"),
		// 			Display: &armbillingbenefits.OperationDisplay{
		// 				Description: to.Ptr("Patch an existing Savings plan"),
		// 				Operation: to.Ptr("Patch SavingsPlans"),
		// 				Provider: to.Ptr("Microsoft Benefits"),
		// 				Resource: to.Ptr("SavingsPlans"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.BillingBenefits/register/action"),
		// 			Display: &armbillingbenefits.OperationDisplay{
		// 				Description: to.Ptr("Registers the Benefits resource provider and enables the creation of Benefits resources."),
		// 				Operation: to.Ptr("Registers the Benefits Resource Provider."),
		// 				Provider: to.Ptr("Microsoft Benefits"),
		// 				Resource: to.Ptr("SavingsPlans"),
		// 			},
		// 	}},
		// }
	}
}
