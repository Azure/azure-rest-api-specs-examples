package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/523ccabf440d8cf1c5b0ea18a8ad1ffedf4902ac/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/RegistryCreateZoneRedundant.json
func ExampleRegistriesClient_BeginCreate_registryCreateZoneRedundant() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRegistriesClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", armcontainerregistry.Registry{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"key": to.Ptr("value"),
		},
		Properties: &armcontainerregistry.RegistryProperties{
			ZoneRedundancy: to.Ptr(armcontainerregistry.ZoneRedundancyEnabled),
		},
		SKU: &armcontainerregistry.SKU{
			Name: to.Ptr(armcontainerregistry.SKUNameStandard),
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
	// res.Registry = armcontainerregistry.Registry{
	// 	Name: to.Ptr("myRegistry"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armcontainerregistry.RegistryProperties{
	// 		AdminUserEnabled: to.Ptr(false),
	// 		AnonymousPullEnabled: to.Ptr(false),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-15T21:38:26.153Z"); return t}()),
	// 		DataEndpointEnabled: to.Ptr(false),
	// 		DataEndpointHostNames: []*string{
	// 		},
	// 		Encryption: &armcontainerregistry.EncryptionProperty{
	// 			Status: to.Ptr(armcontainerregistry.EncryptionStatusDisabled),
	// 		},
	// 		LoginServer: to.Ptr("myRegistry.azurecr.io"),
	// 		NetworkRuleBypassAllowedForTasks: to.Ptr(false),
	// 		NetworkRuleBypassOptions: to.Ptr(armcontainerregistry.NetworkRuleBypassOptionsAzureServices),
	// 		NetworkRuleSet: &armcontainerregistry.NetworkRuleSet{
	// 			DefaultAction: to.Ptr(armcontainerregistry.DefaultActionAllow),
	// 			IPRules: []*armcontainerregistry.IPRule{
	// 			},
	// 		},
	// 		Policies: &armcontainerregistry.Policies{
	// 			AzureADAuthenticationAsArmPolicy: &armcontainerregistry.AzureADAuthenticationAsArmPolicy{
	// 				Status: to.Ptr(armcontainerregistry.AzureADAuthenticationAsArmPolicyStatusEnabled),
	// 			},
	// 			ExportPolicy: &armcontainerregistry.ExportPolicy{
	// 				Status: to.Ptr(armcontainerregistry.ExportPolicyStatusEnabled),
	// 			},
	// 			QuarantinePolicy: &armcontainerregistry.QuarantinePolicy{
	// 				Status: to.Ptr(armcontainerregistry.PolicyStatusDisabled),
	// 			},
	// 			RetentionPolicy: &armcontainerregistry.RetentionPolicy{
	// 				Days: to.Ptr[int32](7),
	// 				LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-15T21:40:12.850Z"); return t}()),
	// 				Status: to.Ptr(armcontainerregistry.PolicyStatusDisabled),
	// 			},
	// 			TrustPolicy: &armcontainerregistry.TrustPolicy{
	// 				Type: to.Ptr(armcontainerregistry.TrustPolicyTypeNotary),
	// 				Status: to.Ptr(armcontainerregistry.PolicyStatusDisabled),
	// 			},
	// 		},
	// 		PrivateEndpointConnections: []*armcontainerregistry.PrivateEndpointConnection{
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armcontainerregistry.PublicNetworkAccessEnabled),
	// 		RoleAssignmentMode: to.Ptr(armcontainerregistry.RoleAssignmentModeLegacyRegistryPermissions),
	// 		ZoneRedundancy: to.Ptr(armcontainerregistry.ZoneRedundancyEnabled),
	// 	},
	// 	SKU: &armcontainerregistry.SKU{
	// 		Name: to.Ptr(armcontainerregistry.SKUNameStandard),
	// 		Tier: to.Ptr(armcontainerregistry.SKUTierStandard),
	// 	},
	// }
}
