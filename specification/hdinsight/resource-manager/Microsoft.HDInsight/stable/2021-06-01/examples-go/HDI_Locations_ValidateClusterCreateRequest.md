Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhdinsight%2Farmhdinsight%2Fv0.2.1/sdk/resourcemanager/hdinsight/armhdinsight/README.md) on how to add the SDK to your project and authenticate.

```go
package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/HDI_Locations_ValidateClusterCreateRequest.json
func ExampleLocationsClient_ValidateClusterCreateRequest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewLocationsClient("<subscription-id>", cred, nil)
	res, err := client.ValidateClusterCreateRequest(ctx,
		"<location>",
		armhdinsight.ClusterCreateRequestValidationParameters{
			Location: to.StringPtr("<location>"),
			Properties: &armhdinsight.ClusterCreateProperties{
				ClusterDefinition: &armhdinsight.ClusterDefinition{
					ComponentVersion: map[string]*string{
						"Spark": to.StringPtr("2.4"),
					},
					Configurations: map[string]interface{}{
						"gateway": map[string]interface{}{
							"restAuthCredential.isEnabled": true,
							"restAuthCredential.password":  "**********",
							"restAuthCredential.username":  "admin",
						},
					},
					Kind: to.StringPtr("<kind>"),
				},
				ClusterVersion: to.StringPtr("<cluster-version>"),
				ComputeProfile: &armhdinsight.ComputeProfile{
					Roles: []*armhdinsight.Role{
						{
							Name: to.StringPtr("<name>"),
							HardwareProfile: &armhdinsight.HardwareProfile{
								VMSize: to.StringPtr("<vmsize>"),
							},
							MinInstanceCount: to.Int32Ptr(1),
							OSProfile: &armhdinsight.OsProfile{
								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
									Password: to.StringPtr("<password>"),
									Username: to.StringPtr("<username>"),
								},
							},
							ScriptActions:       []*armhdinsight.ScriptAction{},
							TargetInstanceCount: to.Int32Ptr(2),
						},
						{
							Name: to.StringPtr("<name>"),
							HardwareProfile: &armhdinsight.HardwareProfile{
								VMSize: to.StringPtr("<vmsize>"),
							},
							OSProfile: &armhdinsight.OsProfile{
								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
									Password: to.StringPtr("<password>"),
									Username: to.StringPtr("<username>"),
								},
							},
							ScriptActions:       []*armhdinsight.ScriptAction{},
							TargetInstanceCount: to.Int32Ptr(4),
						},
						{
							Name: to.StringPtr("<name>"),
							HardwareProfile: &armhdinsight.HardwareProfile{
								VMSize: to.StringPtr("<vmsize>"),
							},
							MinInstanceCount: to.Int32Ptr(1),
							OSProfile: &armhdinsight.OsProfile{
								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
									Password: to.StringPtr("<password>"),
									Username: to.StringPtr("<username>"),
								},
							},
							ScriptActions:       []*armhdinsight.ScriptAction{},
							TargetInstanceCount: to.Int32Ptr(3),
						}},
				},
				MinSupportedTLSVersion: to.StringPtr("<min-supported-tlsversion>"),
				OSType:                 armhdinsight.OSType("Linux").ToPtr(),
				StorageProfile: &armhdinsight.StorageProfile{
					Storageaccounts: []*armhdinsight.StorageAccount{
						{
							Name:       to.StringPtr("<name>"),
							Container:  to.StringPtr("<container>"),
							IsDefault:  to.BoolPtr(true),
							Key:        to.StringPtr("<key>"),
							ResourceID: to.StringPtr("<resource-id>"),
						}},
				},
				Tier: armhdinsight.Tier("Standard").ToPtr(),
			},
			Tags:               map[string]*string{},
			Name:               to.StringPtr("<name>"),
			Type:               to.StringPtr("<type>"),
			FetchAaddsResource: to.BoolPtr(false),
			TenantID:           to.StringPtr("<tenant-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LocationsClientValidateClusterCreateRequestResult)
}
```
