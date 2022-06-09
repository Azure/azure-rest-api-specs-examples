```javascript
const { MicrosoftSerialConsoleClient } = require("@azure/arm-serialconsole");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Handles requests to list all SerialPort resources in a subscription.
 *
 * @summary Handles requests to list all SerialPort resources in a subscription.
 * x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/ListSerialPortSubscription.json
 */
async function listSerialPortsForSubscription() {
  const subscriptionId = "00000000-00000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSerialConsoleClient(credential, subscriptionId);
  const result = await client.serialPorts.listBySubscriptions();
  console.log(result);
}

listSerialPortsForSubscription().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-serialconsole_2.0.1/sdk/serialconsole/arm-serialconsole/README.md) on how to add the SDK to your project and authenticate.
