Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-search_3.0.1/sdk/search/arm-search/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generates a new query key for the specified search service. You can create up to 50 query keys per service.
 *
 * @summary Generates a new query key for the specified search service. You can create up to 50 query keys per service.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchCreateQueryKey.json
 */
async function searchCreateQueryKey() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const searchServiceName = "mysearchservice";
  const name = "Query key for browser-based clients";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.queryKeys.create(resourceGroupName, searchServiceName, name);
  console.log(result);
}

searchCreateQueryKey().catch(console.error);
```
