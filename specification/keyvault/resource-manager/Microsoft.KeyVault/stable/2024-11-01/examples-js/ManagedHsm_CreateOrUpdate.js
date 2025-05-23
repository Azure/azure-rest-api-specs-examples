const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update a managed HSM Pool in the specified subscription.
 *
 * @summary Create or update a managed HSM Pool in the specified subscription.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/ManagedHsm_CreateOrUpdate.json
 */
async function createANewManagedHsmPoolOrUpdateAnExistingManagedHsmPool() {
  const subscriptionId =
    process.env["KEYVAULT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["KEYVAULT_RESOURCE_GROUP"] || "hsm-group";
  const name = "hsm1";
  const parameters = {
    location: "westus",
    properties: {
      enablePurgeProtection: false,
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
    parameters,
  );
  console.log(result);
}
