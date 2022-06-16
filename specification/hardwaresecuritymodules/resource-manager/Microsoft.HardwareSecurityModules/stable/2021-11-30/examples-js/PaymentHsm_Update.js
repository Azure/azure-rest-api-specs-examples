const { AzureDedicatedHSMResourceProvider } = require("@azure/arm-hardwaresecuritymodules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a dedicated HSM in the specified subscription.
 *
 * @summary Update a dedicated HSM in the specified subscription.
 * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/examples/PaymentHsm_Update.json
 */
async function updateAnExistingPaymentHsm() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "hsm-group";
  const name = "hsm1";
  const tags = { dept: "hsm", environment: "dogfood", slice: "A" };
  const options = { tags };
  const credential = new DefaultAzureCredential();
  const client = new AzureDedicatedHSMResourceProvider(credential, subscriptionId);
  const result = await client.dedicatedHsmOperations.beginUpdateAndWait(
    resourceGroupName,
    name,
    options
  );
  console.log(result);
}

updateAnExistingPaymentHsm().catch(console.error);
