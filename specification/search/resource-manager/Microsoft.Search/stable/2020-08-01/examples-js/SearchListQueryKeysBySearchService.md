Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-search_3.0.1/sdk/search/arm-search/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of query API keys for the given Azure Cognitive Search service.
 *
 * @summary Returns the list of query API keys for the given Azure Cognitive Search service.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchListQueryKeysBySearchService.json
 */
async function searchListQueryKeysBySearchService() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const searchServiceName = "mysearchservice";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.queryKeys.listBySearchService(
    resourceGroupName,
    searchServiceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

searchListQueryKeysBySearchService().catch(console.error);
```
