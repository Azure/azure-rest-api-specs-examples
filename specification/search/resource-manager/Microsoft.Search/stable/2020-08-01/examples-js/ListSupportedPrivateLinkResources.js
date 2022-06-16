const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all supported private link resource types for the given service.
 *
 * @summary Gets a list of all supported private link resource types for the given service.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/ListSupportedPrivateLinkResources.json
 */
async function listSupportedPrivateLinkResources() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const searchServiceName = "mysearchservice";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkResources.listSupported(
    resourceGroupName,
    searchServiceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSupportedPrivateLinkResources().catch(console.error);
