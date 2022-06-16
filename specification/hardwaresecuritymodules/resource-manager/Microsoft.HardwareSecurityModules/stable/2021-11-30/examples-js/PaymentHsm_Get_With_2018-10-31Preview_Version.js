const { AzureDedicatedHSMResourceProvider } = require("@azure/arm-hardwaresecuritymodules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified Azure dedicated HSM.
 *
 * @summary Gets the specified Azure dedicated HSM.
 * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/examples/PaymentHsm_Get_With_2018-10-31Preview_Version.json
 */
async function getAPaymentHsmWith20181031PreviewApiVersion() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "hsm-group";
  const name = "hsm1";
  const credential = new DefaultAzureCredential();
  const client = new AzureDedicatedHSMResourceProvider(credential, subscriptionId);
  const result = await client.dedicatedHsmOperations.get(resourceGroupName, name);
  console.log(result);
}

getAPaymentHsmWith20181031PreviewApiVersion().catch(console.error);
