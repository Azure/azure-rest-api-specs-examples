package armmobilenetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"
)

// x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/ServiceCreate.json
func ExampleServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmobilenetwork.NewServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<mobile-network-name>",
		"<service-name>",
		armmobilenetwork.Service{
			Location: to.StringPtr("<location>"),
			Properties: &armmobilenetwork.ServicePropertiesFormat{
				PccRules: []*armmobilenetwork.PccRuleConfiguration{
					{
						RuleName:       to.StringPtr("<rule-name>"),
						RulePrecedence: to.Int32Ptr(255),
						RuleQosPolicy: &armmobilenetwork.PccRuleQosPolicy{
							FiveQi:                              to.Int32Ptr(9),
							AllocationAndRetentionPriorityLevel: to.Int32Ptr(9),
							MaximumBitRate: &armmobilenetwork.Ambr{
								Downlink: to.StringPtr("<downlink>"),
								Uplink:   to.StringPtr("<uplink>"),
							},
							PreemptionCapability:    armmobilenetwork.PreemptionCapability("NotPreempt").ToPtr(),
							PreemptionVulnerability: armmobilenetwork.PreemptionVulnerability("Preemptable").ToPtr(),
						},
						ServiceDataFlowTemplates: []*armmobilenetwork.ServiceDataFlowTemplate{
							{
								Direction: armmobilenetwork.SdfDirection("Uplink").ToPtr(),
								Ports:     []*string{},
								RemoteIPList: []*string{
									to.StringPtr("10.3.4.0/24")},
								TemplateName: to.StringPtr("<template-name>"),
								Protocol: []*string{
									to.StringPtr("ip")},
							}},
						TrafficControl: armmobilenetwork.TrafficControlPermission("Enabled").ToPtr(),
					}},
				ServicePrecedence: to.Int32Ptr(255),
				ServiceQosPolicy: &armmobilenetwork.QosPolicy{
					FiveQi:                              to.Int32Ptr(9),
					AllocationAndRetentionPriorityLevel: to.Int32Ptr(9),
					MaximumBitRate: &armmobilenetwork.Ambr{
						Downlink: to.StringPtr("<downlink>"),
						Uplink:   to.StringPtr("<uplink>"),
					},
					PreemptionCapability:    armmobilenetwork.PreemptionCapability("NotPreempt").ToPtr(),
					PreemptionVulnerability: armmobilenetwork.PreemptionVulnerability("Preemptable").ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServicesClientCreateOrUpdateResult)
}
