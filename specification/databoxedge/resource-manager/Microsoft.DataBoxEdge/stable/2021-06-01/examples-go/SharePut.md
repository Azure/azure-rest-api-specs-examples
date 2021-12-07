Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataboxedge%2Farmdataboxedge%2Fv0.1.0/sdk/resourcemanager/databoxedge/armdataboxedge/README.md) on how to add the SDK to your project and authenticate.

```go
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
				AccessProtocol: armdataboxedge.ShareAccessProtocolSMB.ToPtr(),
				AzureContainerInfo: &armdataboxedge.AzureContainerInfo{
					ContainerName:              to.StringPtr("<container-name>"),
					DataFormat:                 armdataboxedge.AzureContainerDataFormatBlockBlob.ToPtr(),
					StorageAccountCredentialID: to.StringPtr("<storage-account-credential-id>"),
				},
				DataPolicy:       armdataboxedge.DataPolicyCloud.ToPtr(),
				MonitoringStatus: armdataboxedge.MonitoringStatusEnabled.ToPtr(),
				ShareStatus:      armdataboxedge.ShareStatusOffline.ToPtr(),
				UserAccessRights: []*armdataboxedge.UserAccessRight{
					{
						AccessType: armdataboxedge.ShareAccessTypeChange.ToPtr(),
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
	log.Printf("Share.ID: %s\n", *res.ID)
}
```
