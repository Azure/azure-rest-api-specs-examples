const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets latest heatmap for Traffic Manager profile.
 *
 * @summary Gets latest heatmap for Traffic Manager profile.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/HeatMap-GET-With-TopLeft-BotRight.json
 */
async function heatMapGetWithTopLeftBotRight() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "azuresdkfornetautoresttrafficmanager1323";
  const profileName = "azuresdkfornetautoresttrafficmanager3880";
  const topLeft = [10, 50.001];
  const botRight = [-50.001, 80];
  const options = { topLeft, botRight };
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.heatMap.get(resourceGroupName, profileName, options);
  console.log(result);
}

heatMapGetWithTopLeftBotRight().catch(console.error);
