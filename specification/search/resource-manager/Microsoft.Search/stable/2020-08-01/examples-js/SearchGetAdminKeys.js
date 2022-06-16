const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the primary and secondary admin API keys for the specified Azure Cognitive Search service.
 *
 * @summary Gets the primary and secondary admin API keys for the specified Azure Cognitive Search service.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchGetAdminKeys.json
 */
async function searchGetAdminKeys() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const searchServiceName = "mysearchservice";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.adminKeys.get(resourceGroupName, searchServiceName);
  console.log(result);
}

searchGetAdminKeys().catch(console.error);
