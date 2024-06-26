const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The List operation gets information about the private endpoint connections associated with the vault.
 *
 * @summary The List operation gets information about the private endpoint connections associated with the vault.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/listPrivateEndpointConnection.json
 */
async function keyVaultListPrivateEndpointConnection() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sample-group";
  const vaultName = "sample-vault";
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.listByResource(
    resourceGroupName,
    vaultName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

keyVaultListPrivateEndpointConnection().catch(console.error);
