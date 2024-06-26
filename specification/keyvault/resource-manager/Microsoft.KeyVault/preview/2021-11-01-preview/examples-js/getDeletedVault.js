const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the deleted Azure key vault.
 *
 * @summary Gets the deleted Azure key vault.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/getDeletedVault.json
 */
async function retrieveADeletedVault() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "sample-vault";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.vaults.getDeleted(vaultName, location);
  console.log(result);
}

retrieveADeletedVault().catch(console.error);
