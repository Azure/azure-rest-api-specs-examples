const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the specified ExpressRouteLink resource.
 *
 * @summary Retrieves the specified ExpressRouteLink resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/ExpressRouteLinkGet.json
 */
async function expressRouteLinkGet() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const expressRoutePortName = "portName";
  const linkName = "linkName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteLinks.get(
    resourceGroupName,
    expressRoutePortName,
    linkName,
  );
  console.log(result);
}
