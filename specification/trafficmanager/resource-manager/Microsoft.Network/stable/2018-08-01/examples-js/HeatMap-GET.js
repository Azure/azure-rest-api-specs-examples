const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

async function heatMapGet() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "azuresdkfornetautoresttrafficmanager1323";
  const profileName = "azuresdkfornetautoresttrafficmanager3880";
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.heatMap.get(resourceGroupName, profileName);
  console.log(result);
}

heatMapGet().catch(console.error);
