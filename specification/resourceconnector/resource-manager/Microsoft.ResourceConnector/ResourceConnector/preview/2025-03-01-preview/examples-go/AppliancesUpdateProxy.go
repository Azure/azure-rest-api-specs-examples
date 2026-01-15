package armresourceconnector_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourceconnector/armresourceconnector"
)

// Generated from example definition: 2025-03-01-preview/AppliancesUpdateProxy.json
func ExampleAppliancesClient_BeginCreateOrUpdate_updateApplianceProxyConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAppliancesClient().BeginCreateOrUpdate(ctx, "testresourcegroup", "appliance01", armresourceconnector.Appliance{
		Location: to.Ptr("West US"),
		Properties: &armresourceconnector.ApplianceProperties{
			Distro:    to.Ptr(armresourceconnector.DistroAKSEdge),
			PublicKey: to.Ptr("xxxxxxxx"),
			InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
				Provider: to.Ptr(armresourceconnector.ProviderVMWare),
			},
			NetworkProfile: &armresourceconnector.NetworkProfile{
				ProxyConfiguration: &armresourceconnector.ProxyConfiguration{
					Version: to.Ptr("latest"),
				},
			},
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
	// res = armresourceconnector.AppliancesClientCreateOrUpdateResponse{
	// 	Appliance: &armresourceconnector.Appliance{
	// 		ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.ResourceConnector/appliances/appliance01"),
	// 		Name: to.Ptr("appliance01"),
	// 		Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armresourceconnector.ApplianceProperties{
	// 			Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
	// 			Version: to.Ptr("1.0.1"),
	// 			ProvisioningState: to.Ptr("Updating"),
	// 			Status: to.Ptr(armresourceconnector.StatusNetworkProxyUpdatePreparing),
	// 			PublicKey: to.Ptr("xxxxxxxx"),
	// 			InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
	// 				Provider: to.Ptr(armresourceconnector.ProviderVMWare),
	// 			},
	// 			NetworkProfile: &armresourceconnector.NetworkProfile{
	// 				ProxyConfiguration: &armresourceconnector.ProxyConfiguration{
	// 					Version: to.Ptr("1750464301"),
	// 				},
	// 				DNSConfiguration: &armresourceconnector.DNSConfiguration{
	// 					Version: to.Ptr("1744221145"),
	// 				},
	// 				GatewayConfiguration: &armresourceconnector.GatewayConfiguration{
	// 					Version: to.Ptr("1744221145"),
	// 				},
	// 			},
	// 		},
	// 		Identity: &armresourceconnector.Identity{
	// 			Type: to.Ptr(armresourceconnector.ResourceIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 			TenantID: to.Ptr("111111-1111-1111-1111-111111111111"),
	// 		},
	// 		SystemData: &armresourceconnector.SystemData{
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-09T17:52:25.0928001Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-21T00:05:15.0000000Z"); return t}()),
	// 		},
	// 	},
	// }
}
