package armservicelinker_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/stable/2022-05-01/examples/PutLinkWithServiceEndpoint.json
func ExampleLinkerClient_BeginCreateOrUpdate_putLinkWithServiceEndpoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLinkerClient().BeginCreateOrUpdate(ctx, "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", "linkName", armservicelinker.LinkerResource{
		Properties: &armservicelinker.LinkerProperties{
			AuthInfo: &armservicelinker.SecretAuthInfo{
				AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
				Name:     to.Ptr("name"),
				SecretInfo: &armservicelinker.KeyVaultSecretURISecretInfo{
					SecretType: to.Ptr(armservicelinker.SecretTypeKeyVaultSecretURI),
					Value:      to.Ptr("https://vault-name.vault.azure.net/secrets/secret-name/00000000000000000000000000000000"),
				},
			},
			TargetService: &armservicelinker.AzureResource{
				Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
				ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DBforPostgreSQL/servers/test-pg/databases/test-db"),
			},
			VNetSolution: &armservicelinker.VNetSolution{
				Type: to.Ptr(armservicelinker.VNetSolutionTypeServiceEndpoint),
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
	// res.LinkerResource = armservicelinker.LinkerResource{
	// 	Name: to.Ptr("linkName"),
	// 	Type: to.Ptr("Microsoft.ServiceLinker/links"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app/providers/Microsoft.ServiceLinker/links/linkName"),
	// 	Properties: &armservicelinker.LinkerProperties{
	// 		AuthInfo: &armservicelinker.SecretAuthInfo{
	// 			AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
	// 			Name: to.Ptr("name"),
	// 		},
	// 		TargetService: &armservicelinker.AzureResource{
	// 			Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DBforPostgreSQL/servers/test-pg/databases/test-db"),
	// 		},
	// 		VNetSolution: &armservicelinker.VNetSolution{
	// 			Type: to.Ptr(armservicelinker.VNetSolutionTypeServiceEndpoint),
	// 		},
	// 	},
	// }
}
