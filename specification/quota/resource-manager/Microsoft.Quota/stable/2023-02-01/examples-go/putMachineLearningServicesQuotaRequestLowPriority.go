package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/putMachineLearningServicesQuotaRequestLowPriority.json
func ExampleClient_BeginCreateOrUpdate_quotasRequestForMachineLearningServicesLowPriorityResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginCreateOrUpdate(ctx, "TotalLowPriorityCores", "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.MachineLearningServices/locations/eastus", armquota.CurrentQuotaLimitBase{
		Properties: &armquota.Properties{
			Name: &armquota.ResourceName{
				Value: to.Ptr("TotalLowPriorityCores"),
			},
			Limit: &armquota.LimitObject{
				LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
				Value:           to.Ptr[int32](10),
			},
			ResourceType: to.Ptr("lowPriority"),
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
	// res.CurrentQuotaLimitBase = armquota.CurrentQuotaLimitBase{
	// 	Name: to.Ptr("TotalLowPriorityCores"),
	// 	Type: to.Ptr("Microsoft.Quota/quotas"),
	// 	ID: to.Ptr("/subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.MachineLearningServices/locations/eastus/providers/Microsoft.Quota/quotas/TotalLowPriorityCores"),
	// 	Properties: &armquota.Properties{
	// 		Name: &armquota.ResourceName{
	// 			Value: to.Ptr("TotalLowPriorityCores"),
	// 		},
	// 		Limit: &armquota.LimitObject{
	// 			LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
	// 			Value: to.Ptr[int32](10),
	// 		},
	// 		ResourceType: to.Ptr("lowPriority"),
	// 	},
	// }
}
