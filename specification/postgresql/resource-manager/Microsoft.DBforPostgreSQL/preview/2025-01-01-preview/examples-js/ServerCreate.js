const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a new server.
 *
 * @summary Creates a new server.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/ServerCreate.json
 */
async function createANewServer() {
  const subscriptionId =
    process.env["POSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "testrg";
  const serverName = "testpgflex";
  const parameters = {
    administratorLogin: "login",
    administratorLoginPassword: "Password1",
    availabilityZone: "1",
    backup: { backupRetentionDays: 7, geoRedundantBackup: "Enabled" },
    createMode: "Create",
    highAvailability: { mode: "ZoneRedundant" },
    location: "eastus",
    network: {
      delegatedSubnetResourceId:
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/test-subnet",
      privateDnsZoneArmResourceId:
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/privateDnsZones/testpgflex.private.postgres.database",
    },
    sku: { name: "Standard_D4ds_v5", tier: "GeneralPurpose" },
    storage: { autoGrow: "Disabled", storageSizeGB: 512, tier: "P20" },
    tags: { vNetServer: "1" },
    version: "16",
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.beginCreateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}
