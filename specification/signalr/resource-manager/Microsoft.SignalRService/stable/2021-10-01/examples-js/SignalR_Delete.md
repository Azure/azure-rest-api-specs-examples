Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-signalr_5.1.0/sdk/signalr/arm-signalr/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

async function signalRDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "mySignalRService";
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const result = await client.signalR.beginDeleteAndWait(resourceGroupName, resourceName);
  console.log(result);
}

signalRDelete().catch(console.error);
```
