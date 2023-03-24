package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetCatalog.json
func ExampleAzureReservationAPIClient_NewGetCatalogPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armreservations.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAzureReservationAPIClient().NewGetCatalogPager("23bc208b-083f-4901-ae85-4f98c0c3b4b6", &armreservations.AzureReservationAPIClientGetCatalogOptions{ReservedResourceType: to.Ptr("VirtualMachines"),
		Location:    to.Ptr("eastus"),
		PublisherID: nil,
		OfferID:     nil,
		PlanID:      nil,
		Filter:      nil,
		Skip:        nil,
		Take:        nil,
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
		// page.CatalogsResult = armreservations.CatalogsResult{
		// 	TotalItems: to.Ptr[int64](1000),
		// 	Value: []*armreservations.Catalog{
		// 		{
		// 			Name: to.Ptr("Standard_DS5_v2"),
		// 			BillingPlans: map[string][]*armreservations.ReservationBillingPlan{
		// 				"P1Y": []*armreservations.ReservationBillingPlan{
		// 					to.Ptr(armreservations.ReservationBillingPlanUpfront),
		// 					to.Ptr(armreservations.ReservationBillingPlanMonthly)},
		// 					"P3Y": []*armreservations.ReservationBillingPlan{
		// 						to.Ptr(armreservations.ReservationBillingPlanUpfront),
		// 						to.Ptr(armreservations.ReservationBillingPlanMonthly)},
		// 					},
		// 					Locations: []*string{
		// 						to.Ptr("eastus")},
		// 						ResourceType: to.Ptr("VirtualMachines"),
		// 						Restrictions: []*armreservations.SKURestriction{
		// 						},
		// 						SKUProperties: []*armreservations.SKUProperty{
		// 							{
		// 								Name: to.Ptr("Cores"),
		// 								Value: to.Ptr("16"),
		// 							},
		// 							{
		// 								Name: to.Ptr("ProductTitle"),
		// 								Value: to.Ptr("DSv2 Series, DS5"),
		// 							},
		// 							{
		// 								Name: to.Ptr("ProductShortName"),
		// 								Value: to.Ptr("DSv2 Series"),
		// 							},
		// 							{
		// 								Name: to.Ptr("SKUName"),
		// 								Value: to.Ptr("DS5 v2"),
		// 							},
		// 							{
		// 								Name: to.Ptr("MeterId"),
		// 								Value: to.Ptr("12bc208b-083f-4901-ae85-4f98c0c3b4b8"),
		// 						}},
		// 						Terms: []*armreservations.ReservationTerm{
		// 							to.Ptr(armreservations.ReservationTermP1Y),
		// 							to.Ptr(armreservations.ReservationTermP3Y)},
		// 						},
		// 						{
		// 							Name: to.Ptr("Standard_D1"),
		// 							BillingPlans: map[string][]*armreservations.ReservationBillingPlan{
		// 								"P1Y": []*armreservations.ReservationBillingPlan{
		// 									to.Ptr(armreservations.ReservationBillingPlanUpfront),
		// 									to.Ptr(armreservations.ReservationBillingPlanMonthly)},
		// 									"P3Y": []*armreservations.ReservationBillingPlan{
		// 										to.Ptr(armreservations.ReservationBillingPlanUpfront),
		// 										to.Ptr(armreservations.ReservationBillingPlanMonthly)},
		// 									},
		// 									Locations: []*string{
		// 										to.Ptr("eastus")},
		// 										ResourceType: to.Ptr("VirtualMachines"),
		// 										Restrictions: []*armreservations.SKURestriction{
		// 											{
		// 												Type: to.Ptr("Term"),
		// 												ReasonCode: to.Ptr("NotAvailableForSubscription"),
		// 												Values: []*string{
		// 													to.Ptr("P1Y")},
		// 												},
		// 												{
		// 													Type: to.Ptr("Term"),
		// 													ReasonCode: to.Ptr("NotAvailableForSubscription"),
		// 													Values: []*string{
		// 														to.Ptr("P3Y")},
		// 												}},
		// 												SKUProperties: []*armreservations.SKUProperty{
		// 													{
		// 														Name: to.Ptr("Cores"),
		// 														Value: to.Ptr("1"),
		// 													},
		// 													{
		// 														Name: to.Ptr("ProductTitle"),
		// 														Value: to.Ptr("D Series, D1"),
		// 													},
		// 													{
		// 														Name: to.Ptr("ProductShortName"),
		// 														Value: to.Ptr("D Series"),
		// 													},
		// 													{
		// 														Name: to.Ptr("SKUName"),
		// 														Value: to.Ptr("D1"),
		// 													},
		// 													{
		// 														Name: to.Ptr("MeterId"),
		// 														Value: to.Ptr("12bc208b-083f-4901-ae85-4f98c0c3b4b8"),
		// 												}},
		// 												Terms: []*armreservations.ReservationTerm{
		// 													to.Ptr(armreservations.ReservationTermP1Y),
		// 													to.Ptr(armreservations.ReservationTermP3Y)},
		// 												},
		// 												{
		// 													Name: to.Ptr("Standard_F2"),
		// 													BillingPlans: map[string][]*armreservations.ReservationBillingPlan{
		// 														"P1Y": []*armreservations.ReservationBillingPlan{
		// 															to.Ptr(armreservations.ReservationBillingPlanUpfront),
		// 															to.Ptr(armreservations.ReservationBillingPlanMonthly)},
		// 															"P3Y": []*armreservations.ReservationBillingPlan{
		// 																to.Ptr(armreservations.ReservationBillingPlanUpfront),
		// 																to.Ptr(armreservations.ReservationBillingPlanMonthly)},
		// 															},
		// 															Locations: []*string{
		// 																to.Ptr("eastus")},
		// 																ResourceType: to.Ptr("VirtualMachines"),
		// 																Restrictions: []*armreservations.SKURestriction{
		// 																	{
		// 																		Type: to.Ptr("Location"),
		// 																		ReasonCode: to.Ptr("NotAvailableForSubscription"),
		// 																		Values: []*string{
		// 																			to.Ptr("eastus")},
		// 																	}},
		// 																	SKUProperties: []*armreservations.SKUProperty{
		// 																		{
		// 																			Name: to.Ptr("Cores"),
		// 																			Value: to.Ptr("2"),
		// 																		},
		// 																		{
		// 																			Name: to.Ptr("ProductTitle"),
		// 																			Value: to.Ptr("F Series, F2"),
		// 																		},
		// 																		{
		// 																			Name: to.Ptr("ProductShortName"),
		// 																			Value: to.Ptr("F Series"),
		// 																		},
		// 																		{
		// 																			Name: to.Ptr("SKUName"),
		// 																			Value: to.Ptr("F2"),
		// 																		},
		// 																		{
		// 																			Name: to.Ptr("MeterId"),
		// 																			Value: to.Ptr("12bc208b-083f-4901-ae85-4f98c0c3b4b8"),
		// 																	}},
		// 																	Terms: []*armreservations.ReservationTerm{
		// 																		to.Ptr(armreservations.ReservationTermP1Y),
		// 																		to.Ptr(armreservations.ReservationTermP3Y)},
		// 																}},
		// 															}
	}
}
