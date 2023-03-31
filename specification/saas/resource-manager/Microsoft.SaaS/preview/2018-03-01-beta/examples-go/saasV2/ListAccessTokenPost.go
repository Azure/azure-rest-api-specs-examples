package armsaas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/saas/resource-manager/Microsoft.SaaS/preview/2018-03-01-beta/examples/saasV2/ListAccessTokenPost.json
func ExampleResourcesClient_ListAccessToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsaas.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourcesClient().ListAccessToken(ctx, "c825645b-e31b-9cf4-1cee-2aba9e58bc7c", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessTokenResult = armsaas.AccessTokenResult{
	// 	PublisherOfferBaseURI: to.Ptr("https://contoso.website.com/api/v1/marketplace/azure/common/auth"),
	// 	Token: to.Ptr("jyhAtr3AiFTXO0QsEkYizsNzqIAUJ+E0M1SXhl4E5hwAl7+GJT6t+dEwuWpSizgR0Vn6dMhzpS94JRzkUh2Xuq5L1QDgmDkDUeIikTFNviwtwxtF8CwipcB49rq5d8whmVp0CmWEjd/FzU0vvlikBRuy+asxC1UhXv6XxBkSxkZKj29AAFiGgsgRvXIld47C"),
	// }
}
