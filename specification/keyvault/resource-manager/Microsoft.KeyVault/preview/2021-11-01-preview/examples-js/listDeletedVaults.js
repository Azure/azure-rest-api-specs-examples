const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the deleted vaults in a subscription.
 *
 * @summary Gets information about the deleted vaults in a subscription.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/listDeletedVaults.json
 */
async function listDeletedVaultsInTheSpecifiedSubscription() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.vaults.listDeleted()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDeletedVaultsInTheSpecifiedSubscription().catch(console.error);
