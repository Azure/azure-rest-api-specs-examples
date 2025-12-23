package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/CustomIpPrefixListAll.json
func ExampleCustomIPPrefixesClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomIPPrefixesClient().NewListAllPager(nil)
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
		// page.CustomIPPrefixListResult = armnetwork.CustomIPPrefixListResult{
		// 	Value: []*armnetwork.CustomIPPrefix{
		// 		{
		// 			Name: to.Ptr("test-customipprefix"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/customIpPrefixes/test-customipprefix"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.0.0/24"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioned),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix2"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/customIpPrefixes/test-customipprefix2"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.2.0/23"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioned),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix3"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix3"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.4.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateProvisioning),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix4"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix4"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("2607:f0d1:1002:0001::/64"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioned),
		// 				CustomIPPrefixParent: &armnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix5"),
		// 				},
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix5"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix5"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix4"),
		// 				}},
		// 				Cidr: to.Ptr("2607:f0d1:1002::/48"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateProvisioned),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix6"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.5.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedState("ProvisionFailed")),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr("CustomerSignatureNotVerified"),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix7"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.6.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioning),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix8"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.7.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioning),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix8"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				Asn: to.Ptr("11"),
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.7.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioning),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix8"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.7.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioning),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				Geo: to.Ptr(armnetwork.GeoGLOBAL),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix8"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.7.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioning),
		// 				ExpressRouteAdvertise: to.Ptr(true),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix8"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.7.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioning),
		// 				ExpressRouteAdvertise: to.Ptr(true),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				PrefixType: to.Ptr(armnetwork.CustomIPPrefixTypeParent),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 	}},
		// }
	}
}
