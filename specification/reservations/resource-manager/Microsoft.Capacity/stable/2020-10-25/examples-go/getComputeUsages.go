package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getComputeUsages.json
func ExampleQuotaClient_NewListPager_quotasListUsagesForCompute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armreservations.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewQuotaClient().NewListPager("00000000-0000-0000-0000-000000000000", "Microsoft.Compute", "eastus", nil)
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
		// page.QuotaLimits = armreservations.QuotaLimits{
		// 	Value: []*armreservations.CurrentQuotaLimitBase{
		// 		{
		// 			Properties: &armreservations.QuotaProperties{
		// 				Name: &armreservations.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard FSv2 Family vCPUs"),
		// 					Value: to.Ptr("standardFSv2Family"),
		// 				},
		// 				CurrentValue: to.Ptr[int32](0),
		// 				Limit: to.Ptr[int32](100),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Properties: &armreservations.QuotaProperties{
		// 				Name: &armreservations.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard NDS Family vCPUs"),
		// 					Value: to.Ptr("standardNDSFamily"),
		// 				},
		// 				CurrentValue: to.Ptr[int32](0),
		// 				Limit: to.Ptr[int32](0),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Properties: &armreservations.QuotaProperties{
		// 				Name: &armreservations.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard NCSv2 Family vCPUs"),
		// 					Value: to.Ptr("standardNCSv2Family"),
		// 				},
		// 				CurrentValue: to.Ptr[int32](0),
		// 				Limit: to.Ptr[int32](0),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Properties: &armreservations.QuotaProperties{
		// 				Name: &armreservations.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard NCSv3 Family vCPUs"),
		// 					Value: to.Ptr("standardNCSv3Family"),
		// 				},
		// 				CurrentValue: to.Ptr[int32](0),
		// 				Limit: to.Ptr[int32](0),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Properties: &armreservations.QuotaProperties{
		// 				Name: &armreservations.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard LSv2 Family vCPUs"),
		// 					Value: to.Ptr("standardLSv2Family"),
		// 				},
		// 				CurrentValue: to.Ptr[int32](0),
		// 				Limit: to.Ptr[int32](100),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Properties: &armreservations.QuotaProperties{
		// 				Name: &armreservations.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard PBS Family vCPUs"),
		// 					Value: to.Ptr("standardPBSFamily"),
		// 				},
		// 				CurrentValue: to.Ptr[int32](0),
		// 				Limit: to.Ptr[int32](6),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Properties: &armreservations.QuotaProperties{
		// 				Name: &armreservations.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard EIv3 Family vCPUs"),
		// 					Value: to.Ptr("standardEIv3Family"),
		// 				},
		// 				CurrentValue: to.Ptr[int32](0),
		// 				Limit: to.Ptr[int32](100),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Properties: &armreservations.QuotaProperties{
		// 				Name: &armreservations.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard EISv3 Family vCPUs"),
		// 					Value: to.Ptr("standardEISv3Family"),
		// 				},
		// 				CurrentValue: to.Ptr[int32](0),
		// 				Limit: to.Ptr[int32](100),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Properties: &armreservations.QuotaProperties{
		// 				Name: &armreservations.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard DCS Family vCPUs"),
		// 					Value: to.Ptr("standardDCSFamily"),
		// 				},
		// 				CurrentValue: to.Ptr[int32](0),
		// 				Limit: to.Ptr[int32](8),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Properties: &armreservations.QuotaProperties{
		// 				Name: &armreservations.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard NVSv2 Family vCPUs"),
		// 					Value: to.Ptr("standardNVSv2Family"),
		// 				},
		// 				CurrentValue: to.Ptr[int32](0),
		// 				Limit: to.Ptr[int32](0),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Properties: &armreservations.QuotaProperties{
		// 				Name: &armreservations.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard MSv2 Family vCPUs"),
		// 					Value: to.Ptr("standardMSv2Family"),
		// 				},
		// 				CurrentValue: to.Ptr[int32](0),
		// 				Limit: to.Ptr[int32](0),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 	}},
		// }
	}
}
