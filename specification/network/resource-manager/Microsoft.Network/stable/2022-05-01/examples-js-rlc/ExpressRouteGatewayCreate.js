const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a ExpressRoute gateway in a specified resource group.
 *
 * @summary Creates or updates a ExpressRoute gateway in a specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ExpressRouteGatewayCreate.json
 */
async function expressRouteGatewayCreate() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "resourceGroupName";
  const expressRouteGatewayName = "gateway-2";
  const options = {
    body: {
      location: "westus",
      properties: {
        autoScaleConfiguration: { bounds: { min: 3 } },
        virtualHub: {
          id: "/subscriptions/subid/resourceGroups/resourceGroupId/providers/Microsoft.Network/virtualHubs/virtualHubName",
        },
      },
    },
    queryParameters: { "api-version": "2022-05-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}",
      subscriptionId,
      resourceGroupName,
      expressRouteGatewayName,
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

expressRouteGatewayCreate().catch(console.error);
