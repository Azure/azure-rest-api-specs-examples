const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new server or updates an existing server. The update action will overwrite the existing server.
 *
 * @summary creates a new server or updates an existing server. The update action will overwrite the existing server.
 * x-ms-original-file: 2025-06-01-preview/ServerCreate.json
 */
async function createANewServer() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.create("testrg", "mysqltestserver", {
    location: "southeastasia",
    administratorLogin: "cloudsa",
    administratorLoginPassword: "your_password",
    availabilityZone: "1",
    backup: { backupIntervalHours: 24, backupRetentionDays: 7, geoRedundantBackup: "Disabled" },
    createMode: "Default",
    highAvailability: { mode: "ZoneRedundant", standbyAvailabilityZone: "3" },
    storage: {
      autoGrow: "Disabled",
      iops: 600,
      storageRedundancy: "ZoneRedundancy",
      storageSizeGB: 100,
    },
    version: "5.7",
    sku: { name: "Standard_D2ds_v4", tier: "GeneralPurpose" },
    tags: { num: "1" },
  });
  console.log(result);
}
