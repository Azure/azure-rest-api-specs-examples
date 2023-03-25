package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/putMachineLearningServicesQuotaRequestDedicated.json
func ExampleQuotaClient_BeginCreateOrUpdate_quotasRequestPutForMachineLearningServicesDedicatedResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armreservations.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewQuotaClient().BeginCreateOrUpdate(ctx, "D7EC67B3-7657-4966-BFFC-41EFD36BAAB3", "Microsoft.MachineLearningServices", "eastus", "StandardDv2Family", armreservations.CurrentQuotaLimitBase{
		Properties: &armreservations.QuotaProperties{
			Name: &armreservations.ResourceName{
				Value: to.Ptr("StandardDv2Family"),
			},
			Limit:        to.Ptr[int32](200),
			ResourceType: to.Ptr(armreservations.ResourceTypeDedicated),
			Unit:         to.Ptr("Count"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CurrentQuotaLimitBase = armreservations.CurrentQuotaLimitBase{
	// 	Name: to.Ptr("2B5C8515-37D8-4B6A-879B-CD641A2CF605"),
	// 	Type: to.Ptr("Microsoft.Capacity/serviceLimits"),
	// 	ID: to.Ptr("/subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Capacity/resourceProviders/Microsoft.MachineLearningServices/locations/eastus/serviceLimitsRequests/2B5C8515-37D8-4B6A-879B-CD641A2CF605"),
	// 	Properties: &armreservations.QuotaProperties{
	// 		Name: &armreservations.ResourceName{
	// 			LocalizedValue: to.Ptr("Standard Dv2 Family Cluster Dedicated vCPUs"),
	// 			Value: to.Ptr("StandardDv2Family"),
	// 		},
	// 		CurrentValue: to.Ptr[int32](160),
	// 		Limit: to.Ptr[int32](200),
	// 		Properties: map[string]any{
	// 		},
	// 		QuotaPeriod: to.Ptr(""),
	// 		ResourceType: to.Ptr(armreservations.ResourceTypeDedicated),
	// 		Unit: to.Ptr("Count"),
	// 	},
	// }
}
