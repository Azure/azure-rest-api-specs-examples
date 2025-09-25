package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota/v2"
)

// Generated from example definition: 2025-09-01/getQuotaRequestStatusById.json
func ExampleRequestStatusClient_Get_quotaRequestStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRequestStatusClient().Get(ctx, "2B5C8515-37D8-4B6A-879B-CD641A2CF605", "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armquota.RequestStatusClientGetResponse{
	// 	RequestDetails: &armquota.RequestDetails{
	// 		Name: to.Ptr("2B5C8515-37D8-4B6A-879B-CD641A2CF605"),
	// 		Type: to.Ptr("Microsoft.Quota/quotaRequests"),
	// 		ID: to.Ptr("/subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/quotaRequests/2B5C8515-37D8-4B6A-879B-CD641A2CF605"),
	// 		Properties: &armquota.RequestProperties{
	// 			Message: to.Ptr("Request completed"),
	// 			ProvisioningState: to.Ptr(armquota.QuotaRequestStateSucceeded),
	// 			RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-19T19:23:17.904Z"); return t}()),
	// 			Value: []*armquota.SubRequest{
	// 				{
	// 					Name: &armquota.ResourceName{
	// 						LocalizedValue: to.Ptr("Standard NC Promo Family vCPUs"),
	// 						Value: to.Ptr("standardNCPromoFamily"),
	// 					},
	// 					Limit: &armquota.LimitObject{
	// 						LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
	// 						Value: to.Ptr[int32](50),
	// 					},
	// 					Message: to.Ptr("Request completed"),
	// 					ProvisioningState: to.Ptr(armquota.QuotaRequestStateSucceeded),
	// 					Unit: to.Ptr("Count"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
