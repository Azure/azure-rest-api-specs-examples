const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the private link resources supported for the managed hsm pool.
 *
 * @summary Gets the private link resources supported for the managed hsm pool.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/ManagedHsm_listPrivateLinkResources.json
 */
async function keyVaultListPrivateLinkResources() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sample-group";
  const name = "sample-mhsm";
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.mhsmPrivateLinkResources.listByMhsmResource(resourceGroupName, name);
  console.log(result);
}

keyVaultListPrivateLinkResources().catch(console.error);
