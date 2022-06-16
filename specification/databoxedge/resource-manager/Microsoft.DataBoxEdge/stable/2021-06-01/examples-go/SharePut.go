package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/SharePut.json
func ExampleSharesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewSharesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<device-name>",
		"<name>",
		"<resource-group-name>",
		armdataboxedge.Share{
			Properties: &armdataboxedge.ShareProperties{
				Description:    to.StringPtr("<description>"),
				AccessProtocol: armdataboxedge.ShareAccessProtocol("SMB").ToPtr(),
				AzureContainerInfo: &armdataboxedge.AzureContainerInfo{
					ContainerName:              to.StringPtr("<container-name>"),
					DataFormat:                 armdataboxedge.AzureContainerDataFormat("BlockBlob").ToPtr(),
					StorageAccountCredentialID: to.StringPtr("<storage-account-credential-id>"),
				},
				DataPolicy:       armdataboxedge.DataPolicy("Cloud").ToPtr(),
				MonitoringStatus: armdataboxedge.MonitoringStatus("Enabled").ToPtr(),
				ShareStatus:      armdataboxedge.ShareStatus("Online").ToPtr(),
				UserAccessRights: []*armdataboxedge.UserAccessRight{
					{
						AccessType: armdataboxedge.ShareAccessType("Change").ToPtr(),
						UserID:     to.StringPtr("<user-id>"),
					}},
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
	log.Printf("Response result: %#v\n", res.SharesClientCreateOrUpdateResult)
}
