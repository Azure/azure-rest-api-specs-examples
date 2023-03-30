package armengagementfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/engagementfabric/armengagementfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/engagementfabric/resource-manager/Microsoft.EngagementFabric/preview/2018-09-01/examples/AccountsListChannelTypesExample.json
func ExampleAccountsClient_ListChannelTypes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armengagementfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().ListChannelTypes(ctx, "ExampleRg", "ExampleAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ChannelTypeDescriptionList = armengagementfabric.ChannelTypeDescriptionList{
	// 	Value: []*armengagementfabric.ChannelTypeDescription{
	// 		{
	// 			ChannelDescription: to.Ptr("Description of mockChannel1"),
	// 			ChannelFunctions: []*string{
	// 				to.Ptr("MockFunction1"),
	// 				to.Ptr("MockFunction2")},
	// 				ChannelType: to.Ptr("MockChannel1"),
	// 			},
	// 			{
	// 				ChannelDescription: to.Ptr("Description of mockChannel2"),
	// 				ChannelFunctions: []*string{
	// 					to.Ptr("MockFunction1"),
	// 					to.Ptr("MockFunction3")},
	// 					ChannelType: to.Ptr("MockChannel2"),
	// 				},
	// 				{
	// 					ChannelDescription: to.Ptr("Description of mockChannel3"),
	// 					ChannelFunctions: []*string{
	// 						to.Ptr("MockFunction1"),
	// 						to.Ptr("MockFunction2"),
	// 						to.Ptr("MockFunction3")},
	// 						ChannelType: to.Ptr("MockChannel3"),
	// 				}},
	// 			}
}
