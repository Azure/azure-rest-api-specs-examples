package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: 2025-07-15-preview/Domains_Update.json
func ExampleDomainsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("8f6b6269-84f2-4d09-9e31-1127efcd1e40", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDomainsClient().BeginUpdate(ctx, "examplerg", "exampledomain1", armeventgrid.DomainUpdateParameters{
		Properties: &armeventgrid.DomainUpdateParameterProperties{
			InboundIPRules: []*armeventgrid.InboundIPRule{
				{
					Action: to.Ptr(armeventgrid.IPActionTypeAllow),
					IPMask: to.Ptr("12.18.30.15"),
				},
				{
					Action: to.Ptr(armeventgrid.IPActionTypeAllow),
					IPMask: to.Ptr("12.18.176.1"),
				},
			},
			PublicNetworkAccess: to.Ptr(armeventgrid.PublicNetworkAccessEnabled),
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
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
	// res = armeventgrid.DomainsClientUpdateResponse{
	// }
}
