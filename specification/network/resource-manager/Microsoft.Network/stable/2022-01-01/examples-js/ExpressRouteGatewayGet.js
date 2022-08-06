const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Fetches the details of a ExpressRoute gateway in a resource group.
 *
 * @summary Fetches the details of a ExpressRoute gateway in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ExpressRouteGatewayGet.json
 */
async function expressRouteGatewayGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "resourceGroupName";
  const expressRouteGatewayName = "expressRouteGatewayName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteGateways.get(resourceGroupName, expressRouteGatewayName);
  console.log(result);
}

expressRouteGatewayGet().catch(console.error);
