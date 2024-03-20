package armservicelinker_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c16ce913afbdaa073e8ca5e480f3b465db2de542/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2023-04-01-preview/examples/ListDryrun.json
func ExampleLinkersClient_NewListDryrunPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLinkersClient().NewListDryrunPager("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", nil)
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
		// page.DryrunList = armservicelinker.DryrunList{
		// 	Value: []*armservicelinker.DryrunResource{
		// 		{
		// 			Name: to.Ptr("dryrunName"),
		// 			Type: to.Ptr("Microsoft.ServiceLinker/dryruns"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app/providers/Microsoft.ServiceLinker/dryruns/dryrunName"),
		// 			SystemData: &armservicelinker.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-12T22:05:09.000Z"); return t}()),
		// 			},
		// 			Properties: &armservicelinker.DryrunProperties{
		// 				Parameters: &armservicelinker.CreateOrUpdateDryrunParameters{
		// 					ActionName: to.Ptr(armservicelinker.DryrunActionNameCreateOrUpdate),
		// 					AuthInfo: &armservicelinker.SecretAuthInfo{
		// 						AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
		// 						Name: to.Ptr("username"),
		// 					},
		// 					TargetService: &armservicelinker.AzureResource{
		// 						Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
