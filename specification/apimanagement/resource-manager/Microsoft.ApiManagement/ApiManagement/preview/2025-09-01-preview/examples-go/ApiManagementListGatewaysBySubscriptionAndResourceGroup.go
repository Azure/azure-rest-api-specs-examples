package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListGatewaysBySubscriptionAndResourceGroup.json
func ExampleAPIGatewayClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAPIGatewayClient().NewListByResourceGroupPager("rg1", nil)
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
		// page = armapimanagement.APIGatewayClientListByResourceGroupResponse{
		// 	GatewayListResult: armapimanagement.GatewayListResult{
		// 		Value: []*armapimanagement.GatewayResource{
		// 			{
		// 				Name: to.Ptr("standard-gw-1"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/gateways"),
		// 				Etag: to.Ptr(azcore.ETag("AAAAAAAWN/4=")),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/gateways/standard-gw-1"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armapimanagement.GatewayProperties{
		// 					Backend: &armapimanagement.BackendConfiguration{
		// 						Subnet: &armapimanagement.BackendSubnetConfiguration{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vn1/subnets/sn1"),
		// 						},
		// 					},
		// 					ConfigurationAPI: &armapimanagement.GatewayConfigurationAPI{
		// 						Hostname: to.Ptr("standard-gw-1.westus.configuration.gateway.azure-api.net"),
		// 					},
		// 					CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T09:40:00.9453556Z"); return t}()),
		// 					Frontend: &armapimanagement.FrontendConfiguration{
		// 						DefaultHostname: to.Ptr("standard-gw-1.westus.gateway.azure-api.net"),
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					TargetProvisioningState: to.Ptr(""),
		// 				},
		// 				SKU: &armapimanagement.GatewaySKUProperties{
		// 					Name: to.Ptr(armapimanagement.APIGatewaySKUTypeStandard),
		// 					Capacity: to.Ptr[int32](1),
		// 				},
		// 				SystemData: &armapimanagement.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T09:40:00.7106733Z"); return t}()),
		// 					CreatedBy: to.Ptr("bar@contoso.com"),
		// 					CreatedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-20T06:33:09.6159006Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("foo@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"ReleaseName": to.Ptr("Z3"),
		// 					"owner": to.Ptr("v-aswmoh"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("standard-gw-2"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/gateways"),
		// 				Etag: to.Ptr(azcore.ETag("AAAAAAAWKwo=")),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/gateways/standard-gw-2"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armapimanagement.GatewayProperties{
		// 					Backend: &armapimanagement.BackendConfiguration{
		// 						Subnet: &armapimanagement.BackendSubnetConfiguration{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vn2/subnets/sn2"),
		// 						},
		// 					},
		// 					ConfigurationAPI: &armapimanagement.GatewayConfigurationAPI{
		// 						Hostname: to.Ptr("standard-gw-2.eastus.configuration.gateway.azure-api.net"),
		// 					},
		// 					CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T09:40:00.9453556Z"); return t}()),
		// 					Frontend: &armapimanagement.FrontendConfiguration{
		// 						DefaultHostname: to.Ptr("standard-gw-2.eastus.gateway.azure-api.net"),
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					TargetProvisioningState: to.Ptr(""),
		// 				},
		// 				SKU: &armapimanagement.GatewaySKUProperties{
		// 					Name: to.Ptr(armapimanagement.APIGatewaySKUTypeStandard),
		// 					Capacity: to.Ptr[int32](1),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Owner": to.Ptr("vitaliik"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
