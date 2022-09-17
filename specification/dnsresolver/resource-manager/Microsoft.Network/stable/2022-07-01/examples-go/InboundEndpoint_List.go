package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/InboundEndpoint_List.json
func ExampleInboundEndpointsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdnsresolver.NewInboundEndpointsClient("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("sampleResourceGroup", "sampleDnsResolver", &armdnsresolver.InboundEndpointsClientListOptions{Top: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
