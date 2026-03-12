package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/ClusterJobs_CreateOrUpdate_ConfigureSdnIntegration_Enable.json
func ExampleClusterJobsClient_BeginCreateOrUpdate_clusterJobsCreateOrUpdateConfigureSdnIntegrationEnable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClusterJobsClient().BeginCreateOrUpdate(ctx, "test-rg", "myCluster", "configureSdnIntegration", armazurestackhci.ClusterJob{
		Properties: &armazurestackhci.HciConfigureSdnIntegrationJobProperties{
			JobType:              to.Ptr(armazurestackhci.HciJobTypeConfigureSdnIntegration),
			DeploymentMode:       to.Ptr(armazurestackhci.DeploymentModeDeploy),
			SdnIntegrationIntent: to.Ptr(armazurestackhci.SdnIntegrationIntentEnable),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.ClusterJobsClientCreateOrUpdateResponse{
	// 	ClusterJob: &armazurestackhci.ClusterJob{
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/myCluster/jobs/configureSdnIntegration"),
	// 		Name: to.Ptr("configureSdnIntegration"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/clusters/jobs"),
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedBy: to.Ptr("user@microsoft.com"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T10:43:27.9471574Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T10:43:27.9471574Z"); return t}()),
	// 		},
	// 		Properties: &armazurestackhci.HciConfigureSdnIntegrationJobProperties{
	// 			JobType: to.Ptr(armazurestackhci.HciJobTypeConfigureSdnIntegration),
	// 			DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
	// 			SdnIntegrationIntent: to.Ptr(armazurestackhci.SdnIntegrationIntentEnable),
	// 			SdnPrefix: to.Ptr("sdn01-a"),
	// 			ReportedProperties: &armazurestackhci.JobReportedProperties{
	// 				ValidationStatus: &armazurestackhci.EceActionStatus{
	// 					Status: to.Ptr("Success"),
	// 					Steps: []*armazurestackhci.DeploymentStep{
	// 						{
	// 							FullStepIndex: to.Ptr("0"),
	// 							Name: to.Ptr("Configure SDN Integration"),
	// 							Description: to.Ptr("Configure SDN Integration for the cluster."),
	// 							StartTimeUTC: to.Ptr("2023-06-09T00:08:19"),
	// 							EndTimeUTC: to.Ptr("2023-06-09T04:01:47"),
	// 							Status: to.Ptr("Success"),
	// 							Steps: []*armazurestackhci.DeploymentStep{
	// 								{
	// 									FullStepIndex: to.Ptr("0.1"),
	// 									Name: to.Ptr("Validate SDN Configuration"),
	// 									Description: to.Ptr("Validate SDN configuration parameters"),
	// 									StartTimeUTC: to.Ptr("2023-06-09T00:08:23"),
	// 									EndTimeUTC: to.Ptr("2023-06-09T01:10:10"),
	// 									Status: to.Ptr("Success"),
	// 									Steps: []*armazurestackhci.DeploymentStep{
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				DeploymentStatus: &armazurestackhci.EceActionStatus{
	// 					Status: to.Ptr("Success"),
	// 					Steps: []*armazurestackhci.DeploymentStep{
	// 						{
	// 							FullStepIndex: to.Ptr("0"),
	// 							Name: to.Ptr("Configure SDN Integration"),
	// 							Description: to.Ptr("Configure SDN Integration for the cluster."),
	// 							StartTimeUTC: to.Ptr("2023-06-09T00:08:19"),
	// 							EndTimeUTC: to.Ptr("2023-06-09T04:01:47"),
	// 							Status: to.Ptr("Success"),
	// 							Steps: []*armazurestackhci.DeploymentStep{
	// 								{
	// 									FullStepIndex: to.Ptr("0.1"),
	// 									Name: to.Ptr("Enable SDN Integration"),
	// 									Description: to.Ptr("Enable SDN Integration on cluster nodes"),
	// 									StartTimeUTC: to.Ptr("2023-06-09T00:08:23"),
	// 									EndTimeUTC: to.Ptr("2023-06-09T01:10:10"),
	// 									Status: to.Ptr("Success"),
	// 									Steps: []*armazurestackhci.DeploymentStep{
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 			Status: to.Ptr(armazurestackhci.JobStatusSucceeded),
	// 		},
	// 	},
	// }
}
