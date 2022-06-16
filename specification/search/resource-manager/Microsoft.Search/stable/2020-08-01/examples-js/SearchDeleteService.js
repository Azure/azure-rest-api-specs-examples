const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a search service in the given resource group, along with its associated resources.
 *
 * @summary Deletes a search service in the given resource group, along with its associated resources.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchDeleteService.json
 */
async function searchDeleteService() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const searchServiceName = "mysearchservice";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.services.delete(resourceGroupName, searchServiceName);
  console.log(result);
}

searchDeleteService().catch(console.error);
