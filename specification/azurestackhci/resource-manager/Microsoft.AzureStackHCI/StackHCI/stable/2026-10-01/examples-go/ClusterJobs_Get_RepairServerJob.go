package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-10-01/ClusterJobs_Get_RepairServerJob.json
func ExampleClusterJobsClient_Get_clusterJobsGetRepairServerJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClusterJobsClient().Get(ctx, "test-rg", "myCluster", "repairserver", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.ClusterJobsClientGetResponse{
	// 	ClusterJob: armazurestackhci.ClusterJob{
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/myCluster/jobs/repairserver"),
	// 		Name: to.Ptr("repairserver"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/clusters/jobs"),
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedBy: to.Ptr("user@microsoft.com"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2024, time.January, 29, 10, 43, 27, 947157400, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2024, time.January, 29, 10, 43, 27, 947157400, time.UTC)),
	// 		},
	// 		Properties: &armazurestackhci.HciRepairServerJobProperties{
	// 			JobType: to.Ptr(armazurestackhci.HciJobTypeRepairServer),
	// 			DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
	// 			RepairServerJobServerDetails: []*armazurestackhci.RepairServerJobServerDetails{
	// 				{
	// 					ServerName: to.Ptr("Node-2"),
	// 					ServerResourceID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-2"),
	// 				},
	// 			},
	// 			Secrets: []*armazurestackhci.EceDeploymentSecrets{
	// 				{
	// 					SecretName: to.Ptr("cluster1-LocalAdminCredential-f5bcc1d9-23af-4ae9-aca1-041d0f593a63"),
	// 					EceSecretName: to.Ptr(armazurestackhci.EceSecretsLocalAdminCredential),
	// 					SecretLocation: to.Ptr("https://sclusterkvnirhci35.vault.azure.net/secrets/cluster-34232342-LocalAdminCredential-f5bcc1d9-23af-4ae9-aca1-041d0f593a63/9276354aabfc492fa9b2cdbefb54ae4b"),
	// 				},
	// 				{
	// 					SecretName: to.Ptr("cluster2-AzureStackLCMUserCredential-f5bcc1d9-23af-4ae9-aca1-041d0f593a63"),
	// 					EceSecretName: to.Ptr(armazurestackhci.EceSecretsAzureStackLCMUserCredential),
	// 					SecretLocation: to.Ptr("https://sclusterkvnirhci35.vault.azure.net/secrets/cluster-34232342-AzureStackLCMUserCredential-f5bcc1d9-23af-4ae9-aca1-041d0f593a63/9276354aabfc492fa9b2cdbefb54ae4c"),
	// 				},
	// 			},
	// 			ReportedProperties: &armazurestackhci.JobReportedProperties{
	// 				ValidationStatus: &armazurestackhci.EceActionStatus{
	// 					Status: to.Ptr("Error"),
	// 					Steps: []*armazurestackhci.DeploymentStep{
	// 						{
	// 							FullStepIndex: to.Ptr("0"),
	// 							Name: to.Ptr("Repair Server Deployment"),
	// 							Description: to.Ptr("Repair Server."),
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
	// 							Name: to.Ptr("Repair Server Deployment"),
	// 							Description: to.Ptr("Repair Server."),
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
