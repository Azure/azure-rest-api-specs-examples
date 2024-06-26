package armdelegatednetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/delegatednetwork/armdelegatednetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/putController.json
func ExampleControllerClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdelegatednetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewControllerClient().BeginCreate(ctx, "TestRG", "testcontroller", armdelegatednetwork.DelegatedController{
		Location: to.Ptr("West US"),
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
	// res.DelegatedController = armdelegatednetwork.DelegatedController{
	// 	Name: to.Ptr("testcontroller"),
	// 	Type: to.Ptr("Microsoft.DelegatedNetwork/controller"),
	// 	ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.DelegatedNetwork/controller/testcontroller"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armdelegatednetwork.DelegatedControllerProperties{
	// 		DncAppID: to.Ptr("ac1192d7-503f-477a-9cfe-4efc3ee2bd60"),
	// 		DncEndpoint: to.Ptr("https://orch.useast.dnc.azure.com"),
	// 		DncTenantID: to.Ptr("66192d7-503f-477a-9cfe-4efc3ee2bd60"),
	// 		ProvisioningState: to.Ptr(armdelegatednetwork.ControllerStateSucceeded),
	// 		ResourceGUID: to.Ptr("5a82cbcf-e8ea-4175-ac2b-ad36a73f9801"),
	// 	},
	// }
}
