package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/106483d9f698ac3b6c0d481ab0c5fab14152e21f/specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/putComputeOneSkuQuotaRequest.json
func ExampleClient_BeginCreateOrUpdate_quotasPutRequestForCompute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginCreateOrUpdate(ctx, "standardFSv2Family", "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus", armquota.CurrentQuotaLimitBase{
		Properties: &armquota.Properties{
			Name: &armquota.ResourceName{
				Value: to.Ptr("standardFSv2Family"),
			},
			Limit: &armquota.LimitObject{
				LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
				Value:           to.Ptr[int32](10),
			},
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
	// 	Name: to.Ptr("standardFSv2Family"),
	// 	Type: to.Ptr("Microsoft.Quota/quotas"),
	// 	ID: to.Ptr("/subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/quotas/standardFSv2Family"),
	// 	Properties: &armquota.Properties{
	// 		Name: &armquota.ResourceName{
	// 			Value: to.Ptr("standardFSv2Family"),
	// 		},
	// 		Limit: &armquota.LimitObject{
	// 			LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
	// 			Value: to.Ptr[int32](10),
	// 		},
	// 	},
	// }
}
