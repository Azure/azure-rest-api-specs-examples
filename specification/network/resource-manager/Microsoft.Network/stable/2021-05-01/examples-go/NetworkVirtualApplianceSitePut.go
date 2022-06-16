package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkVirtualApplianceSitePut.json
func ExampleVirtualApplianceSitesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewVirtualApplianceSitesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<network-virtual-appliance-name>",
		"<site-name>",
		armnetwork.VirtualApplianceSite{
			Properties: &armnetwork.VirtualApplianceSiteProperties{
				AddressPrefix: to.Ptr("<address-prefix>"),
				O365Policy: &armnetwork.Office365PolicyProperties{
					BreakOutCategories: &armnetwork.BreakOutCategoryPolicies{
						Default:  to.Ptr(true),
						Allow:    to.Ptr(true),
						Optimize: to.Ptr(true),
					},
				},
			},
		},
		&armnetwork.VirtualApplianceSitesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
