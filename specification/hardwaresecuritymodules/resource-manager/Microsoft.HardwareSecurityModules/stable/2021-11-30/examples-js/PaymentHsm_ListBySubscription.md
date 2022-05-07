Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-hardwaresecuritymodules_1.0.0/sdk/hardwaresecuritymodules/arm-hardwaresecuritymodules/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AzureDedicatedHSMResourceProvider } = require("@azure/arm-hardwaresecuritymodules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The List operation gets information about the dedicated HSMs associated with the subscription.
 *
 * @summary The List operation gets information about the dedicated HSMs associated with the subscription.
 * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/examples/PaymentHsm_ListBySubscription.json
 */
async function listDedicatedHsmDevicesInASubscriptionIncludingPaymentHsm() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new AzureDedicatedHSMResourceProvider(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dedicatedHsmOperations.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDedicatedHsmDevicesInASubscriptionIncludingPaymentHsm().catch(console.error);
```
