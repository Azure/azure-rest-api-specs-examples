const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Delete the metrics configuration of the provided cluster.
 *
 * @summary Delete the metrics configuration of the provided cluster.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-10-01-preview/examples/ClusterMetricsConfigurations_Delete.json
 */
async function deleteMetricsConfigurationOfCluster() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const clusterName = "clusterName";
  const metricsConfigurationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.metricsConfigurations.beginDeleteAndWait(
    resourceGroupName,
    clusterName,
    metricsConfigurationName,
  );
  console.log(result);
}
