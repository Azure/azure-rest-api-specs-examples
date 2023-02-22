package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getComputeOneSkuUsages.json
func ExampleQuotaClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armreservations.NewQuotaClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "00000000-0000-0000-0000-000000000000", "Microsoft.Compute", "eastus", "standardNDSFamily", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CurrentQuotaLimitBase = armreservations.CurrentQuotaLimitBase{
	// 	Properties: &armreservations.QuotaProperties{
	// 		Name: &armreservations.ResourceName{
	// 			LocalizedValue: to.Ptr("Standard NDS Family vCPUs"),
	// 			Value: to.Ptr("standardNDSFamily"),
	// 		},
	// 		CurrentValue: to.Ptr[int32](0),
	// 		Limit: to.Ptr[int32](10),
	// 		Unit: to.Ptr("Count"),
	// 	},
	// }
}
