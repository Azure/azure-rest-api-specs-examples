const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new server.
 *
 * @summary creates a new server.
 * x-ms-original-file: 2026-01-01-preview/ServersCreateInYourOwnVirtualNetworkWithSameZoneHighAvailability.json
 */
async function createANewServerInYourOwnVirtualNetworkWithSameZoneHighAvailability() {
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
    highAvailability: { mode: "SameZone" },
    network: {
      delegatedSubnetResourceId:
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.Network/virtualNetworks/examplevirtualnetwork/subnets/examplesubnet",
      privateDnsZoneArmResourceId:
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.Network/privateDnsZones/exampleprivatednszone.private.postgres.database",
    },
    storage: { autoGrow: "Disabled", storageSizeGB: 512, tier: "P20" },
    version: "17",
    sku: { name: "Standard_D4ds_v5", tier: "GeneralPurpose" },
    tags: { InCustomerVnet: "true", InMicrosoftVnet: "false" },
  });
  console.log(result);
}
