Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-postgresql_6.0.1/sdk/postgresql/arm-postgresql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PostgreSQLManagementClient } = require("@azure/arm-postgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates private endpoint connection with the specified tags.
 *
 * @summary Updates private endpoint connection with the specified tags.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/PrivateEndpointConnectionUpdateTags.json
 */
async function updatePrivateEndpointConnectionTags() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const serverName = "test-svr";
  const privateEndpointConnectionName = "private-endpoint-connection-name";
  const parameters = {
    tags: { key1: "val1", key2: "val2", key3: "val3" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginUpdateTagsAndWait(
    resourceGroupName,
    serverName,
    privateEndpointConnectionName,
    parameters
  );
  console.log(result);
}

updatePrivateEndpointConnectionTags().catch(console.error);
```
