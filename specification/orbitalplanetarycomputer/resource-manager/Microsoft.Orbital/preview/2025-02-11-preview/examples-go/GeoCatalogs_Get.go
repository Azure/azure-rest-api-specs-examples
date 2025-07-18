package armplanetarycomputer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/planetarycomputer/armplanetarycomputer"
)

// Generated from example definition: 2025-02-11-preview/GeoCatalogs_Get.json
func ExampleGeoCatalogsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplanetarycomputer.NewClientFactory("cd9b6cdf-dcf0-4dca-ab19-82be07b74704", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGeoCatalogsClient().Get(ctx, "MyResourceGroup", "MyCatalog", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armplanetarycomputer.GeoCatalogsClientGetResponse{
	// 	GeoCatalog: &armplanetarycomputer.GeoCatalog{
	// 		Properties: &armplanetarycomputer.GeoCatalogProperties{
	// 			Tier: to.Ptr(armplanetarycomputer.CatalogTierBasic),
	// 			CatalogURI: to.Ptr("https://mycatalog-0123456789abcdef.geocatalog.spatio.azure.com/"),
	// 			ProvisioningState: to.Ptr(armplanetarycomputer.ProvisioningStateSucceeded),
	// 			AutoGeneratedDomainNameLabelScope: to.Ptr(armplanetarycomputer.AutoGeneratedDomainNameLabelScopeNoReuse),
	// 		},
	// 		Identity: &armplanetarycomputer.ManagedServiceIdentity{
	// 			Type: to.Ptr(armplanetarycomputer.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armplanetarycomputer.UserAssignedIdentity{
	// 				"/subscriptions/cd9b6cdf-dcf0-4dca-ab19-82be07b74704/resourceGroups/MyResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/MyManagedIdentity": &armplanetarycomputer.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("1b1424a8-4d17-4a72-b47f-d1d401c1e2fd"),
	// 					ClientID: to.Ptr("c57c7b17-2c5e-4996-8169-c74d026ffc5d"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"MyTag": to.Ptr("MyValue"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/cd9b6cdf-dcf0-4dca-ab19-82be07b74704/resourceGroups/MyResourceGroup/providers/Microsoft.Orbital/geoCatalogs/MyCatalog"),
	// 		Name: to.Ptr("MyCatalog"),
	// 		Type: to.Ptr("Microsoft.Orbital/geoCatalogs"),
	// 		SystemData: &armplanetarycomputer.SystemData{
	// 			CreatedBy: to.Ptr("Catalog User"),
	// 			CreatedByType: to.Ptr(armplanetarycomputer.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-10T18:34:22.271Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("Catalog User"),
	// 			LastModifiedByType: to.Ptr(armplanetarycomputer.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-10T18:34:22.271Z"); return t}()),
	// 		},
	// 	},
	// }
}
