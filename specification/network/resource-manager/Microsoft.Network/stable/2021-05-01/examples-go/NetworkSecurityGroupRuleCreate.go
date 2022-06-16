package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkSecurityGroupRuleCreate.json
func ExampleSecurityRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewSecurityRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<network-security-group-name>",
		"<security-rule-name>",
		armnetwork.SecurityRule{
			Properties: &armnetwork.SecurityRulePropertiesFormat{
				Access:                   to.Ptr(armnetwork.SecurityRuleAccessDeny),
				DestinationAddressPrefix: to.Ptr("<destination-address-prefix>"),
				DestinationPortRange:     to.Ptr("<destination-port-range>"),
				Direction:                to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
				Priority:                 to.Ptr[int32](100),
				SourceAddressPrefix:      to.Ptr("<source-address-prefix>"),
				SourcePortRange:          to.Ptr("<source-port-range>"),
				Protocol:                 to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
			},
		},
		&armnetwork.SecurityRulesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
