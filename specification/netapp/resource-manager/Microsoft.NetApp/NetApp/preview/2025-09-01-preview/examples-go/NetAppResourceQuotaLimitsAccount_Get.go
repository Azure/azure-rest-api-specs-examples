package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/NetAppResourceQuotaLimitsAccount_Get.json
func ExampleResourceQuotaLimitsAccountClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("D633CC2E-722B-4AE1-B636-BBD9E4C60ED9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceQuotaLimitsAccountClient().Get(ctx, "myRG", "myAccount", "poolsPerAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.ResourceQuotaLimitsAccountClientGetResponse{
	// 	SubscriptionQuotaItem: &armnetapp.SubscriptionQuotaItem{
	// 		ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/myAccount/quotaLimits/poolsPerAccount"),
	// 		Name: to.Ptr("myAccount/poolsPerAccount"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts/quotaLimits"),
	// 		Properties: &armnetapp.SubscriptionQuotaItemProperties{
	// 			Current: to.Ptr[int32](10),
	// 			Default: to.Ptr[int32](10),
	// 			Usage: to.Ptr[int32](10),
	// 		},
	// 	},
	// }
}
