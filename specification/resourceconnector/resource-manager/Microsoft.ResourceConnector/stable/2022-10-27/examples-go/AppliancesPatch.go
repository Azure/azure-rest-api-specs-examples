package armresourceconnector_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourceconnector/armresourceconnector"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesPatch.json
func ExampleAppliancesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppliancesClient().Update(ctx, "testresourcegroup", "appliance01", armresourceconnector.PatchableAppliance{
		Tags: map[string]*string{
			"key": to.Ptr("value"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Appliance = armresourceconnector.Appliance{
	// 	Name: to.Ptr("appliance01"),
	// 	Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ResourceConnector/appliances/appliance01"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
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
	// 		Status: to.Ptr(armresourceconnector.StatusValidating),
	// 		Version: to.Ptr("1.0.1"),
	// 	},
	// }
}
