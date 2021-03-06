package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/InboundNatRuleCreate.json
func ExampleInboundNatRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewInboundNatRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<load-balancer-name>",
		"<inbound-nat-rule-name>",
		armnetwork.InboundNatRule{
			Properties: &armnetwork.InboundNatRulePropertiesFormat{
				BackendPort:      to.Ptr[int32](3389),
				EnableFloatingIP: to.Ptr(false),
				EnableTCPReset:   to.Ptr(false),
				FrontendIPConfiguration: &armnetwork.SubResource{
					ID: to.Ptr("<id>"),
				},
				FrontendPort:         to.Ptr[int32](3390),
				IdleTimeoutInMinutes: to.Ptr[int32](4),
				Protocol:             to.Ptr(armnetwork.TransportProtocolTCP),
			},
		},
		&armnetwork.InboundNatRulesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
