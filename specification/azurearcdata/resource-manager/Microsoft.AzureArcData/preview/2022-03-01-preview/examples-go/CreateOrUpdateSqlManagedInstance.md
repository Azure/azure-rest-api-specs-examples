```go
package armazurearcdata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/CreateOrUpdateSqlManagedInstance.json
func ExampleSQLManagedInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armazurearcdata.NewSQLManagedInstancesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"testrg",
		"testsqlManagedInstance",
		armazurearcdata.SQLManagedInstance{
			Location: to.Ptr("northeurope"),
			Tags: map[string]*string{
				"mytag": to.Ptr("myval"),
			},
			ExtendedLocation: &armazurearcdata.ExtendedLocation{
				Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
				Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
			},
			Properties: &armazurearcdata.SQLManagedInstanceProperties{
				ActiveDirectoryInformation: &armazurearcdata.ActiveDirectoryInformation{
					KeytabInformation: &armazurearcdata.KeytabInformation{
						Keytab: to.Ptr("********"),
					},
				},
				Admin: to.Ptr("Admin user"),
				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
					Password: to.Ptr("********"),
					Username: to.Ptr("username"),
				},
				ClusterID:   to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s"),
				EndTime:     to.Ptr("Instance end time"),
				ExtensionID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"),
				K8SRaw: &armazurearcdata.SQLManagedInstanceK8SRaw{
					AdditionalProperties: map[string]interface{}{
						"additionalProperty": float64(1234),
					},
					Spec: &armazurearcdata.SQLManagedInstanceK8SSpec{
						Replicas: to.Ptr[int32](1),
						Scheduling: &armazurearcdata.K8SScheduling{
							Default: &armazurearcdata.K8SSchedulingOptions{
								Resources: &armazurearcdata.K8SResourceRequirements{
									Limits: map[string]*string{
										"additionalProperty": to.Ptr("additionalValue"),
										"cpu":                to.Ptr("1"),
										"memory":             to.Ptr("8Gi"),
									},
									Requests: map[string]*string{
										"additionalProperty": to.Ptr("additionalValue"),
										"cpu":                to.Ptr("1"),
										"memory":             to.Ptr("8Gi"),
									},
								},
							},
						},
					},
				},
				LicenseType: to.Ptr(armazurearcdata.ArcSQLManagedInstanceLicenseTypeLicenseIncluded),
				StartTime:   to.Ptr("Instance start time"),
			},
			SKU: &armazurearcdata.SQLManagedInstanceSKU{
				Name: to.Ptr("vCore"),
				Dev:  to.Ptr(true),
				Tier: to.Ptr(armazurearcdata.SQLManagedInstanceSKUTierGeneralPurpose),
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
