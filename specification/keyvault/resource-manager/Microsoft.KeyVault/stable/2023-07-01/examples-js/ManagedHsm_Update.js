const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a managed HSM Pool in the specified subscription.
 *
 * @summary Update a managed HSM Pool in the specified subscription.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/ManagedHsm_Update.json
 */
async function updateAnExistingManagedHsmPool() {
  const subscriptionId =
    process.env["KEYVAULT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["KEYVAULT_RESOURCE_GROUP"] || "hsm-group";
  const name = "hsm1";
  const parameters = {
    tags: { dept: "hsm", environment: "dogfood", slice: "A" },
  };
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.managedHsms.beginUpdateAndWait(resourceGroupName, name, parameters);
  console.log(result);
}
