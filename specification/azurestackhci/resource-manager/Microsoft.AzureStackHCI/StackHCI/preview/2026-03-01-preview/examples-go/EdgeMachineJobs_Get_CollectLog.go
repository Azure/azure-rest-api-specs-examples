package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/EdgeMachineJobs_Get_CollectLog.json
func ExampleEdgeMachineJobsClient_Get_edgeMachineJobsGetCollectLog() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEdgeMachineJobsClient().Get(ctx, "ArcInstance-rg", "machine1", "collectLog", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.EdgeMachineJobsClientGetResponse{
	// 	EdgeMachineJob: &armazurestackhci.EdgeMachineJob{
	// 		Properties: &armazurestackhci.EdgeMachineCollectLogJobProperties{
	// 			JobType: to.Ptr(armazurestackhci.EdgeMachineJobTypeCollectLog),
	// 			FromDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T10:43:27.9471574Z"); return t}()),
	// 			ToDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T10:43:27.9471574Z"); return t}()),
	// 			DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeValidate),
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 			JobID: to.Ptr("egmijofqvbbkscxzfghmybzwdhfwny"),
	// 			StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-03T11:31:42.387Z"); return t}()),
	// 			EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-03T11:31:42.387Z"); return t}()),
	// 			Status: to.Ptr(armazurestackhci.JobStatusNotSpecified),
	// 		},
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.AzureStackHCI/edgeMachines/machine-1/jobs/collectLog"),
	// 		Name: to.Ptr("collectLog"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/edgeMachines/jobs"),
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-14T10:46:55.167Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-14T10:46:55.167Z"); return t}()),
	// 		},
	// 	},
	// }
}
