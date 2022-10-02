const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to NOTE: This feature is currently in preview and still being tested for stability. Lists all available internet service providers for a specified Azure region.
 *
 * @summary NOTE: This feature is currently in preview and still being tested for stability. Lists all available internet service providers for a specified Azure region.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkWatcherAvailableProvidersListGet.json
 */
async function getAvailableProvidersList() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const parameters = {
    azureLocations: ["West US"],
    city: "seattle",
    country: "United States",
    state: "washington",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkWatchers.beginListAvailableProvidersAndWait(
    resourceGroupName,
    networkWatcherName,
    parameters
  );
  console.log(result);
}

getAvailableProvidersList().catch(console.error);
