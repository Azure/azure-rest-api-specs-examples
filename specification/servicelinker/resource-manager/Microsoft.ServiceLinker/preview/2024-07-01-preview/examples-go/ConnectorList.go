package armservicelinker_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/ConnectorList.json
func ExampleConnectorClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectorClient().NewListPager("00000000-0000-0000-0000-000000000000", "test-rg", "westus", nil)
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
		// page.ResourceList = armservicelinker.ResourceList{
		// 	Value: []*armservicelinker.LinkerResource{
		// 		{
		// 			Name: to.Ptr("linkName"),
		// 			Type: to.Ptr("Microsoft.ServiceLinker/devConnectors"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.ServiceLinker/linkers/linkName"),
		// 			SystemData: &armservicelinker.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-12T22:05:09.000Z"); return t}()),
		// 			},
		// 			Properties: &armservicelinker.LinkerProperties{
		// 				AuthInfo: &armservicelinker.SecretAuthInfo{
		// 					AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
		// 					Name: to.Ptr("username"),
		// 				},
		// 				TargetService: &armservicelinker.AzureResource{
		// 					Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
		// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
