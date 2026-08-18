package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/GenerateExpressRouteLagsLOA.json
func ExampleExpressRouteLagsClient_GenerateLoa() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRouteLagsClient().GenerateLoa(ctx, "rg1", "lagName", armnetwork.GenerateExpressRouteLagsLOARequest{
		CustomerName: to.Ptr("Customer Name"),
		Members: []*string{
			to.Ptr("member1"),
			to.Ptr("member2"),
			to.Ptr("member3"),
			to.Ptr("member4"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ExpressRouteLagsClientGenerateLoaResponse{
	// 	GenerateExpressRouteLagsLOAResult: armnetwork.GenerateExpressRouteLagsLOAResult{
	// 		EncodedContent: to.Ptr("encoded letter of authorization"),
	// 	},
	// }
}
