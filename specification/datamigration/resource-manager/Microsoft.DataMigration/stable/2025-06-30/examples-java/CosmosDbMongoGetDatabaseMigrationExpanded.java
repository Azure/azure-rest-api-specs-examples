
/**
 * Samples for DatabaseMigrationsMongoToCosmosDbRUMongo Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/
     * CosmosDbMongoGetDatabaseMigrationExpanded.json
     */
    /**
     * Sample code: Get Mongo to CosmosDb Mongo(RU) database Migration with the expand parameter.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void getMongoToCosmosDbMongoRUDatabaseMigrationWithTheExpandParameter(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.databaseMigrationsMongoToCosmosDbRUMongoes().getWithResponse("testrg", "targetCosmosDbClusterName",
            "migrationRequest", com.azure.core.util.Context.NONE);
    }
}
