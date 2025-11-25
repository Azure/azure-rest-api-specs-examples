package armnewrelicobservability_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/newrelic/armnewrelicobservability/v2"
)

// Generated from example definition: 2025-05-01-preview/Monitors_ListHosts_MinimumSet_Gen.json
func ExampleMonitorsClient_NewListHostsPager_monitorsListHostsMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnewrelicobservability.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListHostsPager("rgopenapi", "ipxmlcbonyxtolzejcjshkmlron", armnewrelicobservability.HostsGetRequest{
		UserEmail: to.Ptr("ruxvg@xqkmdhrnoo.hlmbpm"),
		VMIDs: []*string{
			to.Ptr("xzphvxvfmvjrnsgyns"),
		},
	}, nil)
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
		// page = armnewrelicobservability.MonitorsClientListHostsResponse{
		// 	VMHostsListResponse: armnewrelicobservability.VMHostsListResponse{
		// 		Value: []*armnewrelicobservability.VMInfo{
		// 		},
		// 	},
		// }
	}
}
