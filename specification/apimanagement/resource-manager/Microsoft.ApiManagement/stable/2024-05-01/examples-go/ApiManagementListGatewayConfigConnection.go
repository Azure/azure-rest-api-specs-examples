package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListGatewayConfigConnection.json
func ExampleAPIGatewayConfigConnectionClient_NewListByGatewayPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAPIGatewayConfigConnectionClient().NewListByGatewayPager("rg1", "standard-gw-1", nil)
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
		// page.GatewayConfigConnectionListResult = armapimanagement.GatewayConfigConnectionListResult{
		// 	Value: []*armapimanagement.GatewayConfigConnectionResource{
		// 		{
		// 			Name: to.Ptr("gcc-1"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/gateways/configConnections"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/gateways/standard-gw-1/configConnections/gcc-1"),
		// 			Etag: to.Ptr("AAAAAAAWN/4="),
		// 			Properties: &armapimanagement.GatewayConfigConnectionBaseProperties{
		// 				DefaultHostname: to.Ptr("gcc-1-ahg4t5iof8gaafex.standard-gw-1.gateway.eastus.azure-api.net"),
		// 				Hostnames: []*string{
		// 					to.Ptr("gcc1standard-gw-1.gateway.eastus.azure-api.net")},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					SourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/services/apim-service-1/workspaces/ws-001"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("gcc-2"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/gateways/configConnetions"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/gateways/standard-gw-1/configConnections/gcc-2"),
		// 				Etag: to.Ptr("AAAAAAAWKwo="),
		// 				Properties: &armapimanagement.GatewayConfigConnectionBaseProperties{
		// 					DefaultHostname: to.Ptr("gcc-2-ahg4t5iof8gaafex.standard-gw-1.gateway.eastus.azure-api.net"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					SourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/services/apim-service-1/workspaces/ws-002"),
		// 				},
		// 		}},
		// 	}
	}
}
