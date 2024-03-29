const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a managed HSM Pool in the specified subscription.
 *
 * @summary Create or update a managed HSM Pool in the specified subscription.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/ManagedHsm_CreateOrUpdate.json
 */
async function createANewManagedHsmPoolOrUpdateAnExistingManagedHsmPool() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "hsm-group";
  const name = "hsm1";
  const parameters = {
    location: "westus",
    properties: {
      enablePurgeProtection: true,
      enableSoftDelete: true,
      initialAdminObjectIds: ["00000000-0000-0000-0000-000000000000"],
      softDeleteRetentionInDays: 90,
      tenantId: "00000000-0000-0000-0000-000000000000",
    },
    sku: { name: "Standard_B1", family: "B" },
    tags: { dept: "hsm", environment: "dogfood" },
  };
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.managedHsms.beginCreateOrUpdateAndWait(
    resourceGroupName,
    name,
    parameters
  );
  console.log(result);
}

createANewManagedHsmPoolOrUpdateAnExistingManagedHsmPool().catch(console.error);
