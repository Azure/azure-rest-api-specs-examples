const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new server.
 *
 * @summary creates a new server.
 * x-ms-original-file: 2026-01-01-preview/ServersCreateInMicrosoftOwnedVirtualNetworkWithZoneRedundantHighAvailability.json
 */
async function createANewServerInMicrosoftOwnedVirtualNetworkWithZoneRedundantHighAvailability() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.createOrUpdate("exampleresourcegroup", "exampleserver", {
    location: "eastus",
    administratorLogin: "exampleadministratorlogin",
    administratorLoginPassword: "examplepassword",
    availabilityZone: "1",
    backup: { backupRetentionDays: 7, geoRedundantBackup: "Enabled" },
    createMode: "Create",
    highAvailability: { mode: "ZoneRedundant" },
    network: { publicNetworkAccess: "Enabled" },
    storage: { autoGrow: "Disabled", storageSizeGB: 512, tier: "P20" },
    version: "17",
    sku: { name: "Standard_D4ds_v5", tier: "GeneralPurpose" },
    tags: { InCustomerVnet: "false", InMicrosoftVnet: "true" },
  });
  console.log(result);
}
