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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/CreateOrUpdateSqlManagedInstance.json
func ExampleSQLManagedInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armazurearcdata.NewSQLManagedInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<sql-managed-instance-name>",
		armazurearcdata.SQLManagedInstance{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"mytag": to.Ptr("myval"),
			},
			ExtendedLocation: &armazurearcdata.ExtendedLocation{
				Name: to.Ptr("<name>"),
				Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
			},
			Properties: &armazurearcdata.SQLManagedInstanceProperties{
				ActiveDirectoryInformation: &armazurearcdata.ActiveDirectoryInformation{
					KeytabInformation: &armazurearcdata.KeytabInformation{
						Keytab: to.Ptr("<keytab>"),
					},
				},
				Admin: to.Ptr("<admin>"),
				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
					Password: to.Ptr("<password>"),
					Username: to.Ptr("<username>"),
				},
				ClusterID:   to.Ptr("<cluster-id>"),
				EndTime:     to.Ptr("<end-time>"),
				ExtensionID: to.Ptr("<extension-id>"),
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
				StartTime:   to.Ptr("<start-time>"),
			},
			SKU: &armazurearcdata.SQLManagedInstanceSKU{
				Name: to.Ptr("<name>"),
				Dev:  to.Ptr(true),
				Tier: to.Ptr(armazurearcdata.SQLManagedInstanceSKUTierGeneralPurpose),
			},
		},
		&armazurearcdata.SQLManagedInstancesClientBeginCreateOptions{ResumeToken: ""})
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
