const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes the specified peering from the ExpressRouteCrossConnection.
 *
 * @summary Deletes the specified peering from the ExpressRouteCrossConnection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/ExpressRouteCrossConnectionBgpPeeringDelete.json
 */
async function deleteExpressRouteCrossConnectionBgpPeering() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName =
    process.env["NETWORK_RESOURCE_GROUP"] || "CrossConnection-SiliconValley";
  const crossConnectionName = "<circuitServiceKey>";
  const peeringName = "AzurePrivatePeering";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCrossConnectionPeerings.beginDeleteAndWait(
    resourceGroupName,
    crossConnectionName,
    peeringName,
  );
  console.log(result);
}
