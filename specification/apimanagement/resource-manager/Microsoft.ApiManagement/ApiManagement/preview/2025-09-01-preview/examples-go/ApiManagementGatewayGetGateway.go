package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementGatewayGetGateway.json
func ExampleAPIGatewayClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIGatewayClient().Get(ctx, "rg1", "apimService1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.APIGatewayClientGetResponse{
	// 	GatewayResource: armapimanagement.GatewayResource{
	// 		Name: to.Ptr("example-gateway"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/gateway"),
	// 		Etag: to.Ptr(azcore.ETag("AAAAAAAWN/4=")),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/gateway/example-gateway"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armapimanagement.GatewayProperties{
	// 			Backend: &armapimanagement.BackendConfiguration{
	// 				Subnet: &armapimanagement.BackendSubnetConfiguration{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vn1/subnets/sn1"),
	// 				},
	// 			},
	// 			ConfigurationAPI: &armapimanagement.GatewayConfigurationAPI{
	// 				Hostname: to.Ptr("example-gateway.eastus.configuration.gateway.azure-api.net"),
	// 			},
	// 			CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T09:40:00.9453556Z"); return t}()),
	// 			Frontend: &armapimanagement.FrontendConfiguration{
	// 				DefaultHostname: to.Ptr("example-gateway.eastus.gateway.azure-api.net"),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			TargetProvisioningState: to.Ptr(""),
	// 		},
	// 		SKU: &armapimanagement.GatewaySKUProperties{
	// 			Name: to.Ptr(armapimanagement.APIGatewaySKUTypeStandard),
	// 			Capacity: to.Ptr[int32](1),
	// 		},
	// 		SystemData: &armapimanagement.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T09:40:00.7106733Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-20T06:33:09.6159006Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("foo@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"ReleaseName": to.Ptr("Z3"),
	// 			"owner": to.Ptr("v-aswmoh"),
	// 		},
	// 	},
	// }
}
