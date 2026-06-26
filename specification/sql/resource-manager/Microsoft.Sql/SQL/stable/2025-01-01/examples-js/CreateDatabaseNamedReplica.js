const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new database or updates an existing database.
 *
 * @summary creates a new database or updates an existing database.
 * x-ms-original-file: 2025-01-01/CreateDatabaseNamedReplica.json
 */
async function createsADatabaseAsNamedReplicaSecondary() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databases.createOrUpdate(
    "Default-SQL-SouthEastAsia",
    "testsvr",
    "testdb",
    {
      location: "southeastasia",
      createMode: "Secondary",
      secondaryType: "Named",
      sourceDatabaseId:
        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-NorthEurope/providers/Microsoft.Sql/servers/testsvr1/databases/primarydb",
      sku: { name: "HS_Gen4", capacity: 2, tier: "Hyperscale" },
    },
  );
  console.log(result);
}
