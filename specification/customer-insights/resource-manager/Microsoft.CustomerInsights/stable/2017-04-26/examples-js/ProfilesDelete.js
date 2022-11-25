const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a profile within a hub
 *
 * @summary Deletes a profile within a hub
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ProfilesDelete.json
 */
async function profilesDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const profileName = "TestProfileType396";
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.profiles.beginDeleteAndWait(resourceGroupName, hubName, profileName);
  console.log(result);
}

profilesDelete().catch(console.error);
