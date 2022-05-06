Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhdinsight%2Farmhdinsight%2Fv0.4.0/sdk/resourcemanager/hdinsight/armhdinsight/README.md) on how to add the SDK to your project and authenticate.

```go
package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/HDI_Locations_ValidateClusterCreateRequest.json
func ExampleLocationsClient_ValidateClusterCreateRequest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhdinsight.NewLocationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.ValidateClusterCreateRequest(ctx,
		"<location>",
		armhdinsight.ClusterCreateRequestValidationParameters{
			Location: to.Ptr("<location>"),
			Properties: &armhdinsight.ClusterCreateProperties{
				ClusterDefinition: &armhdinsight.ClusterDefinition{
					ComponentVersion: map[string]*string{
						"Spark": to.Ptr("2.4"),
					},
					Configurations: map[string]interface{}{
						"gateway": map[string]interface{}{
							"restAuthCredential.isEnabled": true,
							"restAuthCredential.password":  "**********",
							"restAuthCredential.username":  "admin",
						},
					},
					Kind: to.Ptr("<kind>"),
				},
				ClusterVersion: to.Ptr("<cluster-version>"),
				ComputeProfile: &armhdinsight.ComputeProfile{
					Roles: []*armhdinsight.Role{
						{
							Name: to.Ptr("<name>"),
							HardwareProfile: &armhdinsight.HardwareProfile{
								VMSize: to.Ptr("<vmsize>"),
							},
							MinInstanceCount: to.Ptr[int32](1),
							OSProfile: &armhdinsight.OsProfile{
								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
									Password: to.Ptr("<password>"),
									Username: to.Ptr("<username>"),
								},
							},
							ScriptActions:       []*armhdinsight.ScriptAction{},
							TargetInstanceCount: to.Ptr[int32](2),
						},
						{
							Name: to.Ptr("<name>"),
							HardwareProfile: &armhdinsight.HardwareProfile{
								VMSize: to.Ptr("<vmsize>"),
							},
							OSProfile: &armhdinsight.OsProfile{
								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
									Password: to.Ptr("<password>"),
									Username: to.Ptr("<username>"),
								},
							},
							ScriptActions:       []*armhdinsight.ScriptAction{},
							TargetInstanceCount: to.Ptr[int32](4),
						},
						{
							Name: to.Ptr("<name>"),
							HardwareProfile: &armhdinsight.HardwareProfile{
								VMSize: to.Ptr("<vmsize>"),
							},
							MinInstanceCount: to.Ptr[int32](1),
							OSProfile: &armhdinsight.OsProfile{
								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
									Password: to.Ptr("<password>"),
									Username: to.Ptr("<username>"),
								},
							},
							ScriptActions:       []*armhdinsight.ScriptAction{},
							TargetInstanceCount: to.Ptr[int32](3),
						}},
				},
				MinSupportedTLSVersion: to.Ptr("<min-supported-tlsversion>"),
				OSType:                 to.Ptr(armhdinsight.OSTypeLinux),
				StorageProfile: &armhdinsight.StorageProfile{
					Storageaccounts: []*armhdinsight.StorageAccount{
						{
							Name:       to.Ptr("<name>"),
							Container:  to.Ptr("<container>"),
							IsDefault:  to.Ptr(true),
							Key:        to.Ptr("<key>"),
							ResourceID: to.Ptr("<resource-id>"),
						}},
				},
				Tier: to.Ptr(armhdinsight.TierStandard),
			},
			Tags:               map[string]*string{},
			Name:               to.Ptr("<name>"),
			Type:               to.Ptr("<type>"),
			FetchAaddsResource: to.Ptr(false),
			TenantID:           to.Ptr("<tenant-id>"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
