package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/FavoriteProcessCreate.json
func ExampleFavoriteProcessesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFavoriteProcessesClient().Create(ctx, "contoso-rg1", "contoso-testBaseAccount1", "contoso-package2", "testAppProcess", armtestbase.FavoriteProcessResource{
		Properties: &armtestbase.FavoriteProcessProperties{
			ActualProcessName: to.Ptr("testApp&.exe"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FavoriteProcessResource = armtestbase.FavoriteProcessResource{
	// 	Name: to.Ptr("testAppProcess"),
	// 	Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/packages/favoriteProcesses"),
	// 	ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/packages/contoso-package2/favoriteProcesses/testAppProcess"),
	// 	Properties: &armtestbase.FavoriteProcessProperties{
	// 		ActualProcessName: to.Ptr("testApp&.exe"),
	// 	},
	// }
}
