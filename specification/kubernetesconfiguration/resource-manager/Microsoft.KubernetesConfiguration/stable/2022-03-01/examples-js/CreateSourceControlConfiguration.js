const { SourceControlConfigurationClient } = require("@azure/arm-kubernetesconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new Kubernetes Source Control Configuration.
 *
 * @summary Create a new Kubernetes Source Control Configuration.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/CreateSourceControlConfiguration.json
 */
async function createSourceControlConfiguration() {
  const subscriptionId = "subId1";
  const resourceGroupName = "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const sourceControlConfigurationName = "SRS_GitHubConfig";
  const sourceControlConfiguration = {
    configurationProtectedSettings: {
      protectedSetting1Key: "protectedSetting1Value",
    },
    enableHelmOperator: true,
    helmOperatorProperties: {
      chartValues: "--set git.ssh.secretName=flux-git-deploy --set tillerNamespace=kube-system",
      chartVersion: "0.3.0",
    },
    operatorInstanceName: "SRSGitHubFluxOp-01",
    operatorNamespace: "SRS_Namespace",
    operatorParams: "--git-email=xyzgituser@users.srs.github.com",
    operatorScope: "namespace",
    operatorType: "Flux",
    repositoryUrl: "git@github.com:k8sdeveloper425/flux-get-started",
    sshKnownHostsContents:
      "c3NoLmRldi5henVyZS5jb20gc3NoLXJzYSBBQUFBQjNOemFDMXljMkVBQUFBREFRQUJBQUFCQVFDN0hyMW9UV3FOcU9sekdKT2ZHSjROYWtWeUl6ZjFyWFlkNGQ3d282akJsa0x2Q0E0b2RCbEwwbURVeVowL1FVZlRUcWV1K3RtMjJnT3N2K1ZyVlRNazZ2d1JVNzVnWS95OXV0NU1iM2JSNUJWNThkS1h5cTlBOVVlQjVDYWtlaG41WmdtNngxbUtvVnlmK0ZGbjI2aVlxWEpSZ3pJWlpjWjVWNmhyRTBRZzM5a1ptNGF6NDhvMEFVYmY2U3A0U0xkdm51TWEyc1ZOd0hCYm9TN0VKa201N1hRUFZVMy9RcHlOTEhiV0Rkend0cmxTK2V6MzBTM0FkWWhMS0VPeEFHOHdlT255cnRMSkFVZW45bVRrb2w4b0lJMWVkZjdtV1diV1ZmMG5CbWx5MjErblpjbUNUSVNRQnRkY3lQYUVubzdmRlFNREQyNi9zMGxmS29iNEt3OEg=",
  };
  const credential = new DefaultAzureCredential();
  const client = new SourceControlConfigurationClient(credential, subscriptionId);
  const result = await client.sourceControlConfigurations.createOrUpdate(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
    sourceControlConfigurationName,
    sourceControlConfiguration
  );
  console.log(result);
}

createSourceControlConfiguration().catch(console.error);
