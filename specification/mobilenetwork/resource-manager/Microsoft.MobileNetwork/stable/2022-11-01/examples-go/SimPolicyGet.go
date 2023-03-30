package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/340d577969b7bff5ad0488d79543314bc17daa50/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/SimPolicyGet.json
func ExampleSimPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSimPoliciesClient().Get(ctx, "rg1", "testMobileNetwork", "testPolicy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SimPolicy = armmobilenetwork.SimPolicy{
	// 	Name: to.Ptr("testPolicy"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/simPolicy"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/testPolicy"),
	// 	SystemData: &armmobilenetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armmobilenetwork.SimPolicyPropertiesFormat{
	// 		DefaultSlice: &armmobilenetwork.SliceResourceID{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
	// 		},
	// 		ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
	// 		RegistrationTimer: to.Ptr[int32](3240),
	// 		SiteProvisioningState: map[string]*armmobilenetwork.SiteProvisioningState{
	// 			"/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite": to.Ptr(armmobilenetwork.SiteProvisioningStateAdding),
	// 			"/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite2": to.Ptr(armmobilenetwork.SiteProvisioningStateAdding),
	// 		},
	// 		SliceConfigurations: []*armmobilenetwork.SliceConfiguration{
	// 			{
	// 				DataNetworkConfigurations: []*armmobilenetwork.DataNetworkConfiguration{
	// 					{
	// 						FiveQi: to.Ptr[int32](9),
	// 						AdditionalAllowedSessionTypes: []*armmobilenetwork.PduSessionType{
	// 						},
	// 						AllocationAndRetentionPriorityLevel: to.Ptr[int32](9),
	// 						AllowedServices: []*armmobilenetwork.ServiceResourceID{
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/services/testService"),
	// 						}},
	// 						DataNetwork: &armmobilenetwork.DataNetworkResourceID{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/dataNetworks/testdataNetwork"),
	// 						},
	// 						DefaultSessionType: to.Ptr(armmobilenetwork.PduSessionTypeIPv4),
	// 						MaximumNumberOfBufferedPackets: to.Ptr[int32](200),
	// 						PreemptionCapability: to.Ptr(armmobilenetwork.PreemptionCapabilityNotPreempt),
	// 						PreemptionVulnerability: to.Ptr(armmobilenetwork.PreemptionVulnerabilityPreemptable),
	// 						SessionAmbr: &armmobilenetwork.Ambr{
	// 							Downlink: to.Ptr("1 Gbps"),
	// 							Uplink: to.Ptr("500 Mbps"),
	// 						},
	// 				}},
	// 				DefaultDataNetwork: &armmobilenetwork.DataNetworkResourceID{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/dataNetworks/testdataNetwork"),
	// 				},
	// 				Slice: &armmobilenetwork.SliceResourceID{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
	// 				},
	// 		}},
	// 		UeAmbr: &armmobilenetwork.Ambr{
	// 			Downlink: to.Ptr("1 Gbps"),
	// 			Uplink: to.Ptr("500 Mbps"),
	// 		},
	// 	},
	// }
}
