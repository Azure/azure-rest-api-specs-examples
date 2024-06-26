const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The List operation gets information about the managed HSM Pools associated with the subscription and within the specified resource group.
 *
 * @summary The List operation gets information about the managed HSM Pools associated with the subscription and within the specified resource group.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/ManagedHsm_ListByResourceGroup.json
 */
async function listManagedHsmPoolsInAResourceGroup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "hsm-group";
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedHsms.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagedHsmPoolsInAResourceGroup().catch(console.error);
