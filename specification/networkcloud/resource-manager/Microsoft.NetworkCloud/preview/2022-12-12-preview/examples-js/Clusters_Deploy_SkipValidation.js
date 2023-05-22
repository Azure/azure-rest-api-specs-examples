const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deploy the cluster to the provided rack.
 *
 * @summary Deploy the cluster to the provided rack.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/Clusters_Deploy_SkipValidation.json
 */
async function deployClusterSkippingValidation() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const clusterName = "clusterName";
  const clusterDeployParameters = {
    skipValidationsForMachines: ["bmmName1"],
  };
  const options = { clusterDeployParameters };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.clusters.beginDeployAndWait(resourceGroupName, clusterName, options);
  console.log(result);
}
