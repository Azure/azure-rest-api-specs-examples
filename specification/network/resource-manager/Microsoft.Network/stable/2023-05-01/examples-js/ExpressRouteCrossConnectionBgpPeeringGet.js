const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified peering for the ExpressRouteCrossConnection.
 *
 * @summary Gets the specified peering for the ExpressRouteCrossConnection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/ExpressRouteCrossConnectionBgpPeeringGet.json
 */
async function getExpressRouteCrossConnectionBgpPeering() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName =
    process.env["NETWORK_RESOURCE_GROUP"] || "CrossConnection-SiliconValley";
  const crossConnectionName = "<circuitServiceKey>";
  const peeringName = "AzurePrivatePeering";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCrossConnectionPeerings.get(
    resourceGroupName,
    crossConnectionName,
    peeringName
  );
  console.log(result);
}
