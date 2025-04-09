const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Upgrades the Azure AI Search service to the latest version available.
 *
 * @summary Upgrades the Azure AI Search service to the latest version available.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/preview/2025-02-01-preview/examples/UpgradeSearchServiceToLatestVersion.json
 */
async function upgradeSearchServiceToLatestVersion() {
  const subscriptionId = process.env["SEARCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SEARCH_RESOURCE_GROUP"] || "rg1";
  const searchServiceName = "mysearchservice";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.service.beginUpgradeAndWait(resourceGroupName, searchServiceName);
  console.log(result);
}
