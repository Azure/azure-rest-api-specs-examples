Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhdinsight%2Farmhdinsight%2Fv1.0.0/sdk/resourcemanager/hdinsight/armhdinsight/README.md) on how to add the SDK to your project and authenticate.

```go
package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/CreateHDInsightClusterWithAutoscaleConfig.json
func ExampleClustersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhdinsight.NewClustersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"rg1",
		"cluster1",
		armhdinsight.ClusterCreateParametersExtended{
			Properties: &armhdinsight.ClusterCreateProperties{
				ClusterDefinition: &armhdinsight.ClusterDefinition{
					ComponentVersion: map[string]*string{
						"Hadoop": to.Ptr("2.7"),
					},
					Configurations: map[string]interface{}{
						"gateway": map[string]interface{}{
							"restAuthCredential.isEnabled": true,
							"restAuthCredential.password":  "**********",
							"restAuthCredential.username":  "admin",
						},
					},
					Kind: to.Ptr("hadoop"),
				},
				ClusterVersion: to.Ptr("3.6"),
				ComputeProfile: &armhdinsight.ComputeProfile{
					Roles: []*armhdinsight.Role{
						{
							Name: to.Ptr("workernode"),
							AutoscaleConfiguration: &armhdinsight.Autoscale{
								Recurrence: &armhdinsight.AutoscaleRecurrence{
									Schedule: []*armhdinsight.AutoscaleSchedule{
										{
											Days: []*armhdinsight.DaysOfWeek{
												to.Ptr(armhdinsight.DaysOfWeekMonday),
												to.Ptr(armhdinsight.DaysOfWeekTuesday),
												to.Ptr(armhdinsight.DaysOfWeekWednesday),
												to.Ptr(armhdinsight.DaysOfWeekThursday),
												to.Ptr(armhdinsight.DaysOfWeekFriday)},
											TimeAndCapacity: &armhdinsight.AutoscaleTimeAndCapacity{
												MaxInstanceCount: to.Ptr[int32](3),
												MinInstanceCount: to.Ptr[int32](3),
												Time:             to.Ptr("09:00"),
											},
										},
										{
											Days: []*armhdinsight.DaysOfWeek{
												to.Ptr(armhdinsight.DaysOfWeekMonday),
												to.Ptr(armhdinsight.DaysOfWeekTuesday),
												to.Ptr(armhdinsight.DaysOfWeekWednesday),
												to.Ptr(armhdinsight.DaysOfWeekThursday),
												to.Ptr(armhdinsight.DaysOfWeekFriday)},
											TimeAndCapacity: &armhdinsight.AutoscaleTimeAndCapacity{
												MaxInstanceCount: to.Ptr[int32](6),
												MinInstanceCount: to.Ptr[int32](6),
												Time:             to.Ptr("18:00"),
											},
										},
										{
											Days: []*armhdinsight.DaysOfWeek{
												to.Ptr(armhdinsight.DaysOfWeekSaturday),
												to.Ptr(armhdinsight.DaysOfWeekSunday)},
											TimeAndCapacity: &armhdinsight.AutoscaleTimeAndCapacity{
												MaxInstanceCount: to.Ptr[int32](2),
												MinInstanceCount: to.Ptr[int32](2),
												Time:             to.Ptr("09:00"),
											},
										},
										{
											Days: []*armhdinsight.DaysOfWeek{
												to.Ptr(armhdinsight.DaysOfWeekSaturday),
												to.Ptr(armhdinsight.DaysOfWeekSunday)},
											TimeAndCapacity: &armhdinsight.AutoscaleTimeAndCapacity{
												MaxInstanceCount: to.Ptr[int32](4),
												MinInstanceCount: to.Ptr[int32](4),
												Time:             to.Ptr("18:00"),
											},
										}},
									TimeZone: to.Ptr("China Standard Time"),
								},
							},
							HardwareProfile: &armhdinsight.HardwareProfile{
								VMSize: to.Ptr("Standard_D4_V2"),
							},
							OSProfile: &armhdinsight.OsProfile{
								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
									Password: to.Ptr("**********"),
									Username: to.Ptr("sshuser"),
								},
							},
							ScriptActions:       []*armhdinsight.ScriptAction{},
							TargetInstanceCount: to.Ptr[int32](4),
						}},
				},
				OSType: to.Ptr(armhdinsight.OSTypeLinux),
				StorageProfile: &armhdinsight.StorageProfile{
					Storageaccounts: []*armhdinsight.StorageAccount{
						{
							Name:      to.Ptr("mystorage.blob.core.windows.net"),
							Container: to.Ptr("hdinsight-autoscale-tes-2019-06-18t05-49-16-591z"),
							IsDefault: to.Ptr(true),
							Key:       to.Ptr("storagekey"),
						}},
				},
				Tier: to.Ptr(armhdinsight.TierStandard),
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
