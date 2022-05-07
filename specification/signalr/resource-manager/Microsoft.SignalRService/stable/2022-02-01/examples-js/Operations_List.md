Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-signalr_5.1.0/sdk/signalr/arm-signalr/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available REST API operations of the Microsoft.SignalRService provider.
 *
 * @summary Lists all of the available REST API operations of the Microsoft.SignalRService provider.
 * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/Operations_List.json
 */
async function operationsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

operationsList().catch(console.error);
```
