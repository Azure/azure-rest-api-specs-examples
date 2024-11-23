const { AzureHSMResourceProvider } = require("@azure/arm-hardwaresecuritymodules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified Azure Dedicated HSM.
 *
 * @summary Deletes the specified Azure Dedicated HSM.
 * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-preview/examples/DedicatedHsm_Delete.json
 */
async function deleteADedicatedHsm() {
  const subscriptionId =
    process.env["HARDWARESECURITYMODULES_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["HARDWARESECURITYMODULES_RESOURCE_GROUP"] || "hsm-group";
  const name = "hsm1";
  const credential = new DefaultAzureCredential();
  const client = new AzureHSMResourceProvider(credential, subscriptionId);
  const result = await client.dedicatedHsmOperations.beginDeleteAndWait(resourceGroupName, name);
  console.log(result);
}
