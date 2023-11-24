package armresourceconnector_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourceconnector/armresourceconnector"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5f700acd3d094d8eedca381932f2e7916afd2e55/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesCreate_Update.json
func ExampleAppliancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAppliancesClient().BeginCreateOrUpdate(ctx, "testresourcegroup", "appliance01", armresourceconnector.Appliance{
		Location: to.Ptr("West US"),
		Properties: &armresourceconnector.ApplianceProperties{
			Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
			InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
				Provider: to.Ptr(armresourceconnector.ProviderVMWare),
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
	// res.Appliance = armresourceconnector.Appliance{
	// 	Name: to.Ptr("appliance01"),
	// 	Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.ResourceConnector/appliances/appliance01"),
	// 	SystemData: &armresourceconnector.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.092Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.092Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("West US"),
	// 	Identity: &armresourceconnector.Identity{
	// 		Type: to.Ptr(armresourceconnector.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		TenantID: to.Ptr("111111-1111-1111-1111-111111111111"),
	// 	},
	// 	Properties: &armresourceconnector.ApplianceProperties{
	// 		Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
	// 		InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
	// 			Provider: to.Ptr(armresourceconnector.ProviderVMWare),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Status: to.Ptr(armresourceconnector.StatusConnected),
	// 		Version: to.Ptr("1.0.1"),
	// 	},
	// }
}
