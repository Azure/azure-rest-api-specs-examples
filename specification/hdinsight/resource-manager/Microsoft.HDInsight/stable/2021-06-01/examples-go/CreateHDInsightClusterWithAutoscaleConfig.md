Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhdinsight%2Farmhdinsight%2Fv0.4.0/sdk/resourcemanager/hdinsight/armhdinsight/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/CreateHDInsightClusterWithAutoscaleConfig.json
func ExampleClustersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhdinsight.NewClustersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
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
					Kind: to.Ptr("<kind>"),
				},
				ClusterVersion: to.Ptr("<cluster-version>"),
				ComputeProfile: &armhdinsight.ComputeProfile{
					Roles: []*armhdinsight.Role{
						{
							Name: to.Ptr("<name>"),
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
												Time:             to.Ptr("<time>"),
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
												Time:             to.Ptr("<time>"),
											},
										},
										{
											Days: []*armhdinsight.DaysOfWeek{
												to.Ptr(armhdinsight.DaysOfWeekSaturday),
												to.Ptr(armhdinsight.DaysOfWeekSunday)},
											TimeAndCapacity: &armhdinsight.AutoscaleTimeAndCapacity{
												MaxInstanceCount: to.Ptr[int32](2),
												MinInstanceCount: to.Ptr[int32](2),
												Time:             to.Ptr("<time>"),
											},
										},
										{
											Days: []*armhdinsight.DaysOfWeek{
												to.Ptr(armhdinsight.DaysOfWeekSaturday),
												to.Ptr(armhdinsight.DaysOfWeekSunday)},
											TimeAndCapacity: &armhdinsight.AutoscaleTimeAndCapacity{
												MaxInstanceCount: to.Ptr[int32](4),
												MinInstanceCount: to.Ptr[int32](4),
												Time:             to.Ptr("<time>"),
											},
										}},
									TimeZone: to.Ptr("<time-zone>"),
								},
							},
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
						}},
				},
				OSType: to.Ptr(armhdinsight.OSTypeLinux),
				StorageProfile: &armhdinsight.StorageProfile{
					Storageaccounts: []*armhdinsight.StorageAccount{
						{
							Name:      to.Ptr("<name>"),
							Container: to.Ptr("<container>"),
							IsDefault: to.Ptr(true),
							Key:       to.Ptr("<key>"),
						}},
				},
				Tier: to.Ptr(armhdinsight.TierStandard),
			},
		},
		&armhdinsight.ClustersClientBeginCreateOptions{ResumeToken: ""})
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
