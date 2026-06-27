const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new database or updates an existing database.
 *
 * @summary creates a new database or updates an existing database.
 * x-ms-original-file: 2025-01-01/CreateDatabaseMaintenanceConfiguration.json
 */
async function createsADatabaseWithPreferredMaintenanceWindow() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databases.createOrUpdate(
    "Default-SQL-SouthEastAsia",
    "testsvr",
    "testdb",
    {
      location: "southeastasia",
      collation: "SQL_Latin1_General_CP1_CI_AS",
      createMode: "Default",
      maintenanceConfigurationId:
        "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_SouthEastAsia_1",
      maxSizeBytes: 1073741824,
      sku: { name: "S2", tier: "Standard" },
    },
  );
  console.log(result);
}
