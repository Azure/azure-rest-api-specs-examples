const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all available location names for AFD log analytics report.
 *
 * @summary Get all available location names for AFD log analytics report.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/LogAnalytics_GetLogAnalyticsLocations.json
 */
async function logAnalyticsGetLogAnalyticsLocations() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.logAnalytics.getLogAnalyticsLocations(resourceGroupName, profileName);
  console.log(result);
}
