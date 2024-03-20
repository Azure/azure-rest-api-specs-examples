package armservicelinker_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c16ce913afbdaa073e8ca5e480f3b465db2de542/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2023-04-01-preview/examples/GetDaprConfigurations.json
func ExampleLinkersClient_NewListDaprConfigurationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLinkersClient().NewListDaprConfigurationsPager("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", nil)
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
		// page.DaprConfigurationList = armservicelinker.DaprConfigurationList{
		// 	Value: []*armservicelinker.DaprConfigurationResource{
		// 		{
		// 			Properties: &armservicelinker.DaprConfigurationProperties{
		// 				AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
		// 				DaprProperties: &armservicelinker.DaprProperties{
		// 					BindingComponentDirection: to.Ptr(armservicelinker.DaprBindingComponentDirectionInput),
		// 					ComponentType: to.Ptr("bindings"),
		// 					Metadata: []*armservicelinker.DaprMetadata{
		// 						{
		// 							Name: to.Ptr("containerName"),
		// 							Description: to.Ptr("The name of the container to be used for Dapr state."),
		// 							Required: to.Ptr(armservicelinker.DaprMetadataRequiredTrue),
		// 					}},
		// 					RuntimeVersion: to.Ptr("1.10"),
		// 					Version: to.Ptr("v1"),
		// 				},
		// 				TargetType: to.Ptr("MICROSOFT.STORAGE/STORAGEACCOUNTS/BLOBSERVICES"),
		// 			},
		// 	}},
		// }
	}
}
