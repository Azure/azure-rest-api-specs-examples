Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-signalr_5.1.0/sdk/signalr/arm-signalr/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a custom domain.
 *
 * @summary Get a custom domain.
 * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRCustomDomains_Get.json
 */
async function signalRCustomDomainsGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "mySignalRService";
  const name = "example";
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const result = await client.signalRCustomDomains.get(resourceGroupName, resourceName, name);
  console.log(result);
}

signalRCustomDomainsGet().catch(console.error);
```
