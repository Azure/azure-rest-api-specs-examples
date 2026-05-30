const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing server. The request body can contain one to many of the properties present in the normal server definition.
 *
 * @summary updates an existing server. The request body can contain one to many of the properties present in the normal server definition.
 * x-ms-original-file: 2025-06-01-preview/ServerUpdateWithBYOK.json
 */
async function updateServerWithByok() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.update("testrg", "mysqltestserver", {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-identity":
          {},
      },
    },
    dataEncryption: {
      type: "AzureKeyVault",
      geoBackupKeyURI: "https://test-geo.vault.azure.net/keys/key/c8a92236622244c0a4fdb892666f671a",
      geoBackupUserAssignedIdentityId:
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-geo-identity",
      primaryKeyURI: "https://test.vault.azure.net/keys/key/c8a92236622244c0a4fdb892666f671a",
      primaryUserAssignedIdentityId:
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-identity",
    },
  });
  console.log(result);
}
