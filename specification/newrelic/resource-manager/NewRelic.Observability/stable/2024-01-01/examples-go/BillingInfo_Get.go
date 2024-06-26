package armnewrelicobservability_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/newrelic/armnewrelicobservability"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/BillingInfo_Get.json
func ExampleBillingInfoClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnewrelicobservability.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.BillingInfoResponse = armnewrelicobservability.BillingInfoResponse{
	// 	MarketplaceSaasInfo: &armnewrelicobservability.MarketplaceSaaSInfo{
	// 		BilledAzureSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		MarketplaceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.SaaS/resources/AzNewRelic_b1190c8f"),
	// 		MarketplaceStatus: to.Ptr("Active"),
	// 		MarketplaceSubscriptionID: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 		MarketplaceSubscriptionName: to.Ptr("AzNewRelic_b1190c8f"),
	// 	},
	// 	PartnerBillingEntity: &armnewrelicobservability.PartnerBillingEntity{
	// 		OrganizationID: to.Ptr("1234567890"),
	// 		OrganizationName: to.Ptr("NROrganizationName"),
	// 	},
	// }
}
