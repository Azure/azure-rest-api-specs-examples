const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the inbound routes configured for the Virtual Hub on a particular connection.
 *
 * @summary Gets the inbound routes configured for the Virtual Hub on a particular connection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/GetInboundRoutes.json
 */
async function inboundRoutesForTheVirtualHubOnAParticularConnection() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const virtualHubName = "virtualHub1";
  const options = {
    body: {
      connectionType: "ExpressRouteConnection",
      resourceUri:
        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/exrGw1/expressRouteConnections/exrConn1",
    },
    queryParameters: { "api-version": "2022-05-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/inboundRoutes",
      subscriptionId,
      resourceGroupName,
      virtualHubName,
    )
    .post(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

inboundRoutesForTheVirtualHubOnAParticularConnection().catch(console.error);
