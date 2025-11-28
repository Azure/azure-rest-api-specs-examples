package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/ApplicationActionUpdateUpgrade_example.json
func ExampleApplicationsClient_BeginUpdateUpgrade() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginUpdateUpgrade(ctx, "resRg", "myCluster", "myApp", armservicefabricmanagedclusters.RuntimeUpdateApplicationUpgradeParameters{
		UpgradeKind: to.Ptr(armservicefabricmanagedclusters.RuntimeUpgradeKindRolling),
		Name:        to.Ptr("fabric:/Voting"),
		UpdateDescription: &armservicefabricmanagedclusters.RuntimeRollingUpgradeUpdateMonitoringPolicy{
			RollingUpgradeMode:                      to.Ptr(armservicefabricmanagedclusters.RuntimeRollingUpgradeModeMonitored),
			ForceRestart:                            to.Ptr(true),
			FailureAction:                           to.Ptr(armservicefabricmanagedclusters.RuntimeFailureActionManual),
			HealthCheckWaitDurationInMilliseconds:   to.Ptr("PT0H0M10S"),
			HealthCheckStableDurationInMilliseconds: to.Ptr("PT1H0M0S"),
			HealthCheckRetryTimeoutInMilliseconds:   to.Ptr("PT0H15M0S"),
			UpgradeTimeoutInMilliseconds:            to.Ptr("PT2H0M0S"),
			UpgradeDomainTimeoutInMilliseconds:      to.Ptr("PT2H0M0S"),
		},
		ApplicationHealthPolicy: &armservicefabricmanagedclusters.RuntimeApplicationHealthPolicy{
			ConsiderWarningAsError:                  to.Ptr(true),
			MaxPercentUnhealthyDeployedApplications: to.Ptr[int32](10),
			DefaultServiceTypeHealthPolicy: &armservicefabricmanagedclusters.RuntimeServiceTypeHealthPolicy{
				MaxPercentUnhealthyPartitionsPerService: to.Ptr[int32](10),
				MaxPercentUnhealthyReplicasPerPartition: to.Ptr[int32](11),
				MaxPercentUnhealthyServices:             to.Ptr[int32](12),
			},
			ServiceTypeHealthPolicyMap: map[string]*armservicefabricmanagedclusters.RuntimeServiceTypeHealthPolicy{
				"VotingWeb": {
					MaxPercentUnhealthyPartitionsPerService: to.Ptr[int32](13),
					MaxPercentUnhealthyReplicasPerPartition: to.Ptr[int32](14),
					MaxPercentUnhealthyServices:             to.Ptr[int32](15),
				},
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
