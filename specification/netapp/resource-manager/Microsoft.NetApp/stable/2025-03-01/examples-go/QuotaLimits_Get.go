package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/QuotaLimits_Get.json
func ExampleResourceQuotaLimitsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceQuotaLimitsClient().Get(ctx, "eastus", "totalCoolAccessVolumesPerSubscription", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SubscriptionQuotaItem = armnetapp.SubscriptionQuotaItem{
	// 	Name: to.Ptr("eastus/totalCoolAccessVolumesPerSubscription"),
	// 	Type: to.Ptr("Microsoft.NetApp/locations/quotaLimits"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.NetApp/locations/eastus/quotaLimits/totalCoolAccessVolumesPerSubscription"),
	// 	Properties: &armnetapp.SubscriptionQuotaItemProperties{
	// 		Default: to.Ptr[int32](10),
	// 		Current: to.Ptr[int32](10),
	// 	},
	// }
}
