package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationPutOperation_example_max.json
func ExampleApplicationsClient_BeginCreateOrUpdate_putAnApplicationWithMaximumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", "myApp", armservicefabric.ApplicationResource{
		Tags: map[string]*string{},
		Properties: &armservicefabric.ApplicationResourceProperties{
			MaximumNodes: to.Ptr[int64](3),
			Metrics: []*armservicefabric.ApplicationMetricDescription{
				{
					Name:                     to.Ptr("metric1"),
					MaximumCapacity:          to.Ptr[int64](3),
					ReservationCapacity:      to.Ptr[int64](1),
					TotalApplicationCapacity: to.Ptr[int64](5),
				}},
			MinimumNodes: to.Ptr[int64](1),
			Parameters: map[string]*string{
				"param1": to.Ptr("value1"),
			},
			RemoveApplicationCapacity: to.Ptr(false),
			TypeVersion:               to.Ptr("1.0"),
			UpgradePolicy: &armservicefabric.ApplicationUpgradePolicy{
				ApplicationHealthPolicy: &armservicefabric.ArmApplicationHealthPolicy{
					ConsiderWarningAsError: to.Ptr(true),
					DefaultServiceTypeHealthPolicy: &armservicefabric.ArmServiceTypeHealthPolicy{
						MaxPercentUnhealthyPartitionsPerService: to.Ptr[int32](0),
						MaxPercentUnhealthyReplicasPerPartition: to.Ptr[int32](0),
						MaxPercentUnhealthyServices:             to.Ptr[int32](0),
					},
					MaxPercentUnhealthyDeployedApplications: to.Ptr[int32](0),
				},
				ForceRestart: to.Ptr(false),
				RollingUpgradeMonitoringPolicy: &armservicefabric.ArmRollingUpgradeMonitoringPolicy{
					FailureAction:             to.Ptr(armservicefabric.ArmUpgradeFailureActionRollback),
					HealthCheckRetryTimeout:   to.Ptr("00:10:00"),
					HealthCheckStableDuration: to.Ptr("00:05:00"),
					HealthCheckWaitDuration:   to.Ptr("00:02:00"),
					UpgradeDomainTimeout:      to.Ptr("1.06:00:00"),
					UpgradeTimeout:            to.Ptr("01:00:00"),
				},
				UpgradeMode:                   to.Ptr(armservicefabric.RollingUpgradeModeMonitored),
				UpgradeReplicaSetCheckTimeout: to.Ptr("01:00:00"),
			},
			TypeName: to.Ptr("myAppType"),
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
