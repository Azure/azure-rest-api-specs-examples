package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/EnableOrUpdateAutoScaleWithScheduleBasedConfiguration.json
func ExampleClustersClient_BeginUpdateAutoScaleConfiguration_enableOrUpdateAutoscaleWithTheScheduleBasedConfigurationForHdInsightCluster() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginUpdateAutoScaleConfiguration(ctx, "rg1", "cluster1", armhdinsight.RoleNameWorkernode, armhdinsight.AutoscaleConfigurationUpdateParameter{
		Autoscale: &armhdinsight.Autoscale{
			Recurrence: &armhdinsight.AutoscaleRecurrence{
				Schedule: []*armhdinsight.AutoscaleSchedule{
					{
						Days: []*armhdinsight.DaysOfWeek{
							to.Ptr(armhdinsight.DaysOfWeekThursday)},
						TimeAndCapacity: &armhdinsight.AutoscaleTimeAndCapacity{
							MaxInstanceCount: to.Ptr[int32](4),
							MinInstanceCount: to.Ptr[int32](4),
							Time:             to.Ptr("16:00"),
						},
					}},
				TimeZone: to.Ptr("China Standard Time"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
