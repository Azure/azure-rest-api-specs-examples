package armservicelinker_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c16ce913afbdaa073e8ca5e480f3b465db2de542/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2023-04-01-preview/examples/PutLinker.json
func ExampleLinkerClient_BeginCreateOrUpdate() {
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
				SecretInfo: &armservicelinker.ValueSecretInfo{
					SecretType: to.Ptr(armservicelinker.SecretTypeRawValue),
					Value:      to.Ptr("secret"),
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
