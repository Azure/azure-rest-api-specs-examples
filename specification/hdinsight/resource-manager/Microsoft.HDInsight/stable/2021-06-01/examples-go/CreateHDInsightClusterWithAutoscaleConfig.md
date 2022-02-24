Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhdinsight%2Farmhdinsight%2Fv0.2.1/sdk/resourcemanager/hdinsight/armhdinsight/README.md) on how to add the SDK to your project and authenticate.

```go
package armhdinsight_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/CreateHDInsightClusterWithAutoscaleConfig.json
func ExampleClustersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armhdinsight.ClusterCreateParametersExtended{
			Properties: &armhdinsight.ClusterCreateProperties{
				ClusterDefinition: &armhdinsight.ClusterDefinition{
					ComponentVersion: map[string]*string{
						"Hadoop": to.StringPtr("2.7"),
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
							AutoscaleConfiguration: &armhdinsight.Autoscale{
								Recurrence: &armhdinsight.AutoscaleRecurrence{
									Schedule: []*armhdinsight.AutoscaleSchedule{
										{
											Days: []*armhdinsight.DaysOfWeek{
												armhdinsight.DaysOfWeek("Monday").ToPtr(),
												armhdinsight.DaysOfWeek("Tuesday").ToPtr(),
												armhdinsight.DaysOfWeek("Wednesday").ToPtr(),
												armhdinsight.DaysOfWeek("Thursday").ToPtr(),
												armhdinsight.DaysOfWeek("Friday").ToPtr()},
											TimeAndCapacity: &armhdinsight.AutoscaleTimeAndCapacity{
												MaxInstanceCount: to.Int32Ptr(3),
												MinInstanceCount: to.Int32Ptr(3),
												Time:             to.StringPtr("<time>"),
											},
										},
										{
											Days: []*armhdinsight.DaysOfWeek{
												armhdinsight.DaysOfWeek("Monday").ToPtr(),
												armhdinsight.DaysOfWeek("Tuesday").ToPtr(),
												armhdinsight.DaysOfWeek("Wednesday").ToPtr(),
												armhdinsight.DaysOfWeek("Thursday").ToPtr(),
												armhdinsight.DaysOfWeek("Friday").ToPtr()},
											TimeAndCapacity: &armhdinsight.AutoscaleTimeAndCapacity{
												MaxInstanceCount: to.Int32Ptr(6),
												MinInstanceCount: to.Int32Ptr(6),
												Time:             to.StringPtr("<time>"),
											},
										},
										{
											Days: []*armhdinsight.DaysOfWeek{
												armhdinsight.DaysOfWeek("Saturday").ToPtr(),
												armhdinsight.DaysOfWeek("Sunday").ToPtr()},
											TimeAndCapacity: &armhdinsight.AutoscaleTimeAndCapacity{
												MaxInstanceCount: to.Int32Ptr(2),
												MinInstanceCount: to.Int32Ptr(2),
												Time:             to.StringPtr("<time>"),
											},
										},
										{
											Days: []*armhdinsight.DaysOfWeek{
												armhdinsight.DaysOfWeek("Saturday").ToPtr(),
												armhdinsight.DaysOfWeek("Sunday").ToPtr()},
											TimeAndCapacity: &armhdinsight.AutoscaleTimeAndCapacity{
												MaxInstanceCount: to.Int32Ptr(4),
												MinInstanceCount: to.Int32Ptr(4),
												Time:             to.StringPtr("<time>"),
											},
										}},
									TimeZone: to.StringPtr("<time-zone>"),
								},
							},
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
						}},
				},
				OSType: armhdinsight.OSType("Linux").ToPtr(),
				StorageProfile: &armhdinsight.StorageProfile{
					Storageaccounts: []*armhdinsight.StorageAccount{
						{
							Name:      to.StringPtr("<name>"),
							Container: to.StringPtr("<container>"),
							IsDefault: to.BoolPtr(true),
							Key:       to.StringPtr("<key>"),
						}},
				},
				Tier: armhdinsight.Tier("Standard").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.ClustersClientCreateResult)
}
```
