Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazurearcdata%2Farmazurearcdata%2Fv0.2.0/sdk/resourcemanager/azurearcdata/armazurearcdata/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/CreateOrUpdateDataController.json
func ExampleDataControllersClient_BeginPutDataController() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewDataControllersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginPutDataController(ctx,
		"<resource-group-name>",
		"<data-controller-name>",
		armazurearcdata.DataControllerResource{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"mytag": to.StringPtr("myval"),
			},
			ExtendedLocation: &armazurearcdata.ExtendedLocation{
				Name: to.StringPtr("<name>"),
				Type: armazurearcdata.ExtendedLocationTypes("CustomLocation").ToPtr(),
			},
			Properties: &armazurearcdata.DataControllerProperties{
				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
					Password: to.StringPtr("<password>"),
					Username: to.StringPtr("<username>"),
				},
				ClusterID:      to.StringPtr("<cluster-id>"),
				ExtensionID:    to.StringPtr("<extension-id>"),
				Infrastructure: armazurearcdata.InfrastructureOnpremises.ToPtr(),
				LogAnalyticsWorkspaceConfig: &armazurearcdata.LogAnalyticsWorkspaceConfig{
					PrimaryKey:  to.StringPtr("<primary-key>"),
					WorkspaceID: to.StringPtr("<workspace-id>"),
				},
				LogsDashboardCredential: &armazurearcdata.BasicLoginInformation{
					Password: to.StringPtr("<password>"),
					Username: to.StringPtr("<username>"),
				},
				MetricsDashboardCredential: &armazurearcdata.BasicLoginInformation{
					Password: to.StringPtr("<password>"),
					Username: to.StringPtr("<username>"),
				},
				OnPremiseProperty: &armazurearcdata.OnPremiseProperty{
					ID:               to.StringPtr("<id>"),
					PublicSigningKey: to.StringPtr("<public-signing-key>"),
				},
				UploadServicePrincipal: &armazurearcdata.UploadServicePrincipal{
					Authority:    to.StringPtr("<authority>"),
					ClientID:     to.StringPtr("<client-id>"),
					ClientSecret: to.StringPtr("<client-secret>"),
					TenantID:     to.StringPtr("<tenant-id>"),
				},
				UploadWatermark: &armazurearcdata.UploadWatermark{
					Logs:    to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t }()),
					Metrics: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t }()),
					Usages:  to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t }()),
				},
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
	log.Printf("Response result: %#v\n", res.DataControllersClientPutDataControllerResult)
}
```
