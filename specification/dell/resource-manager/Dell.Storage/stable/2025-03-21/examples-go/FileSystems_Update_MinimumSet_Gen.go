package armdellstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dell/armdellstorage"
)

// Generated from example definition: 2025-03-21/FileSystems_Update_MinimumSet_Gen.json
func ExampleFileSystemsClient_Update_fileSystemsUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdellstorage.NewClientFactory("BF7E7352-2FE4-4163-9CF7-5FF8EC2E9B92", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSystemsClient().Update(ctx, "rgDell", "abcd", armdellstorage.FileSystemResourceUpdate{
		Properties: &armdellstorage.FileSystemResourceUpdateProperties{
			DelegatedSubnetID: to.Ptr("uqfvajvyltgmqvdnxhbrfqbpuey"),
			Capacity: &armdellstorage.Capacity{
				Current: to.Ptr("5"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdellstorage.FileSystemsClientUpdateResponse{
	// 	FileSystemResource: &armdellstorage.FileSystemResource{
	// 		Identity: &armdellstorage.ManagedServiceIdentity{
	// 			PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 			TenantID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 			Type: to.Ptr(armdellstorage.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
	// 			UserAssignedIdentities: map[string]*armdellstorage.UserAssignedIdentity{
	// 				"key7644": &armdellstorage.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					ClientID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 				},
	// 			},
	// 		},
	// 		Properties: &armdellstorage.FileSystemResourceProperties{
	// 			Capacity: &armdellstorage.Capacity{
	// 				Min: to.Ptr("1"),
	// 				Max: to.Ptr("10"),
	// 				Incremental: to.Ptr("1"),
	// 				Current: to.Ptr("5"),
	// 			},
	// 			Marketplace: &armdellstorage.MarketplaceDetails{
	// 				PlanID: to.Ptr("lgozf"),
	// 				OfferID: to.Ptr("pzhjvibxqgeqkndqnjlduwnxqbr"),
	// 				PrivateOfferID: to.Ptr("privateOfferId"),
	// 				PlanName: to.Ptr("planeName"),
	// 				EndDate: to.Ptr("2023-05-27T17:00:00-07:00"),
	// 			},
	// 			ProvisioningState: to.Ptr(armdellstorage.ProvisioningStateAccepted),
	// 			DelegatedSubnetID: to.Ptr("yp"),
	// 			DelegatedSubnetCidr: to.Ptr("10.0.0.1/24"),
	// 			User: &armdellstorage.UserDetails{
	// 			},
	// 			FileSystemID: to.Ptr("filesystemId"),
	// 			DellReferenceNumber: to.Ptr("fhewkj"),
	// 			Encryption: &armdellstorage.EncryptionProperties{
	// 				EncryptionType: to.Ptr(armdellstorage.ResourceEncryptionTypeCustomerManagedKeysCMK),
	// 				KeyURL: to.Ptr("https://contoso.com/keyurl/keyVersion"),
	// 				EncryptionIdentityProperties: &armdellstorage.EncryptionIdentityProperties{
	// 					IdentityType: to.Ptr(armdellstorage.EncryptionIdentityTypeUserAssigned),
	// 					IdentityResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("tbcvhxzpgrijtdygkttnfswwtacs"),
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 	},
	// }
}
