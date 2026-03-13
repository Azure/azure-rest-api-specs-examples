package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/ClusterJobs_List.json
func ExampleClusterJobsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("6D37FF61-4C93-4377-B06B-FC6D6D561A7D", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClusterJobsClient().NewListPager("rghci", "Ql40O4-I77S", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armazurestackhci.ClusterJobsClientListResponse{
		// 	ClusterJobListResult: armazurestackhci.ClusterJobListResult{
		// 		Value: []*armazurestackhci.ClusterJob{
		// 			{
		// 				Properties: &armazurestackhci.HciConfigureSdnIntegrationJobProperties{
		// 					JobType: to.Ptr(armazurestackhci.HciJobTypeConfigureSdnIntegration),
		// 					DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
		// 					SdnIntegrationIntent: to.Ptr(armazurestackhci.SdnIntegrationIntentEnable),
		// 				},
		// 				ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/clu1/jobs/configureSdnIntegration"),
		// 				Name: to.Ptr("configureSdnIntegration"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/clusters/jobs"),
		// 				SystemData: &armazurestackhci.SystemData{
		// 					CreatedBy: to.Ptr("zsnvvvbj"),
		// 					CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T17:02:21.168Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("qxlrx"),
		// 					LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T17:02:21.168Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Properties: &armazurestackhci.HciConfigureCvmJobProperties{
		// 					JobType: to.Ptr(armazurestackhci.HciJobTypeConfigureCVM),
		// 					DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
		// 					ConfidentialVMIntent: to.Ptr(armazurestackhci.ConfidentialVMIntentEnable),
		// 				},
		// 				ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/clu1/jobs/configureCvm"),
		// 				Name: to.Ptr("configureCvm"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/clusters/jobs"),
		// 				SystemData: &armazurestackhci.SystemData{
		// 					CreatedBy: to.Ptr("zsnvvvbj"),
		// 					CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T17:02:21.168Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("qxlrx"),
		// 					LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T17:02:21.168Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/axw"),
		// 	},
		// }
	}
}
