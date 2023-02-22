package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsListSkus.json
func ExampleKustoPoolsClient_NewListSKUsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListSKUsPager(nil)
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
		// page.SKUDescriptionList = armsynapse.SKUDescriptionList{
		// 	Value: []*armsynapse.SKUDescription{
		// 		{
		// 			Name: to.Ptr("Compute optimized"),
		// 			LocationInfo: []*armsynapse.SKULocationInfoItem{
		// 				{
		// 					Location: to.Ptr("West US"),
		// 					Zones: []*string{
		// 						to.Ptr("1"),
		// 						to.Ptr("2"),
		// 						to.Ptr("3")},
		// 					},
		// 					{
		// 						Location: to.Ptr("West Europe"),
		// 						Zones: []*string{
		// 						},
		// 				}},
		// 				Locations: []*string{
		// 					to.Ptr("West US"),
		// 					to.Ptr("West Europe")},
		// 					Size: to.Ptr("Medium"),
		// 				},
		// 				{
		// 					Name: to.Ptr("Compute optimized"),
		// 					LocationInfo: []*armsynapse.SKULocationInfoItem{
		// 						{
		// 							Location: to.Ptr("West US"),
		// 							Zones: []*string{
		// 								to.Ptr("1"),
		// 								to.Ptr("2"),
		// 								to.Ptr("3")},
		// 							},
		// 							{
		// 								Location: to.Ptr("West Europe"),
		// 								Zones: []*string{
		// 								},
		// 						}},
		// 						Locations: []*string{
		// 							to.Ptr("West US"),
		// 							to.Ptr("West Europe")},
		// 							Size: to.Ptr("Large"),
		// 						},
		// 						{
		// 							Name: to.Ptr("Storage optimized"),
		// 							LocationInfo: []*armsynapse.SKULocationInfoItem{
		// 								{
		// 									Location: to.Ptr("West US"),
		// 									Zones: []*string{
		// 										to.Ptr("1"),
		// 										to.Ptr("2"),
		// 										to.Ptr("3")},
		// 									},
		// 									{
		// 										Location: to.Ptr("West Europe"),
		// 										Zones: []*string{
		// 										},
		// 								}},
		// 								Locations: []*string{
		// 									to.Ptr("West US"),
		// 									to.Ptr("West Europe")},
		// 									Size: to.Ptr("Medium"),
		// 								},
		// 								{
		// 									Name: to.Ptr("Storage optimized"),
		// 									LocationInfo: []*armsynapse.SKULocationInfoItem{
		// 										{
		// 											Location: to.Ptr("West US"),
		// 											Zones: []*string{
		// 												to.Ptr("1"),
		// 												to.Ptr("2"),
		// 												to.Ptr("3")},
		// 											},
		// 											{
		// 												Location: to.Ptr("West Europe"),
		// 												Zones: []*string{
		// 												},
		// 										}},
		// 										Locations: []*string{
		// 											to.Ptr("West US"),
		// 											to.Ptr("West Europe")},
		// 											Size: to.Ptr("Large"),
		// 									}},
		// 								}
	}
}
