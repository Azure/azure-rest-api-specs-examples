package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/106483d9f698ac3b6c0d481ab0c5fab14152e21f/specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/getComputeQuotaLimits.json
func ExampleClient_NewListPager_quotasListQuotaLimitsForCompute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus", nil)
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
		// page.Limits = armquota.Limits{
		// 	Value: []*armquota.CurrentQuotaLimitBase{
		// 		{
		// 			Name: to.Ptr("standardFSv2Family"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Quotas/standardFSv2Family"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard FSv2 Family vCPUs"),
		// 					Value: to.Ptr("standardFSv2Family"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("standardNDSFamily"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Quotas/standardNDSFamily"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard NDS Family vCPUs"),
		// 					Value: to.Ptr("standardNDSFamily"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("standardNCSv2Family"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Quotas/standardNCSv2Family"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard NCSv2 Family vCPUs"),
		// 					Value: to.Ptr("standardNCSv2Family"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("standardNCSv3Family"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Quotas/standardNCSv3Family"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard NCSv3 Family vCPUs"),
		// 					Value: to.Ptr("standardNCSv3Family"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("standardLSv2Family"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Quotas/standardLSv2Family"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard LSv2 Family vCPUs"),
		// 					Value: to.Ptr("standardLSv2Family"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("standardPBSFamily"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Quotas/standardPBSFamily"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard PBS Family vCPUs"),
		// 					Value: to.Ptr("standardPBSFamily"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("standardEIv3Family"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Quotas/standardEIv3Family"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard EIv3 Family vCPUs"),
		// 					Value: to.Ptr("standardEIv3Family"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("standardEISv3Family"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Quotas/standardEISv3Family"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard EISv3 Family vCPUs"),
		// 					Value: to.Ptr("standardEISv3Family"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("standardDCSFamily"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Quotas/standardDCSFamily"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard DCS Family vCPUs"),
		// 					Value: to.Ptr("standardDCSFamily"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("standardNVSv2Family"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Quotas/standardNVSv2Family"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard NVSv2 Family vCPUs"),
		// 					Value: to.Ptr("standardNVSv2Family"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("standardMSv2Family"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Quotas/standardMSv2Family"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard MSv2 Family vCPUs"),
		// 					Value: to.Ptr("standardMSv2Family"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("availabilitySets"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Quotas/availabilitySets"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Availability Sets"),
		// 					Value: to.Ptr("availabilitySets"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 	}},
		// }
	}
}
