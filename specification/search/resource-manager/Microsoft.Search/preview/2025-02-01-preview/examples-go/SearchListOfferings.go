package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3db6867b8e524ea6d1bc7a3bbb989fe50dd2f184/specification/search/resource-manager/Microsoft.Search/preview/2025-02-01-preview/examples/SearchListOfferings.json
func ExampleOfferingsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOfferingsClient().NewListPager(nil)
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
		// page.OfferingsListResult = armsearch.OfferingsListResult{
		// 	Value: []*armsearch.OfferingsByRegion{
		// 		{
		// 			Features: []*armsearch.FeatureOffering{
		// 				{
		// 					Name: to.Ptr(armsearch.FeatureName("SemanticRanker")),
		// 				},
		// 				{
		// 					Name: to.Ptr(armsearch.FeatureName("QueryRewriting")),
		// 			}},
		// 			RegionName: to.Ptr("East US"),
		// 			SKUs: []*armsearch.SKUOffering{
		// 				{
		// 					Limits: &armsearch.SKUOfferingLimits{
		// 						Indexers: to.Ptr[int32](50),
		// 						Indexes: to.Ptr[int32](50),
		// 						PartitionStorageInGigabytes: to.Ptr[float32](160),
		// 						PartitionVectorStorageInGigabytes: to.Ptr[float32](35),
		// 						Partitions: to.Ptr[int32](12),
		// 						Replicas: to.Ptr[int32](12),
		// 						SearchUnits: to.Ptr[int32](36),
		// 					},
		// 					SKU: &armsearch.SKU{
		// 						Name: to.Ptr(armsearch.SKUNameStandard),
		// 					},
		// 				},
		// 				{
		// 					Limits: &armsearch.SKUOfferingLimits{
		// 						Indexers: to.Ptr[int32](200),
		// 						Indexes: to.Ptr[int32](200),
		// 						PartitionStorageInGigabytes: to.Ptr[float32](512),
		// 						PartitionVectorStorageInGigabytes: to.Ptr[float32](150),
		// 						Partitions: to.Ptr[int32](12),
		// 						Replicas: to.Ptr[int32](12),
		// 						SearchUnits: to.Ptr[int32](36),
		// 					},
		// 					SKU: &armsearch.SKU{
		// 						Name: to.Ptr(armsearch.SKUNameStandard2),
		// 					},
		// 			}},
		// 		},
		// 		{
		// 			Features: []*armsearch.FeatureOffering{
		// 				{
		// 					Name: to.Ptr(armsearch.FeatureNameGrok),
		// 				},
		// 				{
		// 					Name: to.Ptr(armsearch.FeatureNameImageVectorization),
		// 			}},
		// 			RegionName: to.Ptr("West Europe"),
		// 			SKUs: []*armsearch.SKUOffering{
		// 				{
		// 					Limits: &armsearch.SKUOfferingLimits{
		// 						Indexers: to.Ptr[int32](50),
		// 						Indexes: to.Ptr[int32](50),
		// 						PartitionStorageInGigabytes: to.Ptr[float32](160),
		// 						PartitionVectorStorageInGigabytes: to.Ptr[float32](35),
		// 						Partitions: to.Ptr[int32](12),
		// 						Replicas: to.Ptr[int32](12),
		// 						SearchUnits: to.Ptr[int32](36),
		// 					},
		// 					SKU: &armsearch.SKU{
		// 						Name: to.Ptr(armsearch.SKUNameStandard),
		// 					},
		// 				},
		// 				{
		// 					Limits: &armsearch.SKUOfferingLimits{
		// 						Indexers: to.Ptr[int32](200),
		// 						Indexes: to.Ptr[int32](200),
		// 						PartitionStorageInGigabytes: to.Ptr[float32](512),
		// 						PartitionVectorStorageInGigabytes: to.Ptr[float32](150),
		// 						Partitions: to.Ptr[int32](12),
		// 						Replicas: to.Ptr[int32](12),
		// 						SearchUnits: to.Ptr[int32](36),
		// 					},
		// 					SKU: &armsearch.SKU{
		// 						Name: to.Ptr(armsearch.SKUNameStandard2),
		// 					},
		// 			}},
		// 	}},
		// }
	}
}
