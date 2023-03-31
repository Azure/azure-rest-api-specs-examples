package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/DataBoxEdgeDevicePutWithDataResidency.json
func ExampleDevicesClient_CreateOrUpdate_dataBoxEdgeDevicePutWithDataResidency() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().CreateOrUpdate(ctx, "testedgedevice", "GroupForEdgeAutomation", armdataboxedge.Device{
		Location: to.Ptr("WUS"),
		Properties: &armdataboxedge.DeviceProperties{
			DataResidency: &armdataboxedge.DataResidency{
				Type: to.Ptr(armdataboxedge.DataResidencyTypeZoneReplication),
			},
		},
		SKU: &armdataboxedge.SKUInfo{
			Name: to.Ptr(armdataboxedge.SKUNameEdge),
			Tier: to.Ptr(armdataboxedge.SKUTierStandard),
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Device = armdataboxedge.Device{
	// 	Name: to.Ptr("testedgedevice"),
	// 	Location: to.Ptr("WUS"),
	// 	Properties: &armdataboxedge.DeviceProperties{
	// 		DataResidency: &armdataboxedge.DataResidency{
	// 			Type: to.Ptr(armdataboxedge.DataResidencyTypeZoneReplication),
	// 		},
	// 	},
	// 	SKU: &armdataboxedge.SKUInfo{
	// 		Name: to.Ptr(armdataboxedge.SKUNameEdge),
	// 		Tier: to.Ptr(armdataboxedge.SKUTierStandard),
	// 	},
	// 	Tags: map[string]*string{
	// 	},
	// }
}
