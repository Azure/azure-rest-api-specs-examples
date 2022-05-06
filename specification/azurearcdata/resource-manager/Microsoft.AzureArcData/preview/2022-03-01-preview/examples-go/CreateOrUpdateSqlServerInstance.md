Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazurearcdata%2Farmazurearcdata%2Fv0.4.0/sdk/resourcemanager/azurearcdata/armazurearcdata/README.md) on how to add the SDK to your project and authenticate.

```go
package armazurearcdata_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/CreateOrUpdateSqlServerInstance.json
func ExampleSQLServerInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armazurearcdata.NewSQLServerInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<sql-server-instance-name>",
		armazurearcdata.SQLServerInstance{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"mytag": to.Ptr("myval"),
			},
			Properties: &armazurearcdata.SQLServerInstanceProperties{
				AzureDefenderStatus:            to.Ptr(armazurearcdata.DefenderStatusProtected),
				AzureDefenderStatusLastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t }()),
				Collation:                      to.Ptr("<collation>"),
				ContainerResourceID:            to.Ptr("<container-resource-id>"),
				CurrentVersion:                 to.Ptr("<current-version>"),
				Edition:                        to.Ptr(armazurearcdata.EditionTypeDeveloper),
				HostType:                       to.Ptr(armazurearcdata.HostTypePhysicalServer),
				InstanceName:                   to.Ptr("<instance-name>"),
				LicenseType:                    to.Ptr(armazurearcdata.ArcSQLServerLicenseTypeFree),
				PatchLevel:                     to.Ptr("<patch-level>"),
				ProductID:                      to.Ptr("<product-id>"),
				Status:                         to.Ptr(armazurearcdata.ConnectionStatusRegistered),
				TCPDynamicPorts:                to.Ptr("<tcpdynamic-ports>"),
				TCPStaticPorts:                 to.Ptr("<tcpstatic-ports>"),
				VCore:                          to.Ptr("<vcore>"),
				Version:                        to.Ptr(armazurearcdata.SQLVersionSQLServer2012),
			},
		},
		&armazurearcdata.SQLServerInstancesClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
