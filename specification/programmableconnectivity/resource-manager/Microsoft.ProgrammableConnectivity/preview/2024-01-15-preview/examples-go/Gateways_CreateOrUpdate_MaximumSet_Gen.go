package armprogrammableconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/programmableconnectivity/armprogrammableconnectivity"
)

// Generated from example definition: 2024-01-15-preview/Gateways_CreateOrUpdate_MaximumSet_Gen.json
func ExampleGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprogrammableconnectivity.NewClientFactory("B976474B-99FA-4C25-A3BD-8B05C3C3D07A", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGatewaysClient().BeginCreateOrUpdate(ctx, "rgopenapi", "pgzk", armprogrammableconnectivity.Gateway{
		Properties: &armprogrammableconnectivity.GatewayProperties{},
		Tags: map[string]*string{
			"key2642": to.Ptr("ykmlftvwwpvcmriffxqh"),
		},
		Location: to.Ptr("oryhozfmeohscezl"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armprogrammableconnectivity.GatewaysClientCreateOrUpdateResponse{
	// 	Gateway: &armprogrammableconnectivity.Gateway{
	// 		Properties: &armprogrammableconnectivity.GatewayProperties{
	// 			OperatorAPIConnections: []*string{
	// 				to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/example-rg/providers/Microsoft.ProgrammableConnectivity/operatorApiConnections/uecwablqeufseigocrwf"),
	// 			},
	// 			GatewayBaseURL: to.Ptr("gfohvglkp"),
	// 			ProvisioningState: to.Ptr(armprogrammableconnectivity.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key2642": to.Ptr("ykmlftvwwpvcmriffxqh"),
	// 		},
	// 		Location: to.Ptr("oryhozfmeohscezl"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/example-rg/providers/Microsoft.ProgrammableConnectivity/gateways/cdvcixxcdhjqw"),
	// 		Name: to.Ptr("qwlnmndshgfrtkp"),
	// 		Type: to.Ptr("qqkelyjugwyforzd"),
	// 		SystemData: &armprogrammableconnectivity.SystemData{
	// 			CreatedBy: to.Ptr("kuprrapuolhnvju"),
	// 			CreatedByType: to.Ptr(armprogrammableconnectivity.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T16:41:38.838Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("lsmrhxnvkpmrxncylgqpkr"),
	// 			LastModifiedByType: to.Ptr(armprogrammableconnectivity.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T16:41:38.838Z"); return t}()),
	// 		},
	// 	},
	// }
}
