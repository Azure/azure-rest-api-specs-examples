const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the currently advertised routes table associated with the express route circuit in a resource group.
 *
 * @summary Gets the currently advertised routes table associated with the express route circuit in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/ExpressRouteCircuitRouteTableList.json
 */
async function listRouteTables() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const circuitName = "circuitName";
  const peeringName = "peeringName";
  const devicePath = "devicePath";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuits.beginListRoutesTableAndWait(
    resourceGroupName,
    circuitName,
    peeringName,
    devicePath
  );
  console.log(result);
}
