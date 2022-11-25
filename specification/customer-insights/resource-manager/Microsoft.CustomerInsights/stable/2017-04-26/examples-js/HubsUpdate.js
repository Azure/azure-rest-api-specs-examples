const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a Hub.
 *
 * @summary Updates a Hub.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/HubsUpdate.json
 */
async function hubsUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const parameters = {
    hubBillingInfo: { maxUnits: 5, minUnits: 1, skuName: "B0" },
    location: "West US",
  };
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.hubs.update(resourceGroupName, hubName, parameters);
  console.log(result);
}

hubsUpdate().catch(console.error);
