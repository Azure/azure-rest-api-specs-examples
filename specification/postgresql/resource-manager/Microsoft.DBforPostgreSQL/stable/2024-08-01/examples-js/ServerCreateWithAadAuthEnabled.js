const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a new server.
 *
 * @summary Creates a new server.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/ServerCreateWithAadAuthEnabled.json
 */
async function createANewServerWithActiveDirectoryAuthenticationEnabled() {
  const subscriptionId =
    process.env["POSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "testrg";
  const serverName = "pgtestsvc4";
  const parameters = {
    administratorLogin: "cloudsa",
    administratorLoginPassword: "password",
    authConfig: {
      activeDirectoryAuth: "Enabled",
      passwordAuth: "Enabled",
      tenantId: "tttttt-tttt-tttt-tttt-tttttttttttt",
    },
    availabilityZone: "1",
    backup: { backupRetentionDays: 7, geoRedundantBackup: "Disabled" },
    createMode: "Create",
    dataEncryption: { type: "SystemManaged" },
    highAvailability: { mode: "ZoneRedundant" },
    location: "westus",
    network: {
      delegatedSubnetResourceId:
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/test-vnet-subnet",
      privateDnsZoneArmResourceId:
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourcegroups/testrg/providers/Microsoft.Network/privateDnsZones/test-private-dns-zone.postgres.database.azure.com",
    },
    sku: { name: "Standard_D4s_v3", tier: "GeneralPurpose" },
    storage: { autoGrow: "Disabled", storageSizeGB: 512, tier: "P20" },
    tags: { elasticServer: "1" },
    version: "12",
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.beginCreateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}
