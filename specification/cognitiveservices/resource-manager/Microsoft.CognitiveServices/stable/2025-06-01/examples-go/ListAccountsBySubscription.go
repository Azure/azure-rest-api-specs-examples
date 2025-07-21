package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/ListAccountsBySubscription.json
func ExampleAccountsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AccountListResult = armcognitiveservices.AccountListResult{
		// 	Value: []*armcognitiveservices.Account{
		// 		{
		// 			Name: to.Ptr("bingSearch"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/bvttest/providers/Microsoft.CognitiveServices/accounts/bingSearch"),
		// 			Etag: to.Ptr("W/\"datetime'2017-03-27T11%3A19%3A08.762494Z'\""),
		// 			Kind: to.Ptr("Bing.Search"),
		// 			Location: to.Ptr("global"),
		// 			Properties: &armcognitiveservices.AccountProperties{
		// 				Endpoint: to.Ptr("https://api.cognitive.microsoft.com/bing/v5.0"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("S1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("CrisProd"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/bvttest/providers/Microsoft.CognitiveServices/accounts/CrisProd"),
		// 			Etag: to.Ptr("W/\"datetime'2017-03-31T08%3A57%3A07.4499566Z'\""),
		// 			Kind: to.Ptr("CRIS"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcognitiveservices.AccountProperties{
		// 				Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/sts/v1.0"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("S0"),
		// 			},
		// 			Tags: map[string]*string{
		// 				"can't delete it successfully": to.Ptr("v-yunjin"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("rayrptest0308"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/bvttest/providers/Microsoft.CognitiveServices/accounts/rayrptest0308"),
		// 			Etag: to.Ptr("W/\"datetime'2017-03-27T11%3A15%3A23.5232645Z'\""),
		// 			Kind: to.Ptr("Face"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcognitiveservices.AccountProperties{
		// 				Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/face/v1.0"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("S0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("raytest02"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/bvttest/providers/Microsoft.CognitiveServices/accounts/raytest02"),
		// 			Etag: to.Ptr("W/\"datetime'2017-04-04T02%3A07%3A07.3957572Z'\""),
		// 			Kind: to.Ptr("Emotion"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcognitiveservices.AccountProperties{
		// 				Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/emotion/v1.0"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("S0"),
		// 			},
		// 	}},
		// }
	}
}
