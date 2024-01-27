package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/685aad3f33d355c1d9c89d493ee9398865367bd8/specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Services_ListSupportedApmTypes.json
func ExampleServicesClient_NewListSupportedApmTypesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicesClient().NewListSupportedApmTypesPager("myResourceGroup", "myservice", nil)
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
		// page.SupportedApmTypes = armappplatform.SupportedApmTypes{
		// 	Value: []*armappplatform.SupportedApmType{
		// 		{
		// 			Name: to.Ptr("AppDynamics"),
		// 		},
		// 		{
		// 			Name: to.Ptr("ApplicationInsights"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Dynatrace"),
		// 		},
		// 		{
		// 			Name: to.Ptr("ElasticAPM"),
		// 		},
		// 		{
		// 			Name: to.Ptr("NewRelic"),
		// 	}},
		// }
	}
}
