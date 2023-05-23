const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the metrics configuration of the provided cluster.
 *
 * @summary Delete the metrics configuration of the provided cluster.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/ClusterMetricsConfigurations_Delete.json
 */
async function deleteMetricsConfigurationOfCluster() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const clusterName = "clusterName";
  const metricsConfigurationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.metricsConfigurations.beginDeleteAndWait(
    resourceGroupName,
    clusterName,
    metricsConfigurationName
  );
  console.log(result);
}
