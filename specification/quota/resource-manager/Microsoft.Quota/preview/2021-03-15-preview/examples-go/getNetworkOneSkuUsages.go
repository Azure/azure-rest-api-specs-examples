package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/quota/resource-manager/Microsoft.Quota/preview/2021-03-15-preview/examples/getNetworkOneSkuUsages.json
func ExampleUsagesClient_Get_quotasUsagesRequestForNetwork() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUsagesClient().Get(ctx, "MinPublicIpInterNetworkPrefixLength", "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CurrentUsagesBase = armquota.CurrentUsagesBase{
	// 	Name: to.Ptr("MinPublicIpInterNetworkPrefixLength"),
	// 	Type: to.Ptr("Microsoft.Quota/usages"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/MinPublicIpInterNetworkPrefixLength"),
	// 	Properties: &armquota.UsagesProperties{
	// 		Name: &armquota.ResourceName{
	// 			LocalizedValue: to.Ptr("Min Public Ip InterNetwork Prefix Length"),
	// 			Value: to.Ptr("MinPublicIpInterNetworkPrefixLength"),
	// 		},
	// 		IsQuotaApplicable: to.Ptr(true),
	// 		ResourceType: to.Ptr("MinPublicIpInterNetworkPrefixLength"),
	// 		Unit: to.Ptr("Count"),
	// 		Usages: &armquota.UsagesObject{
	// 			Value: to.Ptr[int32](10),
	// 		},
	// 	},
	// }
}
