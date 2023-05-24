package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/GetSourceControlConfiguration.json
func ExampleSourceControlConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkubernetesconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSourceControlConfigurationsClient().Get(ctx, "rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", "SRS_GitHubConfig", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SourceControlConfiguration = armkubernetesconfiguration.SourceControlConfiguration{
	// 	Name: to.Ptr("SRS_GitHubConfig"),
	// 	Type: to.Ptr("Microsoft.KubernetesConfiguration/sourceControlConfigurations"),
	// 	ID: to.Ptr("/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/clusterName1/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/SRS_GitHubConfig"),
	// 	Properties: &armkubernetesconfiguration.SourceControlConfigurationProperties{
	// 		ComplianceStatus: &armkubernetesconfiguration.ComplianceStatus{
	// 			ComplianceState: to.Ptr(armkubernetesconfiguration.ComplianceStateTypePending),
	// 			LastConfigApplied: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-08T05:25:32.122Z"); return t}()),
	// 			Message: to.Ptr("Configuration successfully created"),
	// 			MessageLevel: to.Ptr(armkubernetesconfiguration.MessageLevelType("Info")),
	// 		},
	// 		HelmOperatorProperties: &armkubernetesconfiguration.HelmOperatorProperties{
	// 			ChartValues: to.Ptr("--set git.ssh.secretName=flux-git-deploy --set tillerNamespace=kube-system"),
	// 			ChartVersion: to.Ptr("0.3.0"),
	// 		},
	// 		OperatorInstanceName: to.Ptr("SRSGitHubFluxOp-01"),
	// 		OperatorNamespace: to.Ptr("SRS_Namespace"),
	// 		OperatorParams: to.Ptr("--git-email=xyzgituser@users.srs.github.com"),
	// 		OperatorScope: to.Ptr(armkubernetesconfiguration.OperatorScopeTypeNamespace),
	// 		OperatorType: to.Ptr(armkubernetesconfiguration.OperatorTypeFlux),
	// 		ProvisioningState: to.Ptr(armkubernetesconfiguration.ProvisioningStateTypeSucceeded),
	// 		RepositoryPublicKey: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAQEAqvTzeL+BWgcHnC1CTBRMg2ZfCh9khlrvb2avFHiGG24rRvjLHlKqtfiiw+cZNCKskUyVKqamD2RHrhyn/wXvJ9fFRt0LhYLKn4hJhJaPx4IawdWnW1MUv4U+Mr8o3Cxps4EmiZemqri3fOrhzEIlPL272whKpzlDLV7L4W1XQIGmVPwQ93HTzKEd5uHuEuw6JyFftDDLlCnd3Q1kQ7HOabFEfcSSr9cMx2MU4j/Pjuf3Rd/CzeksvKtU009KSXSnWKm8LL1fihSc1H1WDTi8iuZtT63hsNYH1yxrPRpMVScs3ufLViAGO9NEHQSgDdl/OERQQqKisUn2Qm6adgmftw== rsa-key-20190909"),
	// 		RepositoryURL: to.Ptr("git@github.com:k8sdeveloper425/flux-get-started"),
	// 		SSHKnownHostsContents: to.Ptr("c3NoLmRldi5henVyZS5jb20gc3NoLXJzYSBBQUFBQjNOemFDMXljMkVBQUFBREFRQUJBQUFCQVFDN0hyMW9UV3FOcU9sekdKT2ZHSjROYWtWeUl6ZjFyWFlkNGQ3d282akJsa0x2Q0E0b2RCbEwwbURVeVowL1FVZlRUcWV1K3RtMjJnT3N2K1ZyVlRNazZ2d1JVNzVnWS95OXV0NU1iM2JSNUJWNThkS1h5cTlBOVVlQjVDYWtlaG41WmdtNngxbUtvVnlmK0ZGbjI2aVlxWEpSZ3pJWlpjWjVWNmhyRTBRZzM5a1ptNGF6NDhvMEFVYmY2U3A0U0xkdm51TWEyc1ZOd0hCYm9TN0VKa201N1hRUFZVMy9RcHlOTEhiV0Rkend0cmxTK2V6MzBTM0FkWWhMS0VPeEFHOHdlT255cnRMSkFVZW45bVRrb2w4b0lJMWVkZjdtV1diV1ZmMG5CbWx5MjErblpjbUNUSVNRQnRkY3lQYUVubzdmRlFNREQyNi9zMGxmS29iNEt3OEg="),
	// 	},
	// 	SystemData: &armkubernetesconfiguration.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-08T05:10:57.027Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-08T05:10:57.027Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
	// 	},
	// }
}
