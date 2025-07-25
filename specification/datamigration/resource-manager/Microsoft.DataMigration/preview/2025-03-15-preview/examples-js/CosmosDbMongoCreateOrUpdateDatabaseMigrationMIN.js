const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or Update Database Migration resource.
 *
 * @summary Create or Update Database Migration resource.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/CosmosDbMongoCreateOrUpdateDatabaseMigrationMIN.json
 */
async function createMongoToCosmosDbMongoRuDatabaseMigrationResourceWithMinimumParameters() {
  const subscriptionId =
    process.env["DATAMIGRATION_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["DATAMIGRATION_RESOURCE_GROUP"] || "testrg";
  const targetResourceName = "targetCosmosDbClusterName";
  const migrationName = "migrationRequest";
  const parameters = {
    collectionList: [
      {
        sourceCollection: "sourceCol1",
        sourceDatabase: "sourceDb1",
        targetCollection: "targetCol1",
        targetDatabase: "targetDb1",
      },
      { sourceCollection: "sourceCol2", sourceDatabase: "sourceDb2" },
    ],
    kind: "MongoToCosmosDbMongo",
    migrationService:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/MigrationServices/testMigrationService",
    scope:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName",
    sourceMongoConnection: {
      host: "abc.mongodb.com",
      password: "placeholder",
      port: 88,
      useSsl: true,
      userName: "abc",
    },
    targetMongoConnection: { connectionString: "placeholder" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.databaseMigrationsMongoToCosmosDbRUMongo.beginCreateAndWait(
    resourceGroupName,
    targetResourceName,
    migrationName,
    parameters,
  );
  console.log(result);
}
