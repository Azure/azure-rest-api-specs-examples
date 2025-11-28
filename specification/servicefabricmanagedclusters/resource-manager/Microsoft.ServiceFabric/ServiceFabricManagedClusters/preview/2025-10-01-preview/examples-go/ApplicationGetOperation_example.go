package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/ApplicationGetOperation_example.json
func ExampleApplicationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationsClient().Get(ctx, "resRg", "myCluster", "myApp", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicefabricmanagedclusters.ApplicationsClientGetResponse{
	// 	ApplicationResource: &armservicefabricmanagedclusters.ApplicationResource{
	// 		Name: to.Ptr("myApp"),
	// 		Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/applications"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applications/myApp"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armservicefabricmanagedclusters.ApplicationResourceProperties{
	// 			Parameters: map[string]*string{
	// 				"param1": to.Ptr("value1"),
	// 			},
	// 			ProvisioningState: to.Ptr("Updating"),
	// 			UpgradePolicy: &armservicefabricmanagedclusters.ApplicationUpgradePolicy{
	// 				ApplicationHealthPolicy: &armservicefabricmanagedclusters.ApplicationHealthPolicy{
	// 					ConsiderWarningAsError: to.Ptr(true),
	// 					DefaultServiceTypeHealthPolicy: &armservicefabricmanagedclusters.ServiceTypeHealthPolicy{
	// 						MaxPercentUnhealthyPartitionsPerService: to.Ptr[int32](0),
	// 						MaxPercentUnhealthyReplicasPerPartition: to.Ptr[int32](0),
	// 						MaxPercentUnhealthyServices: to.Ptr[int32](0),
	// 					},
	// 					MaxPercentUnhealthyDeployedApplications: to.Ptr[int32](0),
	// 					ServiceTypeHealthPolicyMap: map[string]*armservicefabricmanagedclusters.ServiceTypeHealthPolicy{
	// 						"service1": &armservicefabricmanagedclusters.ServiceTypeHealthPolicy{
	// 							MaxPercentUnhealthyPartitionsPerService: to.Ptr[int32](30),
	// 							MaxPercentUnhealthyReplicasPerPartition: to.Ptr[int32](30),
	// 							MaxPercentUnhealthyServices: to.Ptr[int32](30),
	// 						},
	// 					},
	// 				},
	// 				ForceRestart: to.Ptr(false),
	// 				InstanceCloseDelayDuration: to.Ptr[int64](600),
	// 				RecreateApplication: to.Ptr(false),
	// 				RollingUpgradeMonitoringPolicy: &armservicefabricmanagedclusters.RollingUpgradeMonitoringPolicy{
	// 					FailureAction: to.Ptr(armservicefabricmanagedclusters.FailureActionRollback),
	// 					HealthCheckRetryTimeout: to.Ptr("00:10:00"),
	// 					HealthCheckStableDuration: to.Ptr("00:05:00"),
	// 					HealthCheckWaitDuration: to.Ptr("00:02:00"),
	// 					UpgradeDomainTimeout: to.Ptr("00:15:00"),
	// 					UpgradeTimeout: to.Ptr("01:00:00"),
	// 				},
	// 				UpgradeMode: to.Ptr(armservicefabricmanagedclusters.RollingUpgradeModeUnmonitoredAuto),
	// 				UpgradeReplicaSetCheckTimeout: to.Ptr[int64](3600),
	// 			},
	// 			Version: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applicationTypes/myAppType/versions/1.0"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
