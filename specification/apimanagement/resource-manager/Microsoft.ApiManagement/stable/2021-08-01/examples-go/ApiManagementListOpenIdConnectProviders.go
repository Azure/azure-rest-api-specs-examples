package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListOpenIdConnectProviders.json
func ExampleOpenIDConnectProviderClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOpenIDConnectProviderClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.OpenIDConnectProviderClientListByServiceOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
	})
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
		// page.OpenIDConnectProviderCollection = armapimanagement.OpenIDConnectProviderCollection{
		// 	Count: to.Ptr[int64](1),
		// 	Value: []*armapimanagement.OpenidConnectProviderContract{
		// 		{
		// 			Name: to.Ptr("templateOpenIdConnect2"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/openidconnectproviders"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/openidConnectProviders/templateOpenIdConnect2"),
		// 			Properties: &armapimanagement.OpenidConnectProviderContractProperties{
		// 				Description: to.Ptr("open id provider template2"),
		// 				ClientID: to.Ptr("oidprovidertemplate2"),
		// 				DisplayName: to.Ptr("templateoidprovider2"),
		// 				MetadataEndpoint: to.Ptr("https://oidprovider-template2.net"),
		// 			},
		// 	}},
		// }
	}
}
