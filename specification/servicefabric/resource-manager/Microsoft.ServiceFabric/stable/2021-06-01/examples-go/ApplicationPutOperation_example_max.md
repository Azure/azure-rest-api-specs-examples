Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicefabric%2Farmservicefabric%2Fv0.6.0/sdk/resourcemanager/servicefabric/armservicefabric/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicefabric_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationPutOperation_example_max.json
func ExampleApplicationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armservicefabric.NewApplicationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<application-name>",
		armservicefabric.ApplicationResource{
			Tags: map[string]*string{},
			Properties: &armservicefabric.ApplicationResourceProperties{
				MaximumNodes: to.Ptr[int64](3),
				Metrics: []*armservicefabric.ApplicationMetricDescription{
					{
						Name:                     to.Ptr("<name>"),
						MaximumCapacity:          to.Ptr[int64](3),
						ReservationCapacity:      to.Ptr[int64](1),
						TotalApplicationCapacity: to.Ptr[int64](5),
					}},
				MinimumNodes: to.Ptr[int64](1),
				Parameters: map[string]*string{
					"param1": to.Ptr("value1"),
				},
				RemoveApplicationCapacity: to.Ptr(false),
				TypeVersion:               to.Ptr("<type-version>"),
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
						HealthCheckRetryTimeout:   to.Ptr("<health-check-retry-timeout>"),
						HealthCheckStableDuration: to.Ptr("<health-check-stable-duration>"),
						HealthCheckWaitDuration:   to.Ptr("<health-check-wait-duration>"),
						UpgradeDomainTimeout:      to.Ptr("<upgrade-domain-timeout>"),
						UpgradeTimeout:            to.Ptr("<upgrade-timeout>"),
					},
					UpgradeMode:                   to.Ptr(armservicefabric.RollingUpgradeModeMonitored),
					UpgradeReplicaSetCheckTimeout: to.Ptr("<upgrade-replica-set-check-timeout>"),
				},
				TypeName: to.Ptr("<type-name>"),
			},
		},
		&armservicefabric.ApplicationsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```
