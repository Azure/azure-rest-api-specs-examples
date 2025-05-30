const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the outbound routes configured for the Virtual Hub on a particular connection.
 *
 * @summary Gets the outbound routes configured for the Virtual Hub on a particular connection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/GetOutboundRoutes.json
 */
async function outboundRoutesForTheVirtualHubOnAParticularConnection() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualHubName = "virtualHub1";
  const getOutboundRoutesParameters = {
    connectionType: "ExpressRouteConnection",
    resourceUri:
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/exrGw1/expressRouteConnections/exrConn1",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualHubs.beginGetOutboundRoutesAndWait(
    resourceGroupName,
    virtualHubName,
    getOutboundRoutesParameters,
  );
  console.log(result);
}
