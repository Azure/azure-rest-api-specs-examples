const { AzureHSMResourceProvider } = require("@azure/arm-hardwaresecuritymodules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The List operation gets information about the dedicated HSMs associated with the subscription.
 *
 * @summary The List operation gets information about the dedicated HSMs associated with the subscription.
 * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-preview/examples/DedicatedHsm_ListBySubscription.json
 */
async function listDedicatedHsmDevicesInASubscription() {
  const subscriptionId =
    process.env["HARDWARESECURITYMODULES_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new AzureHSMResourceProvider(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dedicatedHsmOperations.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
