const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the versions of the specified key in the specified managed HSM.
 *
 * @summary Lists the versions of the specified key in the specified managed HSM.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-02-01/examples/managedHsmListKeyVersions.json
 */
async function listKeyVersionsInTheManagedHsm() {
  const subscriptionId =
    process.env["KEYVAULT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["KEYVAULT_RESOURCE_GROUP"] || "sample-group";
  const name = "sample-managedhsm-name";
  const keyName = "sample-key-name";
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedHsmKeys.listVersions(resourceGroupName, name, keyName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
