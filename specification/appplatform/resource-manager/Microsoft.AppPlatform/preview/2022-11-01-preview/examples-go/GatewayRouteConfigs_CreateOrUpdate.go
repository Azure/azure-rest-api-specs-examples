package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/GatewayRouteConfigs_CreateOrUpdate.json
func ExampleGatewayRouteConfigsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewGatewayRouteConfigsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myservice", "default", "myRouteConfig", armappplatform.GatewayRouteConfigResource{
		Properties: &armappplatform.GatewayRouteConfigProperties{
			AppResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/apps/myApp"),
			OpenAPI: &armappplatform.GatewayRouteConfigOpenAPIProperties{
				URI: to.Ptr("https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore.json"),
			},
			Routes: []*armappplatform.GatewayAPIRoute{
				{
					Filters: []*string{
						to.Ptr("StripPrefix=2"),
						to.Ptr("RateLimit=1,1s")},
					Predicates: []*string{
						to.Ptr("Path=/api5/customer/**")},
					SsoEnabled: to.Ptr(true),
					Title:      to.Ptr("myApp route config"),
				}},
			Protocol: to.Ptr(armappplatform.GatewayRouteConfigProtocolHTTPS),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
