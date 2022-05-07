Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the private endpoint connections associated with the storage account.
 *
 * @summary List all the private endpoint connections associated with the storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountListPrivateEndpointConnections.json
 */
async function storageAccountListPrivateEndpointConnections() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const accountName = "sto2527";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

storageAccountListPrivateEndpointConnections().catch(console.error);
```
