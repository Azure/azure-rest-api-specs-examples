const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Permanently deletes the specified managed HSM.
 *
 * @summary Permanently deletes the specified managed HSM.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/DeletedManagedHsm_Purge.json
 */
async function purgeAManagedHsmPool() {
  const subscriptionId =
    process.env["KEYVAULT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const name = "hsm1";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.managedHsms.beginPurgeDeletedAndWait(name, location);
  console.log(result);
}
