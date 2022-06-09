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

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/CreateOrUpdateSqlManagedInstance.json
func ExampleSQLManagedInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLManagedInstancesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<sql-managed-instance-name>",
		armazurearcdata.SQLManagedInstance{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"mytag": to.StringPtr("myval"),
			},
			ExtendedLocation: &armazurearcdata.ExtendedLocation{
				Name: to.StringPtr("<name>"),
				Type: armazurearcdata.ExtendedLocationTypes("CustomLocation").ToPtr(),
			},
			Properties: &armazurearcdata.SQLManagedInstanceProperties{
				Admin: to.StringPtr("<admin>"),
				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
					Password: to.StringPtr("<password>"),
					Username: to.StringPtr("<username>"),
				},
				ClusterID:   to.StringPtr("<cluster-id>"),
				EndTime:     to.StringPtr("<end-time>"),
				ExtensionID: to.StringPtr("<extension-id>"),
				K8SRaw: &armazurearcdata.SQLManagedInstanceK8SRaw{
					AdditionalProperties: map[string]map[string]interface{}{
						"additionalProperty": {},
					},
					Spec: &armazurearcdata.SQLManagedInstanceK8SSpec{
						Replicas: to.Int32Ptr(1),
						Scheduling: &armazurearcdata.K8SScheduling{
							Default: &armazurearcdata.K8SSchedulingOptions{
								Resources: &armazurearcdata.K8SResourceRequirements{
									Limits: map[string]*string{
										"additionalProperty": to.StringPtr("additionalValue"),
										"cpu":                to.StringPtr("1"),
										"memory":             to.StringPtr("8Gi"),
									},
									Requests: map[string]*string{
										"additionalProperty": to.StringPtr("additionalValue"),
										"cpu":                to.StringPtr("1"),
										"memory":             to.StringPtr("8Gi"),
									},
								},
							},
						},
					},
				},
				LicenseType: armazurearcdata.ArcSQLManagedInstanceLicenseType("LicenseIncluded").ToPtr(),
				StartTime:   to.StringPtr("<start-time>"),
			},
			SKU: &armazurearcdata.SQLManagedInstanceSKU{
				Name: armazurearcdata.SQLManagedInstanceSKUName("vCore").ToPtr(),
				Dev:  to.BoolPtr(true),
				Tier: armazurearcdata.SQLManagedInstanceSKUTierGeneralPurpose.ToPtr(),
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
	log.Printf("Response result: %#v\n", res.SQLManagedInstancesClientCreateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazurearcdata%2Farmazurearcdata%2Fv0.2.0/sdk/resourcemanager/azurearcdata/armazurearcdata/README.md) on how to add the SDK to your project and authenticate.
