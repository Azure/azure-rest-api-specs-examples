const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the bare metal machine key set of the provided cluster.
 *
 * @summary Delete the bare metal machine key set of the provided cluster.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BareMetalMachineKeySets_Delete.json
 */
async function deleteBareMetalMachineKeySetOfCluster() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const clusterName = "clusterName";
  const bareMetalMachineKeySetName = "bareMetalMachineKeySetName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.bareMetalMachineKeySets.beginDeleteAndWait(
    resourceGroupName,
    clusterName,
    bareMetalMachineKeySetName
  );
  console.log(result);
}
