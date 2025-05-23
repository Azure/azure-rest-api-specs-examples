const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the specified version of the specified key in the specified key vault.
 *
 * @summary Gets the specified version of the specified key in the specified key vault.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/getKeyVersion.json
 */
async function getAKeyVersion() {
  const subscriptionId =
    process.env["KEYVAULT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["KEYVAULT_RESOURCE_GROUP"] || "sample-group";
  const vaultName = "sample-vault-name";
  const keyName = "sample-key-name";
  const keyVersion = "fd618d9519b74f9aae94ade66b876acc";
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.keys.getVersion(resourceGroupName, vaultName, keyName, keyVersion);
  console.log(result);
}
