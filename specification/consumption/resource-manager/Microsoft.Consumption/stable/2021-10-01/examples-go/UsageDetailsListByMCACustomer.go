package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/UsageDetailsListByMCACustomer.json
func ExampleUsageDetailsClient_NewListPager_customerUsageDetailsListModern() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsageDetailsClient().NewListPager("providers/Microsoft.Billing/BillingAccounts/1234:56789/customers/00000000-0000-0000-0000-000000000000", &armconsumption.UsageDetailsClientListOptions{Expand: nil,
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
		// 		&armconsumption.ModernUsageDetail{
		// 			Name: to.Ptr("usageDetails_Id1"),
		// 			Type: to.Ptr("Microsoft.Consumption/usageDetails"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/1234:56789/customers/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/usageDetails/usageDetails_Id1"),
		// 			Tags: map[string]*string{
		// 				"dev": to.Ptr("tools"),
		// 				"env": to.Ptr("newcrp"),
		// 			},
		// 			Kind: to.Ptr(armconsumption.UsageDetailsKindModern),
		// 			Properties: &armconsumption.ModernUsageDetailProperties{
		// 				AdditionalInfo: to.Ptr("{  \"UsageType\": \"ComputeHR\",  \"ImageType\": \"Windows Client BYOL\",  \"ServiceType\": \"Standard_D1\",  \"VMName\": null,  \"VMProperties\": null,  \"VCPUs\": 1,  \"CPUs\": 0}"),
		// 				BenefitID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				BenefitName: to.Ptr("Reservation_purchase_03-09-2018_10-59"),
		// 				BillingAccountID: to.Ptr("1234:56789"),
		// 				BillingAccountName: to.Ptr("Account Name 1"),
		// 				BillingCurrencyCode: to.Ptr("USD"),
		// 				BillingPeriodEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-31T00:00:00.000Z"); return t}()),
		// 				BillingPeriodStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-01T00:00:00.000Z"); return t}()),
		// 				BillingProfileID: to.Ptr("2468"),
		// 				BillingProfileName: to.Ptr("Account Name 1"),
		// 				ChargeType: to.Ptr("Usage"),
		// 				ConsumedService: to.Ptr("Microsoft.Storage"),
		// 				CostCenter: to.Ptr("DEV"),
		// 				CostInBillingCurrency: to.Ptr[float64](1.84763819095477),
		// 				CostInPricingCurrency: to.Ptr[float64](1.84763819095477),
		// 				CostInUSD: to.Ptr[float64](1.84763819095477),
		// 				CustomerName: to.Ptr("Modern Azure Customer 1"),
		// 				CustomerTenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				Date: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-30T00:00:00.000Z"); return t}()),
		// 				ExchangeRate: to.Ptr("1"),
		// 				ExchangeRateDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-01T00:00:00.000Z"); return t}()),
		// 				ExchangeRatePricingToBilling: to.Ptr[float64](0.077),
		// 				Frequency: to.Ptr("UsageBased"),
		// 				InstanceName: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/Resource Group 1/providers/Microsoft.Storage/storageAccounts/Resource Name 1"),
		// 				InvoiceID: to.Ptr(""),
		// 				InvoiceSectionID: to.Ptr("98765"),
		// 				InvoiceSectionName: to.Ptr("Invoice Section 1"),
		// 				IsAzureCreditEligible: to.Ptr(false),
		// 				MarketPrice: to.Ptr[float64](0.077),
		// 				MeterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				PartnerEarnedCreditApplied: to.Ptr("0"),
		// 				PartnerEarnedCreditRate: to.Ptr[float64](0.077),
		// 				PartnerName: to.Ptr("Partner Name 1"),
		// 				PartnerTenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				PaygCostInBillingCurrency: to.Ptr[float64](1.848),
		// 				PaygCostInUSD: to.Ptr[float64](1.848),
		// 				PreviousInvoiceID: to.Ptr(""),
		// 				PricingCurrencyCode: to.Ptr("USD"),
		// 				Product: to.Ptr("Virtual Machines D Series - D1 - US East"),
		// 				ProductIdentifier: to.Ptr("DZH318Z0BQ4B00FV"),
		// 				ProductOrderID: to.Ptr("a3db7880-70eb-4b4c-6a79-1425a058df5a"),
		// 				ProductOrderName: to.Ptr("Azure plan"),
		// 				PublisherID: to.Ptr(""),
		// 				PublisherName: to.Ptr("Microsoft"),
		// 				PublisherType: to.Ptr("Microsoft"),
		// 				Quantity: to.Ptr[float64](0.7329),
		// 				ResellerMpnID: to.Ptr(""),
		// 				ResellerName: to.Ptr("Reseller Name 1"),
		// 				ReservationID: to.Ptr(""),
		// 				ReservationName: to.Ptr(""),
		// 				ResourceGroup: to.Ptr("Resource Group 1"),
		// 				ResourceLocation: to.Ptr("USEast"),
		// 				ResourceLocationNormalized: to.Ptr("US East"),
		// 				ServiceInfo1: to.Ptr(""),
		// 				ServiceInfo2: to.Ptr("Windows Client BYOL"),
		// 				ServicePeriodEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-01T00:00:00.000Z"); return t}()),
		// 				ServicePeriodStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-01T00:00:00.000Z"); return t}()),
		// 				SubscriptionGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				SubscriptionName: to.Ptr("Subscription Name 1"),
		// 				Term: to.Ptr(""),
		// 				UnitPrice: to.Ptr[float64](4.38),
		// 			},
		// 	}},
		// }
	}
}
