const { SourceControlConfigurationClient } = require("@azure/arm-kubernetesconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets details of the Source Control Configuration.
 *
 * @summary Gets details of the Source Control Configuration.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/GetSourceControlConfiguration.json
 */
async function getSourceControlConfiguration() {
  const subscriptionId = process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] || "subId1";
  const resourceGroupName = process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const sourceControlConfigurationName = "SRS_GitHubConfig";
  const credential = new DefaultAzureCredential();
  const client = new SourceControlConfigurationClient(credential, subscriptionId);
  const result = await client.sourceControlConfigurations.get(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
    sourceControlConfigurationName
  );
  console.log(result);
}
