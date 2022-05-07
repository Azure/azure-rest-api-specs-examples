Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-search_3.0.1/sdk/search/arm-search/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a Private Endpoint connection to the search service in the given resource group.
 *
 * @summary Updates a Private Endpoint connection to the search service in the given resource group.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/UpdatePrivateEndpointConnection.json
 */
async function privateEndpointConnectionUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const searchServiceName = "mysearchservice";
  const privateEndpointConnectionName = "testEndpoint.50bf4fbe-d7c1-4b48-a642-4f5892642546";
  const privateEndpointConnection = {
    properties: {
      privateLinkServiceConnectionState: {
        description: "Rejected for some reason",
        status: "Rejected",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.update(
    resourceGroupName,
    searchServiceName,
    privateEndpointConnectionName,
    privateEndpointConnection
  );
  console.log(result);
}

privateEndpointConnectionUpdate().catch(console.error);
```
