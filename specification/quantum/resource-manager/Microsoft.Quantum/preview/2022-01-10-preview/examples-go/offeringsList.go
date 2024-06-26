package armquantum_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quantum/armquantum"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/quantum/resource-manager/Microsoft.Quantum/preview/2022-01-10-preview/examples/offeringsList.json
func ExampleOfferingsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquantum.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOfferingsClient().NewListPager("westus2", nil)
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
		// page.OfferingsListResult = armquantum.OfferingsListResult{
		// 	Value: []*armquantum.ProviderDescription{
		// 		{
		// 			Name: to.Ptr("Quantum Inspired Optimization"),
		// 			ID: to.Ptr("QIO"),
		// 			Properties: &armquantum.ProviderProperties{
		// 				Description: to.Ptr("Provide Quantum Inspired Optimization solvers"),
		// 				PricingDimensions: []*armquantum.PricingDimension{
		// 					{
		// 						Name: to.Ptr("Pricing"),
		// 						ID: to.Ptr("price"),
		// 					},
		// 					{
		// 						Name: to.Ptr("Minimum monthly rate"),
		// 						ID: to.Ptr("rate"),
		// 				}},
		// 				QuotaDimensions: []*armquantum.QuotaDimension{
		// 					{
		// 						Name: to.Ptr("Job count"),
		// 						Description: to.Ptr("The number of jobs you may submit per month"),
		// 						ID: to.Ptr("job-count"),
		// 						Period: to.Ptr("Monthly"),
		// 						Quota: to.Ptr[float32](30),
		// 						Scope: to.Ptr("Workspace"),
		// 						Unit: to.Ptr("job"),
		// 						UnitPlural: to.Ptr("jobs"),
		// 					},
		// 					{
		// 						Name: to.Ptr("Job hours"),
		// 						Description: to.Ptr("The number of hours of solver time you may use per month"),
		// 						ID: to.Ptr("job-hour"),
		// 						Period: to.Ptr("Monthly"),
		// 						Quota: to.Ptr[float32](1000),
		// 						Scope: to.Ptr("Subscription"),
		// 						Unit: to.Ptr("hour"),
		// 						UnitPlural: to.Ptr("hours"),
		// 				}},
		// 				SKUs: []*armquantum.SKUDescription{
		// 					{
		// 						Name: to.Ptr("Standard"),
		// 						Description: to.Ptr("Provider CPU and FPGA QIO based solver."),
		// 						AutoAdd: to.Ptr(true),
		// 						ID: to.Ptr("Standard"),
		// 						PricingDetails: []*armquantum.PricingDetail{
		// 							{
		// 								ID: to.Ptr("price"),
		// 								Value: to.Ptr("200 dollars"),
		// 							},
		// 							{
		// 								ID: to.Ptr("rate"),
		// 								Value: to.Ptr("10 hours/month"),
		// 						}},
		// 						QuotaDimensions: []*armquantum.QuotaDimension{
		// 							{
		// 								ID: to.Ptr("quota1"),
		// 								Scope: to.Ptr("Workspace"),
		// 						}},
		// 						RestrictedAccessURI: to.Ptr("https://endpoint"),
		// 						Targets: []*string{
		// 							to.Ptr("p1")},
		// 							Version: to.Ptr("1.0"),
		// 					}},
		// 					Targets: []*armquantum.TargetDescription{
		// 						{
		// 							Name: to.Ptr("CPU annealer"),
		// 							Description: to.Ptr("CPU annealer algorithm"),
		// 							ID: to.Ptr("p1"),
		// 					}},
		// 				},
		// 		}},
		// 	}
	}
}
