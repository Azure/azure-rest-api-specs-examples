package armservicelinker_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/GetDryrun.json
func ExampleLinkersClient_GetDryrun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLinkersClient().GetDryrun(ctx, "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", "dryrunName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DryrunResource = armservicelinker.DryrunResource{
	// 	Name: to.Ptr("dryrunName"),
	// 	Type: to.Ptr("Microsoft.ServiceLinker/dryruns"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app/providers/Microsoft.ServiceLinker/dryruns/dryrunName"),
	// 	SystemData: &armservicelinker.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-12T22:05:09.000Z"); return t}()),
	// 	},
	// 	Properties: &armservicelinker.DryrunProperties{
	// 		Parameters: &armservicelinker.CreateOrUpdateDryrunParameters{
	// 			ActionName: to.Ptr(armservicelinker.DryrunActionNameCreateOrUpdate),
	// 			AuthInfo: &armservicelinker.SecretAuthInfo{
	// 				AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
	// 				Name: to.Ptr("username"),
	// 			},
	// 			TargetService: &armservicelinker.AzureResource{
	// 				Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"),
	// 			},
	// 		},
	// 	},
	// }
}
