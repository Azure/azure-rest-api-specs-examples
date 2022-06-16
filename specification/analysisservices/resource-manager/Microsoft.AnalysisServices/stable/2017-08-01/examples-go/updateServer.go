package armanalysisservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/updateServer.json
func ExampleServersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armanalysisservices.NewServersClient("613192d7-503f-477a-9cfe-4efc3ee2bd60", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"TestRG",
		"azsdktest",
		armanalysisservices.ServerUpdateParameters{
			Properties: &armanalysisservices.ServerMutableProperties{
				AsAdministrators: &armanalysisservices.ServerAdministrators{
					Members: []*string{
						to.Ptr("azsdktest@microsoft.com"),
						to.Ptr("azsdktest2@microsoft.com")},
				},
			},
			SKU: &armanalysisservices.ResourceSKU{
				Name:     to.Ptr("S1"),
				Capacity: to.Ptr[int32](1),
				Tier:     to.Ptr(armanalysisservices.SKUTierStandard),
			},
			Tags: map[string]*string{
				"testKey": to.Ptr("testValue"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
