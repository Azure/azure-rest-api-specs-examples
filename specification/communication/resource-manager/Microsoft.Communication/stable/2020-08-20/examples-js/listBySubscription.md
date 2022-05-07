Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-communication_3.0.1/sdk/communication/arm-communication/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Handles requests to list all resources in a subscription.
 *
 * @summary Handles requests to list all resources in a subscription.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/listBySubscription.json
 */
async function listBySubscription() {
  const subscriptionId = "12345";
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.communicationService.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listBySubscription().catch(console.error);
```
