const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets latest heatmap for Traffic Manager profile.
 *
 * @summary gets latest heatmap for Traffic Manager profile.
 * x-ms-original-file: 2024-04-01-preview/HeatMap-GET-With-TopLeft-BotRight.json
 */
async function heatMapGETWithTopLeftBotRight() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.heatMap.get(
    "azuresdkfornetautoresttrafficmanager1323",
    "azuresdkfornetautoresttrafficmanager3880",
    { topLeft: [10, 50.001], botRight: [-50.001, 80] },
  );
  console.log(result);
}
