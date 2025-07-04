const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get bare metal machine key set of the provided cluster.
 *
 * @summary Get bare metal machine key set of the provided cluster.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/BareMetalMachineKeySets_Get.json
 */
async function getBareMetalMachineKeySetOfCluster() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const clusterName = "clusterName";
  const bareMetalMachineKeySetName = "bareMetalMachineKeySetName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.bareMetalMachineKeySets.get(
    resourceGroupName,
    clusterName,
    bareMetalMachineKeySetName,
  );
  console.log(result);
}
