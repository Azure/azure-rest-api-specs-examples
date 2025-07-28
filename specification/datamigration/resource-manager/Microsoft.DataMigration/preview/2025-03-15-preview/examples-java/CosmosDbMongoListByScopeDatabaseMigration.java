
/**
 * Samples for DatabaseMigrationsMongoToCosmosDbRUMongo GetForScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * CosmosDbMongoListByScopeDatabaseMigration.json
     */
    /**
     * Sample code: Get Mongo to CosmosDb Mongo(RU) database Migration without the expand parameter.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void getMongoToCosmosDbMongoRUDatabaseMigrationWithoutTheExpandParameter(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.databaseMigrationsMongoToCosmosDbRUMongoes().getForScope("testrg", "targetCosmosDbClusterName",
            com.azure.core.util.Context.NONE);
    }
}
