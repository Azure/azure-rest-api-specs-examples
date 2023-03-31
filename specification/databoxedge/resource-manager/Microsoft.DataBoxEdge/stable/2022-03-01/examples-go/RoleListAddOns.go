package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/RoleListAddOns.json
func ExampleAddonsClient_NewListByRolePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAddonsClient().NewListByRolePager("testedgedevice", "IoTRole1", "GroupForEdgeAutomation", nil)
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
		// page.AddonList = armdataboxedge.AddonList{
		// 	Value: []armdataboxedge.AddonClassification{
		// 		&armdataboxedge.ArcAddon{
		// 			Name: to.Ptr("arcName"),
		// 			Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/roles/addons"),
		// 			ID: to.Ptr("/subscriptions/0d44739e-0563-474f-97e7-24a0cdb23b29/resourcegroups/prpare/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/addonExamples/roles/kubernetesRole/addons/arcName"),
		// 			Kind: to.Ptr(armdataboxedge.AddonTypeArcForKubernetes),
		// 			Properties: &armdataboxedge.ArcAddonProperties{
		// 				HostPlatform: to.Ptr(armdataboxedge.PlatformTypeLinux),
		// 				HostPlatformType: to.Ptr(armdataboxedge.HostPlatformTypeKubernetesCluster),
		// 				ProvisioningState: to.Ptr(armdataboxedge.AddonState("Succeeded")),
		// 				ResourceGroupName: to.Ptr("testrg1"),
		// 				ResourceLocation: to.Ptr("EastUS"),
		// 				ResourceName: to.Ptr("testresource1"),
		// 				SubscriptionID: to.Ptr("0d44739e-0563-474f-97e7-24a0cdb23b29"),
		// 				Version: to.Ptr("0.2.18"),
		// 			},
		// 	}},
		// }
	}
}
