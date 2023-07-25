package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/ListAccountsByResourceGroup.json
func ExampleAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// 			Name: to.Ptr("myAccount"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/myAccount"),
		// 			Etag: to.Ptr("W/\"datetime'2017-04-10T04%3A42%3A19.7067387Z'\""),
		// 			Kind: to.Ptr("Emotion"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcognitiveservices.AccountProperties{
		// 				Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/emotion/v1.0"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("F0"),
		// 			},
		// 			Tags: map[string]*string{
		// 				"ExpiredDate": to.Ptr("2017/09/01"),
		// 				"Owner": to.Ptr("felixwa"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("TestPropertyWU2"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/TestPropertyWU2"),
		// 			Etag: to.Ptr("W/\"datetime'2017-04-07T04%3A32%3A38.9187216Z'\""),
		// 			Kind: to.Ptr("Face"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcognitiveservices.AccountProperties{
		// 				Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/face/v1.0"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("S0"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 	}},
		// }
	}
}
