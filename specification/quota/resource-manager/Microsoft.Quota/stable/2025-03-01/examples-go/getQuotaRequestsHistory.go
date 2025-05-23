package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8691e5081766c7ad602a9e55de841d07bed5196a/specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/getQuotaRequestsHistory.json
func ExampleRequestStatusClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRequestStatusClient().NewListPager("subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus", &armquota.RequestStatusClientListOptions{Filter: nil,
		Top:       nil,
		Skiptoken: nil,
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
		// page.RequestDetailsList = armquota.RequestDetailsList{
		// 	Value: []*armquota.RequestDetails{
		// 		{
		// 			Name: to.Ptr("2B5C8515-37D8-4B6A-879B-CD641A2CF605"),
		// 			Type: to.Ptr("Microsoft.Quota/quotaRequests"),
		// 			ID: to.Ptr("/subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/quotaRequests/2B5C8515-37D8-4B6A-879B-CD641A2CF605"),
		// 			Properties: &armquota.RequestProperties{
		// 				Message: to.Ptr("Request completed"),
		// 				ProvisioningState: to.Ptr(armquota.QuotaRequestStateSucceeded),
		// 				RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-19T19:23:17.904Z"); return t}()),
		// 				Value: []*armquota.SubRequest{
		// 					{
		// 						Name: &armquota.ResourceName{
		// 							LocalizedValue: to.Ptr("Standard NC Promo Family vCPUs"),
		// 							Value: to.Ptr("standardNCPromoFamily"),
		// 						},
		// 						Limit: &armquota.LimitObject{
		// 							LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 							Value: to.Ptr[int32](50),
		// 						},
		// 						Message: to.Ptr("Request completed"),
		// 						ProvisioningState: to.Ptr(armquota.QuotaRequestStateSucceeded),
		// 						SubRequestID: to.Ptr("AD07450A-DE86-4FD3-859B-107BEF218C4C"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("7E73A85C-83BB-4DE4-903F-076F1A2B91D6"),
		// 			Type: to.Ptr("Microsoft.Quota/quotaRequests"),
		// 			ID: to.Ptr("/subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/quotaRequests/7E73A85C-83BB-4DE4-903F-076F1A2B91D6"),
		// 			Properties: &armquota.RequestProperties{
		// 				Message: to.Ptr("Request completed"),
		// 				ProvisioningState: to.Ptr(armquota.QuotaRequestStateSucceeded),
		// 				RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-18T19:23:17.904Z"); return t}()),
		// 				Value: []*armquota.SubRequest{
		// 					{
		// 						Name: &armquota.ResourceName{
		// 							LocalizedValue: to.Ptr("Standard NV Promo Family vCPUs"),
		// 							Value: to.Ptr("standardNVPromoFamily"),
		// 						},
		// 						Limit: &armquota.LimitObject{
		// 							LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 							Value: to.Ptr[int32](150),
		// 						},
		// 						Message: to.Ptr("Request completed"),
		// 						ProvisioningState: to.Ptr(armquota.QuotaRequestStateSucceeded),
		// 						SubRequestID: to.Ptr("AD07450A-DE86-4FD3-859B-107BEF218C4C"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("5E460077-AB53-4802-8997-A6940E0B7649"),
		// 			Type: to.Ptr("Microsoft.Quota/quotaRequests"),
		// 			ID: to.Ptr("/subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/quotaRequests/5E460077-AB53-4802-8997-A6940E0B7649"),
		// 			Properties: &armquota.RequestProperties{
		// 				Error: &armquota.ServiceErrorDetail{
		// 					Code: to.Ptr("ResourceNotAvailableForOffer"),
		// 					Message: to.Ptr("The resource is currently unavailable in the location for the offer type."),
		// 				},
		// 				Message: to.Ptr("The resource is currently unavailable in the location for the offer type."),
		// 				ProvisioningState: to.Ptr(armquota.QuotaRequestStateSucceeded),
		// 				RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-17T19:23:17.904Z"); return t}()),
		// 				Value: []*armquota.SubRequest{
		// 					{
		// 						Name: &armquota.ResourceName{
		// 							LocalizedValue: to.Ptr("Standard HBS Family vCPUs"),
		// 							Value: to.Ptr("standardHBSFamily"),
		// 						},
		// 						Limit: &armquota.LimitObject{
		// 							LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 							Value: to.Ptr[int32](52),
		// 						},
		// 						Message: to.Ptr("Request failed, please contact support."),
		// 						ProvisioningState: to.Ptr(armquota.QuotaRequestStateSucceeded),
		// 						SubRequestID: to.Ptr("AD07450A-DE86-4FD3-859B-107BEF218C4C"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
