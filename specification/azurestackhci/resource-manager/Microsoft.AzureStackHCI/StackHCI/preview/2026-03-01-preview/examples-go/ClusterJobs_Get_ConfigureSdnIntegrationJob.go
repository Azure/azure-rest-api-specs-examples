package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/ClusterJobs_Get_ConfigureSdnIntegrationJob.json
func ExampleClusterJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("6D37FF61-4C93-4377-B06B-FC6D6D561A7D", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClusterJobsClient().Get(ctx, "rghci", "Y-k0MG", "configureSdnIntegration", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.ClusterJobsClientGetResponse{
	// 	ClusterJob: &armazurestackhci.ClusterJob{
	// 		Properties: &armazurestackhci.HciConfigureSdnIntegrationJobProperties{
	// 			JobType: to.Ptr(armazurestackhci.HciJobTypeConfigureSdnIntegration),
	// 			DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
	// 			SdnIntegrationIntent: to.Ptr(armazurestackhci.SdnIntegrationIntentEnable),
	// 			ReportedProperties: &armazurestackhci.JobReportedProperties{
	// 				ValidationStatus: &armazurestackhci.EceActionStatus{
	// 					Status: to.Ptr("Error"),
	// 					Steps: []*armazurestackhci.DeploymentStep{
	// 						{
	// 							FullStepIndex: to.Ptr("0"),
	// 							Name: to.Ptr("Configure SDN Integration Deployment"),
	// 							Description: to.Ptr("Configure SDN Integration."),
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
	// 							Name: to.Ptr("Configure SDN Integration Deployment"),
	// 							Description: to.Ptr("Configure SDN Integration."),
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
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/clu1/jobs/configureSdnIntegration"),
	// 		Name: to.Ptr("configureSdnIntegration"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/clusters/jobs"),
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedBy: to.Ptr("zsnvvvbj"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T17:02:21.168Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("qxlrx"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T17:02:21.168Z"); return t}()),
	// 		},
	// 	},
	// }
}
