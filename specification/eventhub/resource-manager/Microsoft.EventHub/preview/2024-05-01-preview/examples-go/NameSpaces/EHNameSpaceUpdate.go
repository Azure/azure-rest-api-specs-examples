package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5759c77eee2d57bdb9e47aa1805d0ffb61704f2d/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2024-05-01-preview/examples/NameSpaces/EHNameSpaceUpdate.json
func ExampleNamespacesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().Update(ctx, "ResurceGroupSample", "NamespaceSample", armeventhub.EHNamespace{
		Location: to.Ptr("East US"),
		Identity: &armeventhub.Identity{
			Type: to.Ptr(armeventhub.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armeventhub.UserAssignedIdentity{
				"/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2": nil,
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EHNamespace = armeventhub.EHNamespace{
	// 	Name: to.Ptr("NamespaceSample"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces"),
	// 	ID: to.Ptr("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.EventHub/namespaces/NamespaceSample"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armeventhub.Identity{
	// 		Type: to.Ptr(armeventhub.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
	// 		PrincipalID: to.Ptr("PrincipalIdGUID"),
	// 		TenantID: to.Ptr("TenantIdGUID"),
	// 		UserAssignedIdentities: map[string]*armeventhub.UserAssignedIdentity{
	// 			"/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud1": &armeventhub.UserAssignedIdentity{
	// 				ClientID: to.Ptr("ClientIdGUID"),
	// 				PrincipalID: to.Ptr("PrincipalIdGUID"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armeventhub.EHNamespaceProperties{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-30T00:28:38.963Z"); return t}()),
	// 		DisableLocalAuth: to.Ptr(false),
	// 		Encryption: &armeventhub.Encryption{
	// 			KeySource: to.Ptr("Microsoft.KeyVault"),
	// 			KeyVaultProperties: []*armeventhub.KeyVaultProperties{
	// 				{
	// 					Identity: &armeventhub.UserAssignedIdentityProperties{
	// 						UserAssignedIdentity: to.Ptr("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud1"),
	// 					},
	// 					KeyName: to.Ptr("Samplekey"),
	// 					KeyVaultURI: to.Ptr("https://sample-keyvault-user.vault-int.azure-int.net"),
	// 					KeyVersion: to.Ptr(""),
	// 			}},
	// 			RequireInfrastructureEncryption: to.Ptr(false),
	// 		},
	// 		IsAutoInflateEnabled: to.Ptr(false),
	// 		MaximumThroughputUnits: to.Ptr[int32](0),
	// 		MetricID: to.Ptr("MetricGUID:NamespaceSample"),
	// 		MinimumTLSVersion: to.Ptr(armeventhub.TLSVersionOne1),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServiceBusEndpoint: to.Ptr("https://NamespaceSample.servicebus.windows-int.net:443/"),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-30T00:31:13.657Z"); return t}()),
	// 		ZoneRedundant: to.Ptr(false),
	// 	},
	// 	SKU: &armeventhub.SKU{
	// 		Name: to.Ptr(armeventhub.SKUNameStandard),
	// 		Capacity: to.Ptr[int32](1),
	// 		Tier: to.Ptr(armeventhub.SKUTierStandard),
	// 	},
	// }
}
