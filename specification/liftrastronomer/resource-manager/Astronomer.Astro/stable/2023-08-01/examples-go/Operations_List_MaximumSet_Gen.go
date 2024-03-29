package armastro_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/astro/armastro"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/liftrastronomer/resource-manager/Astronomer.Astro/stable/2023-08-01/examples/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armastro.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armastro.OperationListResult{
		// 	Value: []*armastro.Operation{
		// 		{
		// 			Name: to.Ptr("zabhglnki"),
		// 			ActionType: to.Ptr(armastro.ActionTypeInternal),
		// 			Display: &armastro.OperationDisplay{
		// 				Description: to.Ptr("nkucjlsbtriwdgedbxlknbwfz"),
		// 				Operation: to.Ptr("teozafbxkiagahfypii"),
		// 				Provider: to.Ptr("hgepwsvbptqbigephgxoxyll"),
		// 				Resource: to.Ptr("thhzqbtxxi"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr(armastro.OriginUser),
		// 	}},
		// }
	}
}
