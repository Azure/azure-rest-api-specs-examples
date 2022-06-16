const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified Express Route Circuit Connection from the specified express route circuit.
 *
 * @summary Deletes the specified Express Route Circuit Connection from the specified express route circuit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRouteCircuitConnectionDelete.json
 */
async function deleteExpressRouteCircuit() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const circuitName = "ExpressRouteARMCircuitA";
  const peeringName = "AzurePrivatePeering";
  const connectionName = "circuitConnectionUSAUS";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuitConnections.beginDeleteAndWait(
    resourceGroupName,
    circuitName,
    peeringName,
    connectionName
  );
  console.log(result);
}

deleteExpressRouteCircuit().catch(console.error);
