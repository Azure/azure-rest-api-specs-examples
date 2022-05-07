Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-serialconsole_2.0.1/sdk/serialconsole/arm-serialconsole/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MicrosoftSerialConsoleClient } = require("@azure/arm-serialconsole");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Enables the Serial Console service for all VMs and VM scale sets in the provided subscription
 *
 * @summary Enables the Serial Console service for all VMs and VM scale sets in the provided subscription
 * x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/EnableConsoleExamples.json
 */
async function enableSerialConsoleForASubscription() {
  const subscriptionId = "00000000-00000-0000-0000-000000000000";
  const defaultParam = "default";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSerialConsoleClient(credential, subscriptionId);
  const result = await client.enableConsole(defaultParam);
  console.log(result);
}

enableSerialConsoleForASubscription().catch(console.error);
```
