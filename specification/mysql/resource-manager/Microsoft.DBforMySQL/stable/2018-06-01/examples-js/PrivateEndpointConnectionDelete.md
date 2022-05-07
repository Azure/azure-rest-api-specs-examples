Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-mysql_5.0.1/sdk/mysql/arm-mysql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a private endpoint connection with a given name.
 *
 * @summary Deletes a private endpoint connection with a given name.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/PrivateEndpointConnectionDelete.json
 */
async function deletesAPrivateEndpointConnectionWithAGivenName() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const serverName = "test-svr";
  const privateEndpointConnectionName = "private-endpoint-connection-name";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginDeleteAndWait(
    resourceGroupName,
    serverName,
    privateEndpointConnectionName
  );
  console.log(result);
}

deletesAPrivateEndpointConnectionWithAGivenName().catch(console.error);
```
