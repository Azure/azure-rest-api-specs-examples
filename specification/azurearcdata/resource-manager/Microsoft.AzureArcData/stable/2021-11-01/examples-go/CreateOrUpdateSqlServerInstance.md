Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazurearcdata%2Farmazurearcdata%2Fv0.1.0/sdk/resourcemanager/azurearcdata/armazurearcdata/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/CreateOrUpdateSqlServerInstance.json
func ExampleSQLServerInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLServerInstancesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<sql-server-instance-name>",
		armazurearcdata.SQLServerInstance{
			TrackedResource: armazurearcdata.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"mytag": to.StringPtr("myval"),
				},
			},
			Properties: &armazurearcdata.SQLServerInstanceProperties{
				AzureDefenderStatus:            armazurearcdata.DefenderStatusProtected.ToPtr(),
				AzureDefenderStatusLastUpdated: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t }()),
				Collation:                      to.StringPtr("<collation>"),
				ContainerResourceID:            to.StringPtr("<container-resource-id>"),
				CurrentVersion:                 to.StringPtr("<current-version>"),
				Edition:                        armazurearcdata.EditionTypeDeveloper.ToPtr(),
				InstanceName:                   to.StringPtr("<instance-name>"),
				LicenseType:                    armazurearcdata.ArcSQLServerLicenseTypeFree.ToPtr(),
				PatchLevel:                     to.StringPtr("<patch-level>"),
				ProductID:                      to.StringPtr("<product-id>"),
				Status:                         armazurearcdata.ConnectionStatusConnected.ToPtr(),
				TCPDynamicPorts:                to.StringPtr("<tcpdynamic-ports>"),
				TCPStaticPorts:                 to.StringPtr("<tcpstatic-ports>"),
				VCore:                          to.StringPtr("<vcore>"),
				Version:                        armazurearcdata.SQLVersionSQLServer2017.ToPtr(),
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
	log.Printf("SQLServerInstance.ID: %s\n", *res.ID)
}
```
