package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/SimGet.json
func ExampleSimsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSimsClient().Get(ctx, "testResourceGroupName", "testSimGroup", "testSimName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Sim = armmobilenetwork.Sim{
	// 	Name: to.Ptr("testSim"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/simGroups/sims"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/simGroups/testSimGroup/sims/testSim"),
	// 	SystemData: &armmobilenetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmobilenetwork.SimPropertiesFormat{
	// 		DeviceType: to.Ptr("Video camera"),
	// 		IntegratedCircuitCardIdentifier: to.Ptr("8900000000000000000"),
	// 		InternationalMobileSubscriberIdentity: to.Ptr("00000"),
	// 		ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
	// 		SimPolicy: &armmobilenetwork.SimPolicyResourceID{
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
	// 		},
	// 		SimState: to.Ptr(armmobilenetwork.SimStateEnabled),
	// 		SiteProvisioningState: map[string]*armmobilenetwork.SiteProvisioningState{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite": to.Ptr(armmobilenetwork.SiteProvisioningStateProvisioned),
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite2": to.Ptr(armmobilenetwork.SiteProvisioningStateProvisioned),
	// 		},
	// 		StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
	// 			{
	// 				AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
	// 				},
	// 				Slice: &armmobilenetwork.SliceResourceID{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
	// 				},
	// 				StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
	// 					IPv4Address: to.Ptr("2.4.0.1"),
	// 				},
	// 		}},
	// 	},
	// }
}
