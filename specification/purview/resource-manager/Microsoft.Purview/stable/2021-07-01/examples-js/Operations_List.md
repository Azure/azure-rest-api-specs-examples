Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-purview_1.0.1/sdk/purview/arm-purview/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PurviewManagementClient } = require("@azure/arm-purview");
const { DefaultAzureCredential } = require("@azure/identity");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of available operations
 *
 * @summary List of available operations
 * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Operations_List.json
 */
async function operationsList() {
async function operationsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const credential = new DefaultAzureCredential();
  const client = new PurviewManagementClient(credential, subscriptionId);
  const client = new PurviewManagementClient(credential, subscriptionId);
  const resArray = new Array();
  const resArray = new Array();
  for await (let item of client.operations.list()) {
  for await (let item of client.operations.list()) {
    resArray.push(item);
    resArray.push(item);
  }
  }
  console.log(resArray);
  console.log(resArray);
}
}

operationsList().catch(console.error);
operationsList().catch(console.error);
```
