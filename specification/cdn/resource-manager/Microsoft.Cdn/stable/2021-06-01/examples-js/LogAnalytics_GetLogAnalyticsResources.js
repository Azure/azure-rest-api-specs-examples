const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all endpoints and custom domains available for AFD log report
 *
 * @summary Get all endpoints and custom domains available for AFD log report
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetLogAnalyticsResources.json
 */
async function logAnalyticsGetLogAnalyticsResources() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.logAnalytics.getLogAnalyticsResources(resourceGroupName, profileName);
  console.log(result);
}

logAnalyticsGetLogAnalyticsResources().catch(console.error);
