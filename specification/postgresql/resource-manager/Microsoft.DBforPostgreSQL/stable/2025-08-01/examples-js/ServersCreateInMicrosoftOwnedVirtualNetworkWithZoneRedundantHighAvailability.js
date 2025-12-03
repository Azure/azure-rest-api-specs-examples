const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a new server.
 *
 * @summary Creates a new server.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/ServersCreateInMicrosoftOwnedVirtualNetworkWithZoneRedundantHighAvailability.json
 */
async function createANewServerInMicrosoftOwnedVirtualNetworkWithZoneRedundantHighAvailability() {
  const subscriptionId =
    process.env["POSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "exampleresourcegroup";
  const serverName = "exampleserver";
  const parameters = {
    administratorLogin: "exampleadministratorlogin",
    administratorLoginPassword: "examplepassword",
    availabilityZone: "1",
    backup: { backupRetentionDays: 7, geoRedundantBackup: "Enabled" },
    createMode: "Create",
    highAvailability: { mode: "ZoneRedundant" },
    location: "eastus",
    network: { publicNetworkAccess: "Enabled" },
    sku: { name: "Standard_D4ds_v5", tier: "GeneralPurpose" },
    storage: { autoGrow: "Disabled", storageSizeGB: 512, tier: "P20" },
    tags: { inCustomerVnet: "false", inMicrosoftVnet: "true" },
    version: "17",
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    parameters,
  );
  console.log(result);
}
