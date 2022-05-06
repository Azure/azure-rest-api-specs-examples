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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/CreateOrUpdateDataController.json
func ExampleDataControllersClient_BeginPutDataController() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armazurearcdata.NewDataControllersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginPutDataController(ctx,
		"<resource-group-name>",
		"<data-controller-name>",
		armazurearcdata.DataControllerResource{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"mytag": to.Ptr("myval"),
			},
			ExtendedLocation: &armazurearcdata.ExtendedLocation{
				Name: to.Ptr("<name>"),
				Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
			},
			Properties: &armazurearcdata.DataControllerProperties{
				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
					Password: to.Ptr("<password>"),
					Username: to.Ptr("<username>"),
				},
				ClusterID:      to.Ptr("<cluster-id>"),
				ExtensionID:    to.Ptr("<extension-id>"),
				Infrastructure: to.Ptr(armazurearcdata.InfrastructureOnpremises),
				LogAnalyticsWorkspaceConfig: &armazurearcdata.LogAnalyticsWorkspaceConfig{
					PrimaryKey:  to.Ptr("<primary-key>"),
					WorkspaceID: to.Ptr("<workspace-id>"),
				},
				LogsDashboardCredential: &armazurearcdata.BasicLoginInformation{
					Password: to.Ptr("<password>"),
					Username: to.Ptr("<username>"),
				},
				MetricsDashboardCredential: &armazurearcdata.BasicLoginInformation{
					Password: to.Ptr("<password>"),
					Username: to.Ptr("<username>"),
				},
				OnPremiseProperty: &armazurearcdata.OnPremiseProperty{
					ID:               to.Ptr("<id>"),
					PublicSigningKey: to.Ptr("<public-signing-key>"),
				},
				UploadServicePrincipal: &armazurearcdata.UploadServicePrincipal{
					Authority:    to.Ptr("<authority>"),
					ClientID:     to.Ptr("<client-id>"),
					ClientSecret: to.Ptr("<client-secret>"),
					TenantID:     to.Ptr("<tenant-id>"),
				},
				UploadWatermark: &armazurearcdata.UploadWatermark{
					Logs:    to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t }()),
					Metrics: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t }()),
					Usages:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t }()),
				},
			},
		},
		&armazurearcdata.DataControllersClientBeginPutDataControllerOptions{ResumeToken: ""})
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
