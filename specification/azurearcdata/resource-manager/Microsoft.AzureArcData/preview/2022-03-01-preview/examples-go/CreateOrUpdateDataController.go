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
	}
	ctx := context.Background()
	client, err := armazurearcdata.NewDataControllersClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPutDataController(ctx,
		"testrg",
		"testdataController",
		armazurearcdata.DataControllerResource{
			Location: to.Ptr("northeurope"),
			Tags: map[string]*string{
				"mytag": to.Ptr("myval"),
			},
			ExtendedLocation: &armazurearcdata.ExtendedLocation{
				Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
				Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
			},
			Properties: &armazurearcdata.DataControllerProperties{
				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
					Password: to.Ptr("********"),
					Username: to.Ptr("username"),
				},
				ClusterID:      to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s"),
				ExtensionID:    to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"),
				Infrastructure: to.Ptr(armazurearcdata.InfrastructureOnpremises),
				LogAnalyticsWorkspaceConfig: &armazurearcdata.LogAnalyticsWorkspaceConfig{
					PrimaryKey:  to.Ptr("********"),
					WorkspaceID: to.Ptr("00000000-1111-2222-3333-444444444444"),
				},
				LogsDashboardCredential: &armazurearcdata.BasicLoginInformation{
					Password: to.Ptr("********"),
					Username: to.Ptr("username"),
				},
				MetricsDashboardCredential: &armazurearcdata.BasicLoginInformation{
					Password: to.Ptr("********"),
					Username: to.Ptr("username"),
				},
				OnPremiseProperty: &armazurearcdata.OnPremiseProperty{
					ID:               to.Ptr("12345678-1234-1234-ab12-1a2b3c4d5e6f"),
					PublicSigningKey: to.Ptr("publicOnPremSigningKey"),
				},
				UploadServicePrincipal: &armazurearcdata.UploadServicePrincipal{
					Authority:    to.Ptr("https://login.microsoftonline.com/"),
					ClientID:     to.Ptr("00000000-1111-2222-3333-444444444444"),
					ClientSecret: to.Ptr("********"),
					TenantID:     to.Ptr("00000000-1111-2222-3333-444444444444"),
				},
				UploadWatermark: &armazurearcdata.UploadWatermark{
					Logs:    to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t }()),
					Metrics: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t }()),
					Usages:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t }()),
				},
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
