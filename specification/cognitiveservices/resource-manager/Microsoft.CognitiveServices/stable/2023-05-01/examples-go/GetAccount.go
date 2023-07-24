package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/GetAccount.json
func ExampleAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().Get(ctx, "myResourceGroup", "myAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armcognitiveservices.Account{
	// 	Name: to.Ptr("myAccount"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/myAccount"),
	// 	Etag: to.Ptr("W/\"datetime'2017-04-10T04%3A42%3A19.7067387Z'\""),
	// 	Kind: to.Ptr("Emotion"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcognitiveservices.AccountProperties{
	// 		Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/emotion/v1.0"),
	// 		ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armcognitiveservices.SKU{
	// 		Name: to.Ptr("F0"),
	// 	},
	// 	Tags: map[string]*string{
	// 		"ExpiredDate": to.Ptr("2017/09/01"),
	// 		"Owner": to.Ptr("felixwa"),
	// 	},
	// }
}
