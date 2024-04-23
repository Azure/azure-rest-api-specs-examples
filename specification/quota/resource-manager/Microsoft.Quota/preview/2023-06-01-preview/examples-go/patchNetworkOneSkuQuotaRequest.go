package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/106483d9f698ac3b6c0d481ab0c5fab14152e21f/specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/patchNetworkOneSkuQuotaRequest.json
func ExampleClient_BeginUpdate_quotasRequestPatchForNetwork() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginUpdate(ctx, "MinPublicIpInterNetworkPrefixLength", "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Network/locations/eastus", armquota.CurrentQuotaLimitBase{
		Properties: &armquota.Properties{
			Name: &armquota.ResourceName{
				Value: to.Ptr("MinPublicIpInterNetworkPrefixLength"),
			},
			Limit: &armquota.LimitObject{
				LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
				Value:           to.Ptr[int32](10),
			},
			ResourceType: to.Ptr("MinPublicIpInterNetworkPrefixLength"),
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
	// 	Name: to.Ptr("2B5C8515-37D8-4B6A-879B-CD641A2CF605"),
	// 	Type: to.Ptr("Microsoft.Quota/quotas"),
	// 	ID: to.Ptr("/subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/quotaRequests/2B5C8515-37D8-4B6A-879B-CD641A2CF605"),
	// 	Properties: &armquota.Properties{
	// 		Name: &armquota.ResourceName{
	// 			Value: to.Ptr("MinPublicIpInterNetworkPrefixLength"),
	// 		},
	// 		Limit: &armquota.LimitObject{
	// 			LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
	// 			Value: to.Ptr[int32](10),
	// 		},
	// 		ResourceType: to.Ptr("MinPublicIpInterNetworkPrefixLength"),
	// 	},
	// }
}
