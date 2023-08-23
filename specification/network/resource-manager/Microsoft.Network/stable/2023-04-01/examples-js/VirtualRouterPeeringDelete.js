const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified peering from a Virtual Router.
 *
 * @summary Deletes the specified peering from a Virtual Router.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualRouterPeeringDelete.json
 */
async function deleteVirtualRouterPeering() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualRouterName = "virtualRouter";
  const peeringName = "peering1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualRouterPeerings.beginDeleteAndWait(
    resourceGroupName,
    virtualRouterName,
    peeringName
  );
  console.log(result);
}
