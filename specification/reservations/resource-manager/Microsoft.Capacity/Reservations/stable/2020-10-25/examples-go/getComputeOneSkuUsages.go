package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v4"
)

// Generated from example definition: 2020-10-25/getComputeOneSkuUsages.json
func ExampleQuotaClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armreservations.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQuotaClient().Get(ctx, "00000000-0000-0000-0000-000000000000", "Microsoft.Compute", "eastus", "standardNDSFamily", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armreservations.QuotaClientGetResponse{
	// 	CurrentQuotaLimitBase: armreservations.CurrentQuotaLimitBase{
	// 		Name: to.Ptr("standardNDSFamily"),
	// 		Type: to.Ptr("Microsoft.Capacity/ServiceLimits"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Capacity/resourceProviders/Microsoft.Compute/locations/eastus/serviceLimits/standardNDSFamily"),
	// 		Properties: &armreservations.QuotaProperties{
	// 			Name: &armreservations.ResourceName{
	// 				LocalizedValue: to.Ptr("Standard NDS Family vCPUs"),
	// 				Value: to.Ptr("standardNDSFamily"),
	// 			},
	// 			CurrentValue: to.Ptr[int32](0),
	// 			Limit: to.Ptr[int32](10),
	// 			Unit: to.Ptr("Count"),
	// 		},
	// 	},
	// }
}
