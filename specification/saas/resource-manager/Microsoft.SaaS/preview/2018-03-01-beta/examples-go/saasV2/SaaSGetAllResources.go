package armsaas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/saas/resource-manager/Microsoft.SaaS/preview/2018-03-01-beta/examples/saasV2/SaaSGetAllResources.json
func ExampleResourcesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsaas.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourcesClient().NewListPager(nil)
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
		// page.ResourceResponseWithContinuation = armsaas.ResourceResponseWithContinuation{
		// 	Value: []*armsaas.Resource{
		// 		{
		// 			Name: to.Ptr("diyakobo-transfer"),
		// 			Type: to.Ptr("Microsoft.SaaS/saasresources"),
		// 			ID: to.Ptr("/providers/Microsoft.SaaS/saasresources/115c3523-1fae-757f-af86-7b27cfd29805"),
		// 			Properties: &armsaas.ResourceProperties{
		// 				OfferID: to.Ptr("microsofthealthcarebot"),
		// 				PaymentChannelMetadata: map[string]*string{
		// 					"azureSubscriptionId": to.Ptr("155af98a-3205-47e7-883b-a2ab9db9f88d"),
		// 					"resourceId": to.Ptr("b25dba9a-1bd7-4600-9447-3d27d11a6477"),
		// 				},
		// 				PublisherID: to.Ptr("microsoft-hcb"),
		// 				SaasResourceName: to.Ptr("diyakobo-transfer"),
		// 				SKUID: to.Ptr("free"),
		// 				Created: to.Ptr("2020-10-12T05:08:40.9235607Z"),
		// 				IsFreeTrial: to.Ptr(false),
		// 				LastModified: to.Ptr("2020-11-12T21:25:40.2736665Z"),
		// 				Status: to.Ptr(armsaas.SaasResourceStatusSubscribed),
		// 				Term: &armsaas.PropertiesTerm{
		// 					EndDate: to.Ptr("2020-12-11T00:00:00Z"),
		// 					StartDate: to.Ptr("2020-11-12T00:00:00Z"),
		// 					TermUnit: to.Ptr("P1M"),
		// 				},
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-transfer"),
		// 			Type: to.Ptr("Microsoft.SaaS/saasresources"),
		// 			ID: to.Ptr("/providers/Microsoft.SaaS/saasresources/7a4a733c-2204-ee2c-2f51-47c2dfbcb7fd"),
		// 			Properties: &armsaas.ResourceProperties{
		// 				OfferID: to.Ptr("microsofthealthcarebot"),
		// 				PaymentChannelMetadata: map[string]*string{
		// 					"azureSubscriptionId": to.Ptr("155af98a-3205-47e7-883b-a2ab9db9f88d"),
		// 					"resourceId": to.Ptr("f5816b76-67e3-45e1-b331-91fbb021e068"),
		// 				},
		// 				PublisherID: to.Ptr("microsoft-hcb"),
		// 				SaasResourceName: to.Ptr("test-transfer"),
		// 				SKUID: to.Ptr("free"),
		// 				Created: to.Ptr("2020-09-30T16:45:41.2981172Z"),
		// 				IsFreeTrial: to.Ptr(false),
		// 				LastModified: to.Ptr("2020-10-30T16:22:15.4319475Z"),
		// 				Status: to.Ptr(armsaas.SaasResourceStatusSubscribed),
		// 				Term: &armsaas.PropertiesTerm{
		// 					EndDate: to.Ptr("2020-11-29T00:00:00Z"),
		// 					StartDate: to.Ptr("2020-10-30T00:00:00Z"),
		// 					TermUnit: to.Ptr("P1M"),
		// 				},
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 	}},
		// }
	}
}
