package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c178f2467f792a322f56174be244135d2c907f/specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_usages.json
func ExampleResourceProviderCommonClient_GetSubscriptionQuota() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceProviderCommonClient().GetSubscriptionQuota(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UserSubscriptionQuotaListResult = armiothub.UserSubscriptionQuotaListResult{
	// 	Value: []*armiothub.UserSubscriptionQuota{
	// 		{
	// 			Name: &armiothub.Name{
	// 				LocalizedValue: to.Ptr("Free Hub Count"),
	// 				Value: to.Ptr("FreeHubCount"),
	// 			},
	// 			Type: to.Ptr("/subscription/91d12660-3dec-467a-be2a-213b5544ddc0/providers/Microsoft.Devices/usages"),
	// 			CurrentValue: to.Ptr[int32](1),
	// 			ID: to.Ptr("/subscription/91d12660-3dec-467a-be2a-213b5544ddc0/providers/Microsoft.Devices/usages/freeHubCount"),
	// 			Limit: to.Ptr[int32](1),
	// 			Unit: to.Ptr("count"),
	// 	}},
	// }
}
