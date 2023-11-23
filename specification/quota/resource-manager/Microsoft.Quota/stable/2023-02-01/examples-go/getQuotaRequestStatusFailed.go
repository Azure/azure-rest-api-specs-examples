package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/getQuotaRequestStatusFailed.json
func ExampleRequestStatusClient_Get_quotaRequestFailed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRequestStatusClient().Get(ctx, "2B5C8515-37D8-4B6A-879B-CD641A2CF605", "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RequestDetails = armquota.RequestDetails{
	// 	Name: to.Ptr("2B5C8515-37D8-4B6A-879B-CD641A2CF605"),
	// 	Type: to.Ptr("Microsoft.Quota/quotaRequests"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/quotaRequests/2B5C8515-37D8-4B6A-879B-CD641A2CF605"),
	// 	Properties: &armquota.RequestProperties{
	// 		Error: &armquota.ServiceErrorDetail{
	// 			Code: to.Ptr("ContactSupport"),
	// 			Message: to.Ptr("Request failed, please contact support."),
	// 		},
	// 		Message: to.Ptr("Request failed, please contact support."),
	// 		ProvisioningState: to.Ptr(armquota.QuotaRequestStateSucceeded),
	// 		RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-19T19:23:17.904Z"); return t}()),
	// 		Value: []*armquota.SubRequest{
	// 			{
	// 				Name: &armquota.ResourceName{
	// 					LocalizedValue: to.Ptr("Standard NC Promo Family vCPUs"),
	// 					Value: to.Ptr("standardNCPromoFamily"),
	// 				},
	// 				Limit: &armquota.LimitObject{
	// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
	// 					Value: to.Ptr[int32](50),
	// 				},
	// 				Message: to.Ptr("Request failed, please contact support."),
	// 				ProvisioningState: to.Ptr(armquota.QuotaRequestStateSucceeded),
	// 				SubRequestID: to.Ptr("AD07450A-DE86-4FD3-859B-107BEF218C4C"),
	// 		}},
	// 	},
	// }
}
