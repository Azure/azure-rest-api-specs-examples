```javascript
const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the resource name is valid and is not already in use.
 *
 * @summary Checks that the resource name is valid and is not already in use.
 * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalR_CheckNameAvailability.json
 */
async function signalRCheckNameAvailability() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const parameters = {
    name: "mySignalRService",
    type: "Microsoft.SignalRService/SignalR",
  };
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const result = await client.signalR.checkNameAvailability(location, parameters);
  console.log(result);
}

signalRCheckNameAvailability().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-signalr_5.1.0/sdk/signalr/arm-signalr/README.md) on how to add the SDK to your project and authenticate.
