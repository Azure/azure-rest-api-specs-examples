package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAssemblies_List.json
func ExampleIntegrationAccountAssembliesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIntegrationAccountAssembliesClient().NewListPager("testResourceGroup", "testIntegrationAccount", nil)
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
		// page.AssemblyCollection = armlogic.AssemblyCollection{
		// 	Value: []*armlogic.AssemblyDefinition{
		// 		{
		// 			Name: to.Ptr("testAssembly"),
		// 			Type: to.Ptr("Microsoft.Logic/integrationAccounts/assemblies"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationAccounts/testIntegrationAccount/assemblies/testAssembly"),
		// 			Properties: &armlogic.AssemblyProperties{
		// 				ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-24T18:45:23.413Z"); return t}()),
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-24T18:45:23.412Z"); return t}()),
		// 				Metadata: map[string]any{
		// 				},
		// 				ContentLink: &armlogic.ContentLink{
		// 					ContentHash: &armlogic.ContentHash{
		// 						Algorithm: to.Ptr("md5"),
		// 						Value: to.Ptr("A2avz/M0ov2FPI3+Je8vDw=="),
		// 					},
		// 					ContentSize: to.Ptr[int64](3056),
		// 					ContentVersion: to.Ptr("\"0x8D45CE54B058881\""),
		// 					URI: to.Ptr("<Uri>"),
		// 				},
		// 				AssemblyName: to.Ptr("System.IdentityModel.Tokens.Jwt"),
		// 				AssemblyVersion: to.Ptr("0.0.0.0"),
		// 			},
		// 	}},
		// }
	}
}
