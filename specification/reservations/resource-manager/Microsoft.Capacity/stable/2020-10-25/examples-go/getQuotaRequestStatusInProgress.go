package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getQuotaRequestStatusInProgress.json
func ExampleQuotaRequestStatusClient_Get_quotaRequestInProgress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armreservations.NewQuotaRequestStatusClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "00000000-0000-0000-0000-000000000000", "Microsoft.Compute", "eastus", "2B5C8515-37D8-4B6A-879B-CD641A2CF605", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QuotaRequestDetails = armreservations.QuotaRequestDetails{
	// 	Name: to.Ptr("00000000-0000-0000-0000-000000000000/2B5C8515-37D8-4B6A-879B-CD641A2CF605"),
	// 	Type: to.Ptr("Microsoft.Capacity/serviceLimitsRequests"),
	// 	ID: to.Ptr("2B5C8515-37D8-4B6A-879B-CD641A2CF605"),
	// 	Properties: &armreservations.QuotaRequestProperties{
	// 		Message: to.Ptr("Request processing"),
	// 		ProvisioningState: to.Ptr(armreservations.QuotaRequestStateSucceeded),
	// 		RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-19T19:23:17.904Z"); return t}()),
	// 		Value: []*armreservations.SubRequest{
	// 			{
	// 				Name: &armreservations.ResourceName{
	// 					LocalizedValue: to.Ptr("Standard NC Promo Family vCPUs"),
	// 					Value: to.Ptr("standardNCPromoFamily"),
	// 				},
	// 				Limit: to.Ptr[int32](50),
	// 				Message: to.Ptr("Request processing"),
	// 				ProvisioningState: to.Ptr(armreservations.QuotaRequestStateSucceeded),
	// 				SubRequestID: to.Ptr("AD07450A-DE86-4FD3-859B-107BEF218C4C"),
	// 		}},
	// 	},
	// }
}
