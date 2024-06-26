package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/685aad3f33d355c1d9c89d493ee9398865367bd8/specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/GatewayRouteConfigs_List.json
func ExampleGatewayRouteConfigsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGatewayRouteConfigsClient().NewListPager("myResourceGroup", "myservice", "default", nil)
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
		// page.GatewayRouteConfigResourceCollection = armappplatform.GatewayRouteConfigResourceCollection{
		// 	Value: []*armappplatform.GatewayRouteConfigResource{
		// 		{
		// 			Name: to.Ptr("myRouteConfig"),
		// 			Type: to.Ptr("Microsoft.AppPlatform/Spring/gateways/routeConfigs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/gateways/default/routeConfigs/myRouteConfig"),
		// 			SystemData: &armappplatform.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
		// 				CreatedBy: to.Ptr("sample-user"),
		// 				CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sample-user"),
		// 				LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armappplatform.GatewayRouteConfigProperties{
		// 				AppResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/apps/myApp"),
		// 				OpenAPI: &armappplatform.GatewayRouteConfigOpenAPIProperties{
		// 					URI: to.Ptr("https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore.json"),
		// 				},
		// 				ProvisioningState: to.Ptr(armappplatform.GatewayProvisioningStateSucceeded),
		// 				Routes: []*armappplatform.GatewayAPIRoute{
		// 					{
		// 						Filters: []*string{
		// 							to.Ptr("StripPrefix=2"),
		// 							to.Ptr("RateLimit=1,1s")},
		// 							Predicates: []*string{
		// 								to.Ptr("Path=/api5/customer/**")},
		// 								SsoEnabled: to.Ptr(true),
		// 								Title: to.Ptr("myApp route config"),
		// 						}},
		// 						Protocol: to.Ptr(armappplatform.GatewayRouteConfigProtocolHTTPS),
		// 					},
		// 			}},
		// 		}
	}
}
