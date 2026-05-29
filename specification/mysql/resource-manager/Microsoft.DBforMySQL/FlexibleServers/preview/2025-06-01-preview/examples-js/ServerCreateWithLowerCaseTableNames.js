const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new server or updates an existing server. The update action will overwrite the existing server.
 *
 * @summary creates a new server or updates an existing server. The update action will overwrite the existing server.
 * x-ms-original-file: 2025-06-01-preview/ServerCreateWithLowerCaseTableNames.json
 */
async function createAServerWithLowerCaseTableNames() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.create("testrg", "mysqltestserver", {
    sku: { name: "Standard_D2ds_v4", tier: "GeneralPurpose" },
    administratorLogin: "cloudsa",
    administratorLoginPassword: "your_password",
    availabilityZone: "1",
    version: "8.0.21",
    createMode: "Default",
    storage: {
      storageSizeGB: 100,
      iops: 600,
      autoGrow: "Disabled",
      storageRedundancy: "LocalRedundancy",
    },
    backup: { backupRetentionDays: 7, backupIntervalHours: 24, geoRedundantBackup: "Disabled" },
    highAvailability: {
      mode: "ZoneRedundant",
      standbyAvailabilityZone: "3",
      replicationMode: "BinaryLog",
    },
    databasePort: 8888,
    lowerCaseTableNames: 1,
    location: "southeastasia",
    tags: { num: "1" },
  });
  console.log(result);
}
