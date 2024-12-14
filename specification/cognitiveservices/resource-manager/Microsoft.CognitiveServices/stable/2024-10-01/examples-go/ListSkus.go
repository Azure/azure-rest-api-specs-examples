package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/ListSkus.json
func ExampleAccountsClient_ListSKUs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().ListSKUs(ctx, "myResourceGroup", "myAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccountSKUListResult = armcognitiveservices.AccountSKUListResult{
	// 	Value: []*armcognitiveservices.AccountSKU{
	// 		{
	// 			ResourceType: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 			SKU: &armcognitiveservices.SKU{
	// 				Name: to.Ptr("F0"),
	// 				Tier: to.Ptr(armcognitiveservices.SKUTierFree),
	// 			},
	// 		},
	// 		{
	// 			ResourceType: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 			SKU: &armcognitiveservices.SKU{
	// 				Name: to.Ptr("S0"),
	// 				Tier: to.Ptr(armcognitiveservices.SKUTierStandard),
	// 			},
	// 	}},
	// }
}
