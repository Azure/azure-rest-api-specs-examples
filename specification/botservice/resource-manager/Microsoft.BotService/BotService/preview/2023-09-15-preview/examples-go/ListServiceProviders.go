package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice/v2"
)

// Generated from example definition: 2023-09-15-preview/ListServiceProviders.json
func ExampleBotConnectionClient_ListServiceProviders() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbotservice.NewClientFactory("subscription-id", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBotConnectionClient().ListServiceProviders(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbotservice.BotConnectionClientListServiceProvidersResponse{
	// 	ServiceProviderResponseList: &armbotservice.ServiceProviderResponseList{
	// 		Value: []*armbotservice.ServiceProvider{
	// 			{
	// 				Properties: &armbotservice.ServiceProviderProperties{
	// 					DevPortalURL: to.Ptr("sampleDevPortalUrl"),
	// 					DisplayName: to.Ptr("sample service provider display name"),
	// 					IconURL: to.Ptr("sampleIconUrl"),
	// 					ID: to.Ptr("sampleServiceProviderId"),
	// 					Parameters: []*armbotservice.ServiceProviderParameter{
	// 						{
	// 							Name: to.Ptr("sampleParameterName"),
	// 							Type: to.Ptr("sampleParameterType"),
	// 							Description: to.Ptr("sampleDescription"),
	// 							Default: to.Ptr("sampleDefaultValue"),
	// 							DisplayName: to.Ptr("sampleDisplayName"),
	// 							HelpURL: to.Ptr("sampleHelpUrl"),
	// 						},
	// 					},
	// 					ServiceProviderName: to.Ptr("sampleServiceProvider"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
