package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/ServiceGet.json
func ExampleServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().Get(ctx, "rg1", "testMobileNetwork", "TestService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Service = armmobilenetwork.Service{
	// 	Name: to.Ptr("testPolicy"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/service"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/services/TestService"),
	// 	SystemData: &armmobilenetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armmobilenetwork.ServicePropertiesFormat{
	// 		PccRules: []*armmobilenetwork.PccRuleConfiguration{
	// 			{
	// 				RuleName: to.Ptr("default-rule"),
	// 				RulePrecedence: to.Ptr[int32](255),
	// 				RuleQosPolicy: &armmobilenetwork.PccRuleQosPolicy{
	// 					FiveQi: to.Ptr[int32](9),
	// 					AllocationAndRetentionPriorityLevel: to.Ptr[int32](9),
	// 					MaximumBitRate: &armmobilenetwork.Ambr{
	// 						Downlink: to.Ptr("1 Gbps"),
	// 						Uplink: to.Ptr("500 Mbps"),
	// 					},
	// 					PreemptionCapability: to.Ptr(armmobilenetwork.PreemptionCapabilityNotPreempt),
	// 					PreemptionVulnerability: to.Ptr(armmobilenetwork.PreemptionVulnerabilityPreemptable),
	// 				},
	// 				ServiceDataFlowTemplates: []*armmobilenetwork.ServiceDataFlowTemplate{
	// 					{
	// 						Direction: to.Ptr(armmobilenetwork.SdfDirectionUplink),
	// 						Ports: []*string{
	// 						},
	// 						RemoteIPList: []*string{
	// 							to.Ptr("10.3.4.0/24")},
	// 							TemplateName: to.Ptr("IP-to-server"),
	// 							Protocol: []*string{
	// 								to.Ptr("ip")},
	// 						}},
	// 						TrafficControl: to.Ptr(armmobilenetwork.TrafficControlPermissionEnabled),
	// 				}},
	// 				ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
	// 				ServicePrecedence: to.Ptr[int32](255),
	// 				ServiceQosPolicy: &armmobilenetwork.QosPolicy{
	// 					FiveQi: to.Ptr[int32](9),
	// 					AllocationAndRetentionPriorityLevel: to.Ptr[int32](9),
	// 					MaximumBitRate: &armmobilenetwork.Ambr{
	// 						Downlink: to.Ptr("1 Gbps"),
	// 						Uplink: to.Ptr("500 Mbps"),
	// 					},
	// 					PreemptionCapability: to.Ptr(armmobilenetwork.PreemptionCapabilityNotPreempt),
	// 					PreemptionVulnerability: to.Ptr(armmobilenetwork.PreemptionVulnerabilityPreemptable),
	// 				},
	// 			},
	// 		}
}
