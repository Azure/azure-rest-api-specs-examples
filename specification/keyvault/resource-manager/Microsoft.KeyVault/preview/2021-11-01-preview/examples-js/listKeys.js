const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the keys in the specified key vault.
 *
 * @summary Lists the keys in the specified key vault.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/listKeys.json
 */
async function listKeysInTheVault() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sample-group";
  const vaultName = "sample-vault-name";
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.keys.list(resourceGroupName, vaultName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listKeysInTheVault().catch(console.error);
