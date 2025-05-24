const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a new server.
 *
 * @summary Creates a new server.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/ClusterCreate.json
 */
async function clusterCreate() {
  const subscriptionId =
    process.env["POSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "testrg";
  const serverName = "pgtestcluster";
  const parameters = {
    administratorLogin: "login",
    administratorLoginPassword: "Password1",
    backup: { backupRetentionDays: 7, geoRedundantBackup: "Disabled" },
    cluster: { clusterSize: 2 },
    createMode: "Create",
    highAvailability: { mode: "Disabled" },
    location: "westus",
    network: { publicNetworkAccess: "Disabled" },
    sku: { name: "Standard_D4ds_v5", tier: "GeneralPurpose" },
    storage: { autoGrow: "Disabled", storageSizeGB: 256, tier: "P15" },
    version: "16",
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.beginCreateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}
