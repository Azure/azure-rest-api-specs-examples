package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Put.json
func ExampleIntegrationServiceEnvironmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armlogic.NewIntegrationServiceEnvironmentsClient("f34b22a3-2202-4fb1-b040-1332bd928c84", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testResourceGroup",
		"testIntegrationServiceEnvironment",
		armlogic.IntegrationServiceEnvironment{
			Location: to.Ptr("brazilsouth"),
			Properties: &armlogic.IntegrationServiceEnvironmentProperties{
				EncryptionConfiguration: &armlogic.IntegrationServiceEnvironmenEncryptionConfiguration{
					EncryptionKeyReference: &armlogic.IntegrationServiceEnvironmenEncryptionKeyReference{
						KeyName: to.Ptr("testKeyName"),
						KeyVault: &armlogic.ResourceReference{
							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.KeyVault/vaults/testKeyVault"),
						},
						KeyVersion: to.Ptr("13b261d30b984753869902d7f47f4d55"),
					},
				},
				NetworkConfiguration: &armlogic.NetworkConfiguration{
					AccessEndpoint: &armlogic.IntegrationServiceEnvironmentAccessEndpoint{
						Type: to.Ptr(armlogic.IntegrationServiceEnvironmentAccessEndpointTypeInternal),
					},
					Subnets: []*armlogic.ResourceReference{
						{
							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s1"),
						},
						{
							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s2"),
						},
						{
							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s3"),
						},
						{
							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s4"),
						}},
				},
			},
			SKU: &armlogic.IntegrationServiceEnvironmentSKU{
				Name:     to.Ptr(armlogic.IntegrationServiceEnvironmentSKUNamePremium),
				Capacity: to.Ptr[int32](2),
			},
		},
		nil)
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
