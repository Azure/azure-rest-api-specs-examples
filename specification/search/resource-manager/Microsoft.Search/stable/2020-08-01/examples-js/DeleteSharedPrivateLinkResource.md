```javascript
const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Initiates the deletion of the shared private link resource from the search service.
 *
 * @summary Initiates the deletion of the shared private link resource from the search service.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/DeleteSharedPrivateLinkResource.json
 */
async function sharedPrivateLinkResourceDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const searchServiceName = "mysearchservice";
  const sharedPrivateLinkResourceName = "testResource";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.sharedPrivateLinkResources.beginDeleteAndWait(
    resourceGroupName,
    searchServiceName,
    sharedPrivateLinkResourceName
  );
  console.log(result);
}

sharedPrivateLinkResourceDelete().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-search_3.0.1/sdk/search/arm-search/README.md) on how to add the SDK to your project and authenticate.
