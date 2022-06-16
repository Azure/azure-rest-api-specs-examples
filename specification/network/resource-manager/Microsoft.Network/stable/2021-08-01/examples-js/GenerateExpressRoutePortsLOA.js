const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generate a letter of authorization for the requested ExpressRoutePort resource.
 *
 * @summary Generate a letter of authorization for the requested ExpressRoutePort resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/GenerateExpressRoutePortsLOA.json
 */
async function generateExpressRoutePortLoa() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const expressRoutePortName = "portName";
  const request = {
    customerName: "customerName",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRoutePorts.generateLOA(
    resourceGroupName,
    expressRoutePortName,
    request
  );
  console.log(result);
}

generateExpressRoutePortLoa().catch(console.error);
