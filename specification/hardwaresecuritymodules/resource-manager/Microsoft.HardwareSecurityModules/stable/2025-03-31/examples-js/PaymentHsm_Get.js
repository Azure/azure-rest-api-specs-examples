const { AzureDedicatedHSMResourceProvider } = require("@azure/arm-hardwaresecuritymodules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the specified Azure dedicated HSM.
 *
 * @summary gets the specified Azure dedicated HSM.
 * x-ms-original-file: 2025-03-31/PaymentHsm_Get.json
 */
async function getAPaymentHSM() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new AzureDedicatedHSMResourceProvider(credential, subscriptionId);
  const result = await client.dedicatedHsm.get("hsm-group", "hsm1");
  console.log(result);
}
