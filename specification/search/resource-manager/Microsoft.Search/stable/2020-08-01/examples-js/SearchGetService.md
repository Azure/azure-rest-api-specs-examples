Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-search_3.0.1/sdk/search/arm-search/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the search service with the given name in the given resource group.
 *
 * @summary Gets the search service with the given name in the given resource group.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchGetService.json
 */
async function searchGetService() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const searchServiceName = "mysearchservice";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.services.get(resourceGroupName, searchServiceName);
  console.log(result);
}

searchGetService().catch(console.error);
```
