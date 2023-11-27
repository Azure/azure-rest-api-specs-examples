package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/UsageDetailsList.json
func ExampleUsageDetailsClient_NewListPager_usageDetailsListLegacy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsageDetailsClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000", &armconsumption.UsageDetailsClientListOptions{Expand: nil,
		Filter:    nil,
		Skiptoken: nil,
		Top:       nil,
		Metric:    nil,
	})
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
		// page.UsageDetailsListResult = armconsumption.UsageDetailsListResult{
		// 	Value: []armconsumption.UsageDetailClassification{
		// 		&armconsumption.LegacyUsageDetail{
		// 			Name: to.Ptr("usageDetails_Id1"),
		// 			Type: to.Ptr("Microsoft.Consumption/usageDetails"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Billing/billingPeriods/201903/providers/Microsoft.Consumption/usageDetails/usageDetails_Id1"),
		// 			Tags: map[string]*string{
		// 				"dev": to.Ptr("tools"),
		// 				"env": to.Ptr("newcrp"),
		// 			},
		// 			Kind: to.Ptr(armconsumption.UsageDetailsKindLegacy),
		// 			Properties: &armconsumption.LegacyUsageDetailProperties{
		// 				AccountName: to.Ptr("AccountName"),
		// 				BenefitID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				BenefitName: to.Ptr("Reservation_purchase_03-09-2018_10-59"),
		// 				BillingAccountID: to.Ptr("xxxxxxxx"),
		// 				BillingAccountName: to.Ptr("Customer Name 1"),
		// 				BillingCurrency: to.Ptr("CAD"),
		// 				BillingPeriodEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-30T00:00:00.000Z"); return t}()),
		// 				BillingPeriodStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-01T00:00:00.000Z"); return t}()),
		// 				BillingProfileID: to.Ptr("xxxxxxxx"),
		// 				BillingProfileName: to.Ptr("Customer Name 1"),
		// 				ChargeType: to.Ptr("Usage"),
		// 				ConsumedService: to.Ptr("Microsoft.ClassicStorage"),
		// 				Cost: to.Ptr[float64](0.000001980949998),
		// 				CostCenter: to.Ptr("BAS"),
		// 				Date: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-09T00:00:00.000Z"); return t}()),
		// 				EffectivePrice: to.Ptr[float64](0.054693034210767),
		// 				Frequency: to.Ptr("UsageBased"),
		// 				InvoiceSection: to.Ptr("Invoice Section 1"),
		// 				IsAzureCreditEligible: to.Ptr(false),
		// 				MeterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				OfferID: to.Ptr("Offer Id 1"),
		// 				PartNumber: to.Ptr("Part Number 1"),
		// 				Product: to.Ptr("Product1"),
		// 				Quantity: to.Ptr[float64](0.000036),
		// 				ResourceGroup: to.Ptr("Resource-Group-westus"),
		// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/Resource-Group-westus/providers/Microsoft.ClassicStorage/storageAccounts/ResourceName1"),
		// 				ResourceLocation: to.Ptr("uswest"),
		// 				ResourceName: to.Ptr("ResourceName1"),
		// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				SubscriptionName: to.Ptr("SubscriptionName 1"),
		// 				UnitPrice: to.Ptr[float64](5.47),
		// 			},
		// 	}},
		// }
	}
}
