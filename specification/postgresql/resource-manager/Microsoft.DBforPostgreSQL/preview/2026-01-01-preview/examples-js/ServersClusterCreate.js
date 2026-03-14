const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new server.
 *
 * @summary creates a new server.
 * x-ms-original-file: 2026-01-01-preview/ServersClusterCreate.json
 */
async function createANewElasticCluster() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.createOrUpdate("exampleresourcegroup", "exampleserver", {
    location: "eastus",
    administratorLogin: "examplelogin",
    administratorLoginPassword: "examplepassword",
    backup: { backupRetentionDays: 7, geoRedundantBackup: "Disabled" },
    cluster: { clusterSize: 2, defaultDatabaseName: "clusterdb" },
    createMode: "Create",
    highAvailability: { mode: "Disabled" },
    network: { publicNetworkAccess: "Disabled" },
    storage: { autoGrow: "Disabled", storageSizeGB: 256, tier: "P15" },
    version: "16",
    sku: { name: "Standard_D4ds_v5", tier: "GeneralPurpose" },
  });
  console.log(result);
}
