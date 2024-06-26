package armchangeanalysis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/changeanalysis/armchangeanalysis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/changeanalysis/resource-manager/Microsoft.ChangeAnalysis/stable/2021-04-01/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchangeanalysis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(&armchangeanalysis.OperationsClientListOptions{SkipToken: nil})
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
		// page.ResourceProviderOperationList = armchangeanalysis.ResourceProviderOperationList{
		// 	Value: []*armchangeanalysis.ResourceProviderOperationDefinition{
		// 		{
		// 			Name: to.Ptr("Microsoft.ChangeAnalysis/register/action"),
		// 			Display: &armchangeanalysis.ResourceProviderOperationDisplay{
		// 				Operation: to.Ptr("Register Microsoft Change Analysis resource provider with a subscriptions"),
		// 				Provider: to.Ptr("Microsoft Change Analysis"),
		// 				Resource: to.Ptr(""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ChangeAnalysis/changes/read"),
		// 			Display: &armchangeanalysis.ResourceProviderOperationDisplay{
		// 				Operation: to.Ptr("Read Azure Application Change Analysis Servie Change"),
		// 				Provider: to.Ptr("Microsoft Change Analysis"),
		// 				Resource: to.Ptr("Azure Application Change Analysis Service Change"),
		// 			},
		// 	}},
		// }
	}
}
