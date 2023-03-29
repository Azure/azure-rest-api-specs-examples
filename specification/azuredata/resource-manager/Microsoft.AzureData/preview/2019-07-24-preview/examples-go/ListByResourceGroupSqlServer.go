package armazuredata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azuredata/resource-manager/Microsoft.AzureData/preview/2019-07-24-preview/examples/ListByResourceGroupSqlServer.json
func ExampleSQLServersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuredata.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLServersClient().NewListByResourceGroupPager("testrg", "testsqlregistration", &armazuredata.SQLServersClientListByResourceGroupOptions{Expand: nil})
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
		// page.SQLServerListResult = armazuredata.SQLServerListResult{
		// 	Value: []*armazuredata.SQLServer{
		// 		{
		// 			Name: to.Ptr("testsqlserver1"),
		// 			Type: to.Ptr("Microsoft.AzureData/SqlServerRegistrations/SqlServers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureData/SqlServerRegistrations/testsqlregistration/sqlServers/testsqlserver1"),
		// 			Properties: &armazuredata.SQLServerProperties{
		// 				Cores: to.Ptr[int32](8),
		// 				Edition: to.Ptr("Latin"),
		// 				PropertyBag: to.Ptr(""),
		// 				RegistrationID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureData/SqlServerRegistrations/testsqlregistration"),
		// 				Version: to.Ptr("2008"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testsqlserver2"),
		// 			Type: to.Ptr("Microsoft.AzureData/SqlServerRegistrations/SqlServers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureData/SqlServerRegistrations/testsqlregistration/sqlServers/testsqlserver2"),
		// 			Properties: &armazuredata.SQLServerProperties{
		// 				Cores: to.Ptr[int32](16),
		// 				Edition: to.Ptr("Latin"),
		// 				PropertyBag: to.Ptr(""),
		// 				RegistrationID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureData/SqlServerRegistrations/testsqlregistration"),
		// 				Version: to.Ptr("2008"),
		// 			},
		// 	}},
		// }
	}
}
