const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing server. The request body can contain one or multiple of the properties present in the normal server definition.
 *
 * @summary updates an existing server. The request body can contain one or multiple of the properties present in the normal server definition.
 * x-ms-original-file: 2026-01-01-preview/ServersUpdate.json
 */
async function updateAnExistingServer() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.update("exampleresourcegroup", "exampleserver", {
    administratorLoginPassword: "examplenewpassword",
    backup: { backupRetentionDays: 20 },
    createMode: "Update",
    storage: { autoGrow: "Enabled", storageSizeGB: 1024, tier: "P30" },
    sku: { name: "Standard_D8s_v3", tier: "GeneralPurpose" },
  });
  console.log(result);
}
