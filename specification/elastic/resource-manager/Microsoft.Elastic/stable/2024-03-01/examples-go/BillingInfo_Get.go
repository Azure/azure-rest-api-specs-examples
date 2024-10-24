package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/BillingInfo_Get.json
func ExampleBillingInfoClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelastic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBillingInfoClient().Get(ctx, "myResourceGroup", "myMonitor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BillingInfoResponse = armelastic.BillingInfoResponse{
	// 	MarketplaceSaasInfo: &armelastic.MarketplaceSaaSInfo{
	// 		BilledAzureSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		MarketplaceName: to.Ptr("MP_RESOURCE"),
	// 		MarketplaceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.SaaS/resources/AzElastic_b1190c8f"),
	// 		MarketplaceStatus: to.Ptr("Status"),
	// 		MarketplaceSubscription: &armelastic.MarketplaceSaaSInfoMarketplaceSubscription{
	// 			ID: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 		},
	// 		Subscribed: to.Ptr(true),
	// 	},
	// 	PartnerBillingEntity: &armelastic.PartnerBillingEntity{
	// 		Name: to.Ptr("elasticOrganizationName"),
	// 		ID: to.Ptr("1234567890"),
	// 		PartnerEntityURI: to.Ptr("https://example.com"),
	// 	},
	// }
}
