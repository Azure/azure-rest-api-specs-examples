package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5759c77eee2d57bdb9e47aa1805d0ffb61704f2d/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2024-05-01-preview/examples/NameSpaces/EHNameSpaceListByResourceGroup.json
func ExampleNamespacesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespacesClient().NewListByResourceGroupPager("ResurceGroupSample", nil)
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
		// page.EHNamespaceListResult = armeventhub.EHNamespaceListResult{
		// 	Value: []*armeventhub.EHNamespace{
		// 		{
		// 			Name: to.Ptr("NamespaceSample"),
		// 			Type: to.Ptr("Microsoft.EventHub/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.EventHub/namespaces/NamespaceSample"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armeventhub.Identity{
		// 				Type: to.Ptr(armeventhub.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
		// 				PrincipalID: to.Ptr("PrincipalIdGUID"),
		// 				TenantID: to.Ptr("TenantIdGUID"),
		// 				UserAssignedIdentities: map[string]*armeventhub.UserAssignedIdentity{
		// 					"/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud1": &armeventhub.UserAssignedIdentity{
		// 						ClientID: to.Ptr("ClientIdGUID"),
		// 						PrincipalID: to.Ptr("PrincipalIdGUID"),
		// 					},
		// 					"/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2": &armeventhub.UserAssignedIdentity{
		// 						ClientID: to.Ptr("6a35400f-6ccb-4817-8f1a-ce19ea4523bc"),
		// 						PrincipalID: to.Ptr("ce2d5953-5c15-40ca-9d51-cc3f4a63b0f5"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armeventhub.EHNamespaceProperties{
		// 				ClusterArmID: to.Ptr("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.EventHub/clusters/enc-test"),
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-30T00:28:38.963Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				Encryption: &armeventhub.Encryption{
		// 					KeySource: to.Ptr("Microsoft.KeyVault"),
		// 					KeyVaultProperties: []*armeventhub.KeyVaultProperties{
		// 						{
		// 							Identity: &armeventhub.UserAssignedIdentityProperties{
		// 								UserAssignedIdentity: to.Ptr("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud1"),
		// 							},
		// 							KeyName: to.Ptr("Samplekey"),
		// 							KeyVaultURI: to.Ptr("https://sample-keyvault-user.vault-int.azure-int.net"),
		// 							KeyVersion: to.Ptr(""),
		// 					}},
		// 					RequireInfrastructureEncryption: to.Ptr(false),
		// 				},
		// 				IsAutoInflateEnabled: to.Ptr(false),
		// 				MaximumThroughputUnits: to.Ptr[int32](0),
		// 				MetricID: to.Ptr("MetricGUID:NamespaceSample"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NamespaceSample.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-30T00:30:55.143Z"); return t}()),
		// 				ZoneRedundant: to.Ptr(false),
		// 			},
		// 			SKU: &armeventhub.SKU{
		// 				Name: to.Ptr(armeventhub.SKUNameStandard),
		// 				Capacity: to.Ptr[int32](1),
		// 				Tier: to.Ptr(armeventhub.SKUTierStandard),
		// 			},
		// 	}},
		// }
	}
}
