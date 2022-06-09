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
	}
	ctx := context.Background()
	client, err := armazurearcdata.NewSQLServerInstancesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"testrg",
		"testsqlServerInstance",
		armazurearcdata.SQLServerInstance{
			Location: to.Ptr("northeurope"),
			Tags: map[string]*string{
				"mytag": to.Ptr("myval"),
			},
			Properties: &armazurearcdata.SQLServerInstanceProperties{
				AzureDefenderStatus:            to.Ptr(armazurearcdata.DefenderStatusProtected),
				AzureDefenderStatusLastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t }()),
				Collation:                      to.Ptr("collation"),
				ContainerResourceID:            to.Ptr("Resource id of hosting Arc Machine"),
				CurrentVersion:                 to.Ptr("2012"),
				Edition:                        to.Ptr(armazurearcdata.EditionTypeDeveloper),
				HostType:                       to.Ptr(armazurearcdata.HostTypePhysicalServer),
				InstanceName:                   to.Ptr("name of instance"),
				LicenseType:                    to.Ptr(armazurearcdata.ArcSQLServerLicenseTypeFree),
				PatchLevel:                     to.Ptr("patchlevel"),
				ProductID:                      to.Ptr("sql id"),
				Status:                         to.Ptr(armazurearcdata.ConnectionStatusRegistered),
				TCPDynamicPorts:                to.Ptr("1433"),
				TCPStaticPorts:                 to.Ptr("1433"),
				VCore:                          to.Ptr("4"),
				Version:                        to.Ptr(armazurearcdata.SQLVersionSQLServer2012),
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazurearcdata%2Farmazurearcdata%2Fv0.5.0/sdk/resourcemanager/azurearcdata/armazurearcdata/README.md) on how to add the SDK to your project and authenticate.
