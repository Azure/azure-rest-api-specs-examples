package armservicelinker_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/Linker.json
func ExampleLinkerClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLinkerClient().Get(ctx, "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", "linkName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LinkerResource = armservicelinker.LinkerResource{
	// 	Name: to.Ptr("linkName"),
	// 	Type: to.Ptr("Microsoft.ServiceLinker/links"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app/providers/Microsoft.ServiceLinker/links/linkName"),
	// 	SystemData: &armservicelinker.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-12T22:05:09.000Z"); return t}()),
	// 	},
	// 	Properties: &armservicelinker.LinkerProperties{
	// 		AuthInfo: &armservicelinker.SecretAuthInfo{
	// 			AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
	// 			Name: to.Ptr("name"),
	// 		},
	// 		ClientType: to.Ptr(armservicelinker.ClientTypeDotnet),
	// 		ConfigurationInfo: &armservicelinker.ConfigurationInfo{
	// 			AdditionalConfigurations: map[string]*string{
	// 				"throttlingLimit": to.Ptr("100"),
	// 			},
	// 			CustomizedKeys: map[string]*string{
	// 				"AZURE_MYSQL_CONNECTIONSTRING": to.Ptr("myConnectionstring"),
	// 				"AZURE_MYSQL_SSLMODE": to.Ptr("mySslmode"),
	// 			},
	// 			DeleteOrUpdateBehavior: to.Ptr(armservicelinker.DeleteOrUpdateBehaviorForcedCleanup),
	// 		},
	// 		PublicNetworkSolution: &armservicelinker.PublicNetworkSolution{
	// 			Action: to.Ptr(armservicelinker.ActionTypeEnable),
	// 		},
	// 		Scope: to.Ptr("AKS-Namespace"),
	// 		SecretStore: &armservicelinker.SecretStore{
	// 			KeyVaultID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.KeyVault/vaults/kvname"),
	// 		},
	// 		TargetService: &armservicelinker.AzureResource{
	// 			Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"),
	// 		},
	// 	},
	// }
}
