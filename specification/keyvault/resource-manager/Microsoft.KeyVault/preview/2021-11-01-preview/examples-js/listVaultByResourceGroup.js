const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The List operation gets information about the vaults associated with the subscription and within the specified resource group.
 *
 * @summary The List operation gets information about the vaults associated with the subscription and within the specified resource group.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/listVaultByResourceGroup.json
 */
async function listVaultsInTheSpecifiedResourceGroup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sample-group";
  const top = 1;
  const options = { top };
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.vaults.listByResourceGroup(resourceGroupName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listVaultsInTheSpecifiedResourceGroup().catch(console.error);
