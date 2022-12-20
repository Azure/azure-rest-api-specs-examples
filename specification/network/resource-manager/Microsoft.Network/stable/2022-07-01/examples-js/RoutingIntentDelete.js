const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a RoutingIntent.
 *
 * @summary Deletes a RoutingIntent.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/RoutingIntentDelete.json
 */
async function routeTableDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualHubName = "virtualHub1";
  const routingIntentName = "Intent1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.routingIntentOperations.beginDeleteAndWait(
    resourceGroupName,
    virtualHubName,
    routingIntentName
  );
  console.log(result);
}

routeTableDelete().catch(console.error);
