Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-batch_7.1.1/sdk/batch/arm-batch/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified private endpoint connection.
 *
 * @summary Gets information about the specified private endpoint connection.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PrivateEndpointConnectionGet.json
 */
async function getPrivateEndpointConnection() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const privateEndpointConnectionName =
    "testprivateEndpointConnection5testprivateEndpointConnection5.24d6b4b5-e65c-4330-bbe9-3a290d62f8e0";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnectionOperations.get(
    resourceGroupName,
    accountName,
    privateEndpointConnectionName
  );
  console.log(result);
}

getPrivateEndpointConnection().catch(console.error);
```
