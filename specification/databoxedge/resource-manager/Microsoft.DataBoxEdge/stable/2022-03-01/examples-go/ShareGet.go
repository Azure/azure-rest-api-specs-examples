package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/ShareGet.json
func ExampleSharesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSharesClient().Get(ctx, "testedgedevice", "smbshare", "GroupForEdgeAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Share = armdataboxedge.Share{
	// 	Name: to.Ptr("smbshare"),
	// 	Type: to.Ptr("dataBoxEdgeDevices/shares"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/shares/smbshare"),
	// 	Properties: &armdataboxedge.ShareProperties{
	// 		Description: to.Ptr(""),
	// 		AccessProtocol: to.Ptr(armdataboxedge.ShareAccessProtocolSMB),
	// 		AzureContainerInfo: &armdataboxedge.AzureContainerInfo{
	// 			ContainerName: to.Ptr("testContainerSMB"),
	// 			DataFormat: to.Ptr(armdataboxedge.AzureContainerDataFormatBlockBlob),
	// 			StorageAccountCredentialID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccountCredentials/sac1"),
	// 		},
	// 		ClientAccessRights: []*armdataboxedge.ClientAccessRight{
	// 		},
	// 		DataPolicy: to.Ptr(armdataboxedge.DataPolicyCloud),
	// 		MonitoringStatus: to.Ptr(armdataboxedge.MonitoringStatusDisabled),
	// 		RefreshDetails: &armdataboxedge.RefreshDetails{
	// 		},
	// 		ShareMappings: []*armdataboxedge.MountPointMap{
	// 		},
	// 		ShareStatus: to.Ptr(armdataboxedge.ShareStatus("Online")),
	// 		UserAccessRights: []*armdataboxedge.UserAccessRight{
	// 			{
	// 				AccessType: to.Ptr(armdataboxedge.ShareAccessTypeChange),
	// 				UserID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/users/user2"),
	// 		}},
	// 	},
	// }
}
