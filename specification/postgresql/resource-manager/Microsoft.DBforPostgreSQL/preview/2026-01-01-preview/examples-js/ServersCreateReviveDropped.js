const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new server.
 *
 * @summary creates a new server.
 * x-ms-original-file: 2026-01-01-preview/ServersCreateReviveDropped.json
 */
async function createANewServerUsingABackupOfAServerThatWasDeletedOrDroppedRecently() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.createOrUpdate("exampleresourcegroup", "exampleserver", {
    location: "eastus",
    createMode: "ReviveDropped",
    pointInTimeUTC: new Date("2025-06-01T18:30:22.123456Z"),
    sourceServerResourceId:
      "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampledeletedserver",
  });
  console.log(result);
}
