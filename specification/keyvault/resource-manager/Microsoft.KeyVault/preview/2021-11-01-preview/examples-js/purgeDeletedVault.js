const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Permanently deletes the specified vault. aka Purges the deleted Azure key vault.
 *
 * @summary Permanently deletes the specified vault. aka Purges the deleted Azure key vault.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/purgeDeletedVault.json
 */
async function purgeADeletedVault() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "sample-vault";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.vaults.beginPurgeDeletedAndWait(vaultName, location);
  console.log(result);
}

purgeADeletedVault().catch(console.error);
