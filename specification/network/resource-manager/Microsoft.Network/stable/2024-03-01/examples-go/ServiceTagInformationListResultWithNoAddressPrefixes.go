package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/ServiceTagInformationListResultWithNoAddressPrefixes.json
func ExampleServiceTagInformationClient_NewListPager_getListOfServiceTagsWithNoAddressPrefixes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServiceTagInformationClient().NewListPager("westeurope", &armnetwork.ServiceTagInformationClientListOptions{NoAddressPrefixes: to.Ptr(true),
		TagName: nil,
	})
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
		// page.ServiceTagInformationListResult = armnetwork.ServiceTagInformationListResult{
		// 	Value: []*armnetwork.ServiceTagInformation{
		// 		{
		// 			Name: to.Ptr("ApiManagement"),
		// 			ID: to.Ptr("ApiManagement"),
		// 			Properties: &armnetwork.ServiceTagInformationPropertiesFormat{
		// 				AddressPrefixes: []*string{
		// 				},
		// 				ChangeNumber: to.Ptr("7"),
		// 				Region: to.Ptr(""),
		// 				SystemService: to.Ptr("AzureApiManagement"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ApiManagement.AustraliaCentral"),
		// 			ID: to.Ptr("ApiManagement.AustraliaCentral"),
		// 			Properties: &armnetwork.ServiceTagInformationPropertiesFormat{
		// 				AddressPrefixes: []*string{
		// 				},
		// 				ChangeNumber: to.Ptr("2"),
		// 				Region: to.Ptr("australiacentral"),
		// 				SystemService: to.Ptr("AzureApiManagement"),
		// 			},
		// 	}},
		// }
	}
}
