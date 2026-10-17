package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-10-01/EdgeMachineJobs_CreateOrUpdate_CollectLog.json
func ExampleEdgeMachineJobsClient_BeginCreateOrUpdate_edgeMachineJobsCreateOrUpdateCollectLog() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEdgeMachineJobsClient().BeginCreateOrUpdate(ctx, "ArcInstance-rg", "machine1", "triggerLogCollection", armazurestackhci.EdgeMachineJob{
		Properties: &armazurestackhci.EdgeMachineJobProperties{
			JobType:        to.Ptr(armazurestackhci.EdgeMachineJobType("EdgeMachineJobProperties")),
			DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeValidate),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.EdgeMachineJobsClientCreateOrUpdateResponse{
	// 	EdgeMachineJob: armazurestackhci.EdgeMachineJob{
	// 		Properties: &armazurestackhci.EdgeMachineCollectLogJobProperties{
	// 			JobType: to.Ptr(armazurestackhci.EdgeMachineJobTypeCollectLog),
	// 			FromDate: to.Ptr(time.Date(2024, time.January, 29, 10, 43, 27, 947157400, time.UTC)),
	// 			ToDate: to.Ptr(time.Date(2024, time.January, 29, 10, 43, 27, 947157400, time.UTC)),
	// 			DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeValidate),
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 			JobID: to.Ptr("egmijofqvbbkscxzfghmybzwdhfwny"),
	// 			StartTimeUTC: to.Ptr(time.Date(2024, time.December, 3, 11, 31, 42, 387000000, time.UTC)),
	// 			EndTimeUTC: to.Ptr(time.Date(2024, time.December, 3, 11, 31, 42, 387000000, time.UTC)),
	// 			Status: to.Ptr(armazurestackhci.JobStatusNotSpecified),
	// 		},
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.AzureStackHCI/edgeMachines/machine-1/jobs/collectLog"),
	// 		Name: to.Ptr("collectLog"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/edgeMachines/jobs"),
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2024, time.November, 14, 10, 46, 55, 167000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2024, time.November, 14, 10, 46, 55, 167000000, time.UTC)),
	// 		},
	// 	},
	// }
}
