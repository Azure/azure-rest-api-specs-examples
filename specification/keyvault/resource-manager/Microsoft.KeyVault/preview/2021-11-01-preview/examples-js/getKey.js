const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the current version of the specified key from the specified key vault.
 *
 * @summary Gets the current version of the specified key from the specified key vault.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/getKey.json
 */
async function getAKey() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sample-group";
  const vaultName = "sample-vault-name";
  const keyName = "sample-key-name";
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.keys.get(resourceGroupName, vaultName, keyName);
  console.log(result);
}

getAKey().catch(console.error);
