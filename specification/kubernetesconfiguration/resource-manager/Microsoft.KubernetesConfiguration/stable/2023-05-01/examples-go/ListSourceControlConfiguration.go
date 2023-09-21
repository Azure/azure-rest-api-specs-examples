package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/examples/ListSourceControlConfiguration.json
func ExampleSourceControlConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkubernetesconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSourceControlConfigurationsClient().NewListPager("rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", nil)
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
		// page.SourceControlConfigurationList = armkubernetesconfiguration.SourceControlConfigurationList{
		// 	Value: []*armkubernetesconfiguration.SourceControlConfiguration{
		// 		{
		// 			Name: to.Ptr("SRS_GitHubConfig"),
		// 			Type: to.Ptr("Microsoft.KubernetesConfiguration/sourceControlConfigurations"),
		// 			ID: to.Ptr("/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/clusterName1/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/SRS_GitHubConfig"),
		// 			Properties: &armkubernetesconfiguration.SourceControlConfigurationProperties{
		// 				ComplianceStatus: &armkubernetesconfiguration.ComplianceStatus{
		// 					ComplianceState: to.Ptr(armkubernetesconfiguration.ComplianceStateTypeCompliant),
		// 				},
		// 				OperatorInstanceName: to.Ptr("SRSGitHubFluxOp-01"),
		// 				OperatorNamespace: to.Ptr("SRS_Namespace"),
		// 				OperatorScope: to.Ptr(armkubernetesconfiguration.OperatorScopeTypeNamespace),
		// 				OperatorType: to.Ptr(armkubernetesconfiguration.OperatorTypeFlux),
		// 				ProvisioningState: to.Ptr(armkubernetesconfiguration.ProvisioningStateTypeSucceeded),
		// 				RepositoryURL: to.Ptr("git@github.com:k8sdeveloper425/SRSClusterconfigs"),
		// 				SSHKnownHostsContents: to.Ptr("c3NoLmRldi5henVyZS5jb20gc3NoLXJzYSBBQUFBQjNOemFDMXljMkVBQUFBREFRQUJBQUFCQVFDN0hyMW9UV3FOcU9sekdKT2ZHSjROYWtWeUl6ZjFyWFlkNGQ3d282akJsa0x2Q0E0b2RCbEwwbURVeVowL1FVZlRUcWV1K3RtMjJnT3N2K1ZyVlRNazZ2d1JVNzVnWS95OXV0NU1iM2JSNUJWNThkS1h5cTlBOVVlQjVDYWtlaG41WmdtNngxbUtvVnlmK0ZGbjI2aVlxWEpSZ3pJWlpjWjVWNmhyRTBRZzM5a1ptNGF6NDhvMEFVYmY2U3A0U0xkdm51TWEyc1ZOd0hCYm9TN0VKa201N1hRUFZVMy9RcHlOTEhiV0Rkend0cmxTK2V6MzBTM0FkWWhMS0VPeEFHOHdlT255cnRMSkFVZW45bVRrb2w4b0lJMWVkZjdtV1diV1ZmMG5CbWx5MjErblpjbUNUSVNRQnRkY3lQYUVubzdmRlFNREQyNi9zMGxmS29iNEt3OEg="),
		// 			},
		// 			SystemData: &armkubernetesconfiguration.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-08T05:10:57.027Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-08T05:10:57.027Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SCRS_GitHubConfig"),
		// 			Type: to.Ptr("Microsoft.KubernetesConfiguration/sourceControlConfigurations"),
		// 			ID: to.Ptr("/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/clusterName1/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/SCRS_GitHubConfig"),
		// 			Properties: &armkubernetesconfiguration.SourceControlConfigurationProperties{
		// 				ComplianceStatus: &armkubernetesconfiguration.ComplianceStatus{
		// 					ComplianceState: to.Ptr(armkubernetesconfiguration.ComplianceStateTypeCompliant),
		// 				},
		// 				OperatorInstanceName: to.Ptr("SCRSGitHubFluxOp-02"),
		// 				OperatorNamespace: to.Ptr("SCRS_Namespace"),
		// 				OperatorScope: to.Ptr(armkubernetesconfiguration.OperatorScopeTypeCluster),
		// 				OperatorType: to.Ptr(armkubernetesconfiguration.OperatorTypeFlux),
		// 				ProvisioningState: to.Ptr(armkubernetesconfiguration.ProvisioningStateTypeSucceeded),
		// 				RepositoryURL: to.Ptr("git@github.com:k8sdeveloper425/SCRSClusterconfigs"),
		// 				SSHKnownHostsContents: to.Ptr("c3NoLmRldi5henVyZS5jb20gc3NoLXJzYSBBQUFBQjNOemFDMXljMkVBQUFBREFRQUJBQUFCQVFDN0hyMW9UV3FOcU9sekdKT2ZHSjROYWtWeUl6ZjFyWFlkNGQ3d282akJsa0x2Q0E0b2RCbEwwbURVeVowL1FVZlRUcWV1K3RtMjJnT3N2K1ZyVlRNazZ2d1JVNzVnWS95OXV0NU1iM2JSNUJWNThkS1h5cTlBOVVlQjVDYWtlaG41WmdtNngxbUtvVnlmK0ZGbjI2aVlxWEpSZ3pJWlpjWjVWNmhyRTBRZzM5a1ptNGF6NDhvMEFVYmY2U3A0U0xkdm51TWEyc1ZOd0hCYm9TN0VKa201N1hRUFZVMy9RcHlOTEhiV0Rkend0cmxTK2V6MzBTM0FkWWhMS0VPeEFHOHdlT255cnRMSkFVZW45bVRrb2w4b0lJMWVkZjdtV1diV1ZmMG5CbWx5MjErblpjbUNUSVNRQnRkY3lQYUVubzdmRlFNREQyNi9zMGxmS29iNEt3OEg="),
		// 			},
		// 			SystemData: &armkubernetesconfiguration.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-08T05:10:57.027Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-08T05:10:57.027Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
		// 			},
		// 	}},
		// }
	}
}
