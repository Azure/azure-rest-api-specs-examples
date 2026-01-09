const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list hosts by zone in a cluster
 *
 * @summary list hosts by zone in a cluster
 * x-ms-original-file: 2025-09-01/Clusters_ListZones.json
 */
async function clustersListZones() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.clusters.listZones("group1", "cloud1", "cluster1");
  console.log(result);
}
