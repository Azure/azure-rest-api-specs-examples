const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new server or updates an existing server. The update action will overwrite the existing server.
 *
 * @summary creates a new server or updates an existing server. The update action will overwrite the existing server.
 * x-ms-original-file: 2025-06-01-preview/ServerCreateWithBYOK.json
 */
async function createAServerWithByok() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.create("testrg", "mysqltestserver", {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-identity":
          {},
      },
    },
    location: "southeastasia",
    administratorLogin: "cloudsa",
    administratorLoginPassword: "your_password",
    availabilityZone: "1",
    backup: { backupIntervalHours: 24, backupRetentionDays: 7, geoRedundantBackup: "Disabled" },
    createMode: "Default",
    dataEncryption: {
      type: "AzureKeyVault",
      geoBackupKeyURI: "https://test-geo.vault.azure.net/keys/key/c8a92236622244c0a4fdb892666f671a",
      geoBackupUserAssignedIdentityId:
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-geo-identity",
      primaryKeyURI: "https://test.vault.azure.net/keys/key/c8a92236622244c0a4fdb892666f671a",
      primaryUserAssignedIdentityId:
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-identity",
    },
    highAvailability: { mode: "ZoneRedundant", standbyAvailabilityZone: "3" },
    storage: {
      autoGrow: "Disabled",
      iops: 600,
      storageRedundancy: "LocalRedundancy",
      storageSizeGB: 100,
    },
    version: "5.7",
    sku: { name: "Standard_D2ds_v4", tier: "GeneralPurpose" },
    tags: { num: "1" },
  });
  console.log(result);
}
