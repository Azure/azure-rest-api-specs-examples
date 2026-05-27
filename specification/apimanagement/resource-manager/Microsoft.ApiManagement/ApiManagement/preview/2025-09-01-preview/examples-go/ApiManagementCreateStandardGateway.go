package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementCreateStandardGateway.json
func ExampleAPIGatewayClient_BeginCreateOrUpdate_apiManagementCreateStandardGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAPIGatewayClient().BeginCreateOrUpdate(ctx, "rg1", "apimGateway1", armapimanagement.GatewayResource{
		Location: to.Ptr("South Central US"),
		Properties: &armapimanagement.GatewayProperties{
			Backend: &armapimanagement.BackendConfiguration{
				Subnet: &armapimanagement.BackendSubnetConfiguration{
					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vn1/subnets/sn1"),
				},
			},
		},
		SKU: &armapimanagement.GatewaySKUProperties{
			Name:     to.Ptr(armapimanagement.APIGatewaySKUTypeStandard),
			Capacity: to.Ptr[int32](1),
		},
		Tags: map[string]*string{
			"Name": to.Ptr("Contoso"),
			"Test": to.Ptr("User"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.APIGatewayClientCreateOrUpdateResponse{
	// 	GatewayResource: armapimanagement.GatewayResource{
	// 		Name: to.Ptr("apimGateway1"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/gateways"),
	// 		Etag: to.Ptr(azcore.ETag("AAAAAAAmREI=")),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/gateways/apimGateway1"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armapimanagement.GatewayProperties{
	// 			Backend: &armapimanagement.BackendConfiguration{
	// 				Subnet: &armapimanagement.BackendSubnetConfiguration{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vn1/subnets/sn1"),
	// 				},
	// 			},
	// 			ConfigurationAPI: &armapimanagement.GatewayConfigurationAPI{
	// 				Hostname: to.Ptr("apimGateway1.eastus.configuration.gateway.azure-api.net"),
	// 			},
	// 			CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-11T18:41:01.2506031Z"); return t}()),
	// 			Frontend: &armapimanagement.FrontendConfiguration{
	// 				DefaultHostname: to.Ptr("apimGateway1.eastus.gateway.azure-api.net"),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			TargetProvisioningState: to.Ptr(""),
	// 		},
	// 		SKU: &armapimanagement.GatewaySKUProperties{
	// 			Name: to.Ptr(armapimanagement.APIGatewaySKUTypeStandard),
	// 			Capacity: to.Ptr[int32](1),
	// 		},
	// 		SystemData: &armapimanagement.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-11T18:41:00.9390609Z"); return t}()),
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-11T18:41:00.9390609Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"api-version": to.Ptr("2025-09-01-preview"),
	// 		},
	// 	},
	// }
}
