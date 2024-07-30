package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a651ba25cda4eec698a3a4e35f867ecc2681d126/specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/ApplicationPatchOperation_example.json
func ExampleApplicationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationsClient().Update(ctx, "resRg", "myCluster", "myApp", armservicefabricmanagedclusters.ApplicationUpdateParameters{
		Tags: map[string]*string{
			"a": to.Ptr("b"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationResource = armservicefabricmanagedclusters.ApplicationResource{
	// 	Name: to.Ptr("myApp"),
	// 	Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/applications"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applications/myApp"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"a": to.Ptr("b"),
	// 	},
	// 	Properties: &armservicefabricmanagedclusters.ApplicationResourceProperties{
	// 		Parameters: map[string]*string{
	// 			"param1": to.Ptr("value1"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		UpgradePolicy: &armservicefabricmanagedclusters.ApplicationUpgradePolicy{
	// 			ApplicationHealthPolicy: &armservicefabricmanagedclusters.ApplicationHealthPolicy{
	// 				ConsiderWarningAsError: to.Ptr(true),
	// 				DefaultServiceTypeHealthPolicy: &armservicefabricmanagedclusters.ServiceTypeHealthPolicy{
	// 					MaxPercentUnhealthyPartitionsPerService: to.Ptr[int32](0),
	// 					MaxPercentUnhealthyReplicasPerPartition: to.Ptr[int32](0),
	// 					MaxPercentUnhealthyServices: to.Ptr[int32](0),
	// 				},
	// 				MaxPercentUnhealthyDeployedApplications: to.Ptr[int32](0),
	// 				ServiceTypeHealthPolicyMap: map[string]*armservicefabricmanagedclusters.ServiceTypeHealthPolicy{
	// 					"service1": &armservicefabricmanagedclusters.ServiceTypeHealthPolicy{
	// 						MaxPercentUnhealthyPartitionsPerService: to.Ptr[int32](30),
	// 						MaxPercentUnhealthyReplicasPerPartition: to.Ptr[int32](30),
	// 						MaxPercentUnhealthyServices: to.Ptr[int32](30),
	// 					},
	// 				},
	// 			},
	// 			ForceRestart: to.Ptr(false),
	// 			InstanceCloseDelayDuration: to.Ptr[int64](600),
	// 			RecreateApplication: to.Ptr(false),
	// 			RollingUpgradeMonitoringPolicy: &armservicefabricmanagedclusters.RollingUpgradeMonitoringPolicy{
	// 				FailureAction: to.Ptr(armservicefabricmanagedclusters.FailureActionRollback),
	// 				HealthCheckRetryTimeout: to.Ptr("00:10:00"),
	// 				HealthCheckStableDuration: to.Ptr("00:05:00"),
	// 				HealthCheckWaitDuration: to.Ptr("00:02:00"),
	// 				UpgradeDomainTimeout: to.Ptr("00:15:00"),
	// 				UpgradeTimeout: to.Ptr("01:00:00"),
	// 			},
	// 			UpgradeMode: to.Ptr(armservicefabricmanagedclusters.RollingUpgradeModeUnmonitoredAuto),
	// 			UpgradeReplicaSetCheckTimeout: to.Ptr[int64](3600),
	// 		},
	// 		Version: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applicationTypes/myAppType/versions/1.0"),
	// 	},
	// }
}
