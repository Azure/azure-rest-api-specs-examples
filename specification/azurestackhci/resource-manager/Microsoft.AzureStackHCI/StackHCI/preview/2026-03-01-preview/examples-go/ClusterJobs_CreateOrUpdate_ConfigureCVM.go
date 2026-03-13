package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/ClusterJobs_CreateOrUpdate_ConfigureCVM.json
func ExampleClusterJobsClient_BeginCreateOrUpdate_clusterJobsCreateOrUpdateConfigureCvmJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClusterJobsClient().BeginCreateOrUpdate(ctx, "test-rg", "myCluster", "ConfigureCVM", armazurestackhci.ClusterJob{
		Properties: &armazurestackhci.HciConfigureCvmJobProperties{
			JobType:              to.Ptr(armazurestackhci.HciJobTypeConfigureCVM),
			DeploymentMode:       to.Ptr(armazurestackhci.DeploymentModeDeploy),
			ConfidentialVMIntent: to.Ptr(armazurestackhci.ConfidentialVMIntentEnable),
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
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/user-pvtrp-rg/providers/Microsoft.AzureStackHCI/clusters/hciclu3/jobs/ConfigureCVM"),
	// 		Name: to.Ptr("ConfigureCVM"),
	// 		Type: to.Ptr("Microsoft.azurestackhci/clusters/jobs"),
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedBy: to.Ptr("user@microsoft.com"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T10:43:27.9471574Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T10:43:27.9471574Z"); return t}()),
	// 		},
	// 		Properties: &armazurestackhci.HciConfigureCvmJobProperties{
	// 			JobType: to.Ptr(armazurestackhci.HciJobTypeConfigureCVM),
	// 			ConfidentialVMIntent: to.Ptr(armazurestackhci.ConfidentialVMIntentEnable),
	// 			DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
	// 			ReportedProperties: &armazurestackhci.JobReportedProperties{
	// 				ValidationStatus: &armazurestackhci.EceActionStatus{
	// 					Status: to.Ptr("Error"),
	// 					Steps: []*armazurestackhci.DeploymentStep{
	// 						{
	// 							FullStepIndex: to.Ptr("0"),
	// 							Name: to.Ptr("Configure CVM Intent"),
	// 							Description: to.Ptr("Configure CVM Intent."),
	// 							StartTimeUTC: to.Ptr("2023-06-09T00:08:19"),
	// 							EndTimeUTC: to.Ptr("2023-06-09T04:01:47"),
	// 							Status: to.Ptr("Error"),
	// 							Exception: []*string{
	// 								to.Ptr("exception1"),
	// 								to.Ptr("exception2"),
	// 							},
	// 							Steps: []*armazurestackhci.DeploymentStep{
	// 								{
	// 									FullStepIndex: to.Ptr("0.1"),
	// 									Name: to.Ptr("Before Cloud Deployment"),
	// 									Description: to.Ptr("Before Cloud Deployment"),
	// 									StartTimeUTC: to.Ptr("2023-06-09T00:08:23"),
	// 									EndTimeUTC: to.Ptr("2023-06-09T01:10:10"),
	// 									Exception: []*string{
	// 										to.Ptr("exception1"),
	// 										to.Ptr("exception2"),
	// 									},
	// 									Steps: []*armazurestackhci.DeploymentStep{
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				DeploymentStatus: &armazurestackhci.EceActionStatus{
	// 					Status: to.Ptr("Error"),
	// 					Steps: []*armazurestackhci.DeploymentStep{
	// 						{
	// 							FullStepIndex: to.Ptr("0"),
	// 							Name: to.Ptr("Configure CVM Intent"),
	// 							Description: to.Ptr("Configure CVM Intent."),
	// 							StartTimeUTC: to.Ptr("2023-06-09T00:08:19"),
	// 							EndTimeUTC: to.Ptr("2023-06-09T04:01:47"),
	// 							Status: to.Ptr("Error"),
	// 							Exception: []*string{
	// 								to.Ptr("exception1"),
	// 								to.Ptr("exception2"),
	// 							},
	// 							Steps: []*armazurestackhci.DeploymentStep{
	// 								{
	// 									FullStepIndex: to.Ptr("0.1"),
	// 									Name: to.Ptr("Before Cloud Deployment"),
	// 									Description: to.Ptr("Before Cloud Deployment"),
	// 									StartTimeUTC: to.Ptr("2023-06-09T00:08:23"),
	// 									EndTimeUTC: to.Ptr("2023-06-09T01:10:10"),
	// 									Exception: []*string{
	// 										to.Ptr("exception1"),
	// 										to.Ptr("exception2"),
	// 									},
	// 									Steps: []*armazurestackhci.DeploymentStep{
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 			Status: to.Ptr(armazurestackhci.JobStatusNotSpecified),
	// 		},
	// 	},
	// }
}
