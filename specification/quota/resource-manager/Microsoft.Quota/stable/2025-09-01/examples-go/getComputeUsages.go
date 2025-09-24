package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota/v2"
)

// Generated from example definition: 2025-09-01/getComputeUsages.json
func ExampleUsagesClient_NewListPager_quotasListUsagesForCompute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus", nil)
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
		// page = armquota.UsagesClientListResponse{
		// 	UsagesLimits: armquota.UsagesLimits{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/usages?api-version=2025-07-15&$skipToken=eyJDb250aW51YXRpb25Ub2tlbiI6eyJUb2tlbiI6IitSSUQiOiIxMiIsIlBLUmFuZ2UiOnsibWluIjoiIiwibWF4IjoiRkYifX19"),
		// 		Value: []*armquota.CurrentUsagesBase{
		// 			{
		// 				Name: to.Ptr("standardFSv2Family"),
		// 				Type: to.Ptr("Microsoft.Quota/Usages"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Usages/standardFSv2Family"),
		// 				Properties: &armquota.UsagesProperties{
		// 					Name: &armquota.ResourceName{
		// 						LocalizedValue: to.Ptr("Standard FSv2 Family vCPUs"),
		// 						Value: to.Ptr("standardFSv2Family"),
		// 					},
		// 					IsQuotaApplicable: to.Ptr(true),
		// 					Unit: to.Ptr("Count"),
		// 					Usages: &armquota.UsagesObject{
		// 						Value: to.Ptr[int32](10),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("standardNDSFamily"),
		// 				Type: to.Ptr("Microsoft.Quota/Usages"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Usages/standardNDSFamily"),
		// 				Properties: &armquota.UsagesProperties{
		// 					Name: &armquota.ResourceName{
		// 						LocalizedValue: to.Ptr("Standard NDS Family vCPUs"),
		// 						Value: to.Ptr("standardNDSFamily"),
		// 					},
		// 					IsQuotaApplicable: to.Ptr(true),
		// 					Unit: to.Ptr("Count"),
		// 					Usages: &armquota.UsagesObject{
		// 						Value: to.Ptr[int32](10),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("standardNCSv2Family"),
		// 				Type: to.Ptr("Microsoft.Quota/Usages"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Usages/standardNCSv2Family"),
		// 				Properties: &armquota.UsagesProperties{
		// 					Name: &armquota.ResourceName{
		// 						LocalizedValue: to.Ptr("Standard NCSv2 Family vCPUs"),
		// 						Value: to.Ptr("standardNCSv2Family"),
		// 					},
		// 					IsQuotaApplicable: to.Ptr(true),
		// 					Unit: to.Ptr("Count"),
		// 					Usages: &armquota.UsagesObject{
		// 						Value: to.Ptr[int32](10),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("standardNCSv3Family"),
		// 				Type: to.Ptr("Microsoft.Quota/Usages"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Usages/standardNCSv3Family"),
		// 				Properties: &armquota.UsagesProperties{
		// 					Name: &armquota.ResourceName{
		// 						LocalizedValue: to.Ptr("Standard NCSv3 Family vCPUs"),
		// 						Value: to.Ptr("standardNCSv3Family"),
		// 					},
		// 					IsQuotaApplicable: to.Ptr(true),
		// 					Unit: to.Ptr("Count"),
		// 					Usages: &armquota.UsagesObject{
		// 						Value: to.Ptr[int32](10),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("standardLSv2Family"),
		// 				Type: to.Ptr("Microsoft.Quota/Usages"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Usages/standardLSv2Family"),
		// 				Properties: &armquota.UsagesProperties{
		// 					Name: &armquota.ResourceName{
		// 						LocalizedValue: to.Ptr("Standard LSv2 Family vCPUs"),
		// 						Value: to.Ptr("standardLSv2Family"),
		// 					},
		// 					IsQuotaApplicable: to.Ptr(true),
		// 					Unit: to.Ptr("Count"),
		// 					Usages: &armquota.UsagesObject{
		// 						Value: to.Ptr[int32](10),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("standardPBSFamily"),
		// 				Type: to.Ptr("Microsoft.Quota/Usages"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Usages/standardPBSFamily"),
		// 				Properties: &armquota.UsagesProperties{
		// 					Name: &armquota.ResourceName{
		// 						LocalizedValue: to.Ptr("Standard PBS Family vCPUs"),
		// 						Value: to.Ptr("standardPBSFamily"),
		// 					},
		// 					IsQuotaApplicable: to.Ptr(true),
		// 					Unit: to.Ptr("Count"),
		// 					Usages: &armquota.UsagesObject{
		// 						Value: to.Ptr[int32](10),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("standardEIv3Family"),
		// 				Type: to.Ptr("Microsoft.Quota/Usages"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Usages/standardEIv3Family"),
		// 				Properties: &armquota.UsagesProperties{
		// 					Name: &armquota.ResourceName{
		// 						LocalizedValue: to.Ptr("Standard EIv3 Family vCPUs"),
		// 						Value: to.Ptr("standardEIv3Family"),
		// 					},
		// 					IsQuotaApplicable: to.Ptr(true),
		// 					Unit: to.Ptr("Count"),
		// 					Usages: &armquota.UsagesObject{
		// 						Value: to.Ptr[int32](10),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("standardEISv3Family"),
		// 				Type: to.Ptr("Microsoft.Quota/Usages"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Usages/standardEISv3Family"),
		// 				Properties: &armquota.UsagesProperties{
		// 					Name: &armquota.ResourceName{
		// 						LocalizedValue: to.Ptr("Standard EISv3 Family vCPUs"),
		// 						Value: to.Ptr("standardEISv3Family"),
		// 					},
		// 					IsQuotaApplicable: to.Ptr(true),
		// 					Unit: to.Ptr("Count"),
		// 					Usages: &armquota.UsagesObject{
		// 						Value: to.Ptr[int32](10),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("standardDCSFamily"),
		// 				Type: to.Ptr("Microsoft.Quota/Usages"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Usages/standardDCSFamily"),
		// 				Properties: &armquota.UsagesProperties{
		// 					Name: &armquota.ResourceName{
		// 						LocalizedValue: to.Ptr("Standard DCS Family vCPUs"),
		// 						Value: to.Ptr("standardDCSFamily"),
		// 					},
		// 					IsQuotaApplicable: to.Ptr(true),
		// 					Unit: to.Ptr("Count"),
		// 					Usages: &armquota.UsagesObject{
		// 						Value: to.Ptr[int32](10),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("standardNVSv2Family"),
		// 				Type: to.Ptr("Microsoft.Quota/Usages"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Usages/standardNVSv2Family"),
		// 				Properties: &armquota.UsagesProperties{
		// 					Name: &armquota.ResourceName{
		// 						LocalizedValue: to.Ptr("Standard NVSv2 Family vCPUs"),
		// 						Value: to.Ptr("standardNVSv2Family"),
		// 					},
		// 					IsQuotaApplicable: to.Ptr(true),
		// 					Unit: to.Ptr("Count"),
		// 					Usages: &armquota.UsagesObject{
		// 						Value: to.Ptr[int32](10),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("standardMSv2Family"),
		// 				Type: to.Ptr("Microsoft.Quota/Usages"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Usages/standardMSv2Family"),
		// 				Properties: &armquota.UsagesProperties{
		// 					Name: &armquota.ResourceName{
		// 						LocalizedValue: to.Ptr("Standard MSv2 Family vCPUs"),
		// 						Value: to.Ptr("standardMSv2Family"),
		// 					},
		// 					IsQuotaApplicable: to.Ptr(true),
		// 					Unit: to.Ptr("Count"),
		// 					Usages: &armquota.UsagesObject{
		// 						Value: to.Ptr[int32](10),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("availabilitySets"),
		// 				Type: to.Ptr("Microsoft.Quota/Usages"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Usages/availabilitySets"),
		// 				Properties: &armquota.UsagesProperties{
		// 					Name: &armquota.ResourceName{
		// 						LocalizedValue: to.Ptr("Availability Sets"),
		// 						Value: to.Ptr("availabilitySets"),
		// 					},
		// 					IsQuotaApplicable: to.Ptr(false),
		// 					Unit: to.Ptr("Count"),
		// 					Usages: &armquota.UsagesObject{
		// 						Value: to.Ptr[int32](10),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
