package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/putMachineLearningServicesQuotaRequestLowPriority.json
func ExampleQuotaClient_BeginCreateOrUpdate_quotasRequestPutForMachineLearningServicesLowPriorityResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armreservations.NewQuotaClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "D7EC67B3-7657-4966-BFFC-41EFD36BAAB3", "Microsoft.MachineLearningServices", "eastus", "TotalLowPriorityCores", armreservations.CurrentQuotaLimitBase{
		Properties: &armreservations.QuotaProperties{
			Name: &armreservations.ResourceName{
				Value: to.Ptr("TotalLowPriorityCores"),
			},
			Limit:        to.Ptr[int32](200),
			ResourceType: to.Ptr(armreservations.ResourceTypeLowPriority),
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
	// TODO: use response item
	_ = res
}
