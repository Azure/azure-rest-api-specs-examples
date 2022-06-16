package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/CreateSourceControlConfiguration.json
func ExampleSourceControlConfigurationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkubernetesconfiguration.NewSourceControlConfigurationsClient("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg1",
		"Microsoft.Kubernetes",
		"connectedClusters",
		"clusterName1",
		"SRS_GitHubConfig",
		armkubernetesconfiguration.SourceControlConfiguration{
			Properties: &armkubernetesconfiguration.SourceControlConfigurationProperties{
				ConfigurationProtectedSettings: map[string]*string{
					"protectedSetting1Key": to.Ptr("protectedSetting1Value"),
				},
				EnableHelmOperator: to.Ptr(true),
				HelmOperatorProperties: &armkubernetesconfiguration.HelmOperatorProperties{
					ChartValues:  to.Ptr("--set git.ssh.secretName=flux-git-deploy --set tillerNamespace=kube-system"),
					ChartVersion: to.Ptr("0.3.0"),
				},
				OperatorInstanceName:  to.Ptr("SRSGitHubFluxOp-01"),
				OperatorNamespace:     to.Ptr("SRS_Namespace"),
				OperatorParams:        to.Ptr("--git-email=xyzgituser@users.srs.github.com"),
				OperatorScope:         to.Ptr(armkubernetesconfiguration.OperatorScopeTypeNamespace),
				OperatorType:          to.Ptr(armkubernetesconfiguration.OperatorTypeFlux),
				RepositoryURL:         to.Ptr("git@github.com:k8sdeveloper425/flux-get-started"),
				SSHKnownHostsContents: to.Ptr("c3NoLmRldi5henVyZS5jb20gc3NoLXJzYSBBQUFBQjNOemFDMXljMkVBQUFBREFRQUJBQUFCQVFDN0hyMW9UV3FOcU9sekdKT2ZHSjROYWtWeUl6ZjFyWFlkNGQ3d282akJsa0x2Q0E0b2RCbEwwbURVeVowL1FVZlRUcWV1K3RtMjJnT3N2K1ZyVlRNazZ2d1JVNzVnWS95OXV0NU1iM2JSNUJWNThkS1h5cTlBOVVlQjVDYWtlaG41WmdtNngxbUtvVnlmK0ZGbjI2aVlxWEpSZ3pJWlpjWjVWNmhyRTBRZzM5a1ptNGF6NDhvMEFVYmY2U3A0U0xkdm51TWEyc1ZOd0hCYm9TN0VKa201N1hRUFZVMy9RcHlOTEhiV0Rkend0cmxTK2V6MzBTM0FkWWhMS0VPeEFHOHdlT255cnRMSkFVZW45bVRrb2w4b0lJMWVkZjdtV1diV1ZmMG5CbWx5MjErblpjbUNUSVNRQnRkY3lQYUVubzdmRlFNREQyNi9zMGxmS29iNEt3OEg="),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
