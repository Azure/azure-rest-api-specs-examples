package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateStandardGateway.json
func ExampleAPIGatewayClient_BeginCreateOrUpdate_apiManagementCreateStandardGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAPIGatewayClient().BeginCreateOrUpdate(ctx, "rg1", "apimGateway1", armapimanagement.GatewayResource{
		Tags: map[string]*string{
			"Name": to.Ptr("Contoso"),
			"Test": to.Ptr("User"),
		},
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GatewayResource = armapimanagement.GatewayResource{
	// 	Name: to.Ptr("apimGateway1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/gateways"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/gateways/apimGateway1"),
	// 	Tags: map[string]*string{
	// 		"api-version": to.Ptr("2024-05-01"),
	// 	},
	// 	Etag: to.Ptr("AAAAAAAmREI="),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armapimanagement.GatewayProperties{
	// 		Backend: &armapimanagement.BackendConfiguration{
	// 			Subnet: &armapimanagement.BackendSubnetConfiguration{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vn1/subnets/sn1"),
	// 			},
	// 		},
	// 		ConfigurationAPI: &armapimanagement.GatewayConfigurationAPI{
	// 			Hostname: to.Ptr("apimGateway1.eastus.configuration.gateway.azure-api.net"),
	// 		},
	// 		CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-11T18:41:01.250Z"); return t}()),
	// 		Frontend: &armapimanagement.FrontendConfiguration{
	// 			DefaultHostname: to.Ptr("apimGateway1.eastus.gateway.azure-api.net"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		TargetProvisioningState: to.Ptr(""),
	// 	},
	// 	SKU: &armapimanagement.GatewaySKUProperties{
	// 		Name: to.Ptr(armapimanagement.APIGatewaySKUTypeStandard),
	// 		Capacity: to.Ptr[int32](1),
	// 	},
	// 	SystemData: &armapimanagement.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-11T18:41:00.939Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@contoso.com"),
	// 		CreatedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-11T18:41:00.939Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 	},
	// }
}
