package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/SharePut.json
func ExampleSharesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataboxedge.NewSharesClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testedgedevice",
		"smbshare",
		"GroupForEdgeAutomation",
		armdataboxedge.Share{
			Properties: &armdataboxedge.ShareProperties{
				Description:    to.Ptr(""),
				AccessProtocol: to.Ptr(armdataboxedge.ShareAccessProtocolSMB),
				AzureContainerInfo: &armdataboxedge.AzureContainerInfo{
					ContainerName:              to.Ptr("testContainerSMB"),
					DataFormat:                 to.Ptr(armdataboxedge.AzureContainerDataFormatBlockBlob),
					StorageAccountCredentialID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccountCredentials/sac1"),
				},
				DataPolicy:       to.Ptr(armdataboxedge.DataPolicyCloud),
				MonitoringStatus: to.Ptr(armdataboxedge.MonitoringStatusEnabled),
				ShareStatus:      to.Ptr(armdataboxedge.ShareStatus("Online")),
				UserAccessRights: []*armdataboxedge.UserAccessRight{
					{
						AccessType: to.Ptr(armdataboxedge.ShareAccessTypeChange),
						UserID:     to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/users/user2"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
