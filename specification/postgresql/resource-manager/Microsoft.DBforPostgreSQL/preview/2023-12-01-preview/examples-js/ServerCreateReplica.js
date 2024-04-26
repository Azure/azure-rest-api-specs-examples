const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new server.
 *
 * @summary Creates a new server.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-12-01-preview/examples/ServerCreateReplica.json
 */
async function serverCreateReplica() {
  const subscriptionId =
    process.env["POSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "testrg";
  const serverName = "pgtestsvc5rep";
  const parameters = {
    createMode: "Replica",
    dataEncryption: {
      type: "AzureKeyVault",
      geoBackupKeyURI: "",
      geoBackupUserAssignedIdentityId: "",
      primaryKeyURI:
        "https://test-kv.vault.azure.net/keys/test-key1/77f57315bab34b0189daa113fbc78787",
      primaryUserAssignedIdentityId:
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity",
    },
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/ffffffffFfffFfffFfffFfffffffffff/resourceGroups/testresourcegroup/providers/MicrosoftManagedIdentity/userAssignedIdentities/testUsermanagedidentity":
          {},
      },
    },
    location: "westus",
    pointInTimeUTC: new Date("2021-06-27T00:04:59.4078005+00:00"),
    sourceServerResourceId:
      "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/sourcepgservername",
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.beginCreateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}
