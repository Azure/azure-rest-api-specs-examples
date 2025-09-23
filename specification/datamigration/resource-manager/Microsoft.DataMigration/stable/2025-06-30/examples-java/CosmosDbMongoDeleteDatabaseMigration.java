
/**
 * Samples for DatabaseMigrationsMongoToCosmosDbvCoreMongo Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/
     * CosmosDbMongoDeleteDatabaseMigration.json
     */
    /**
     * Sample code: Delete Mongo to CosmosDb Mongo(vCore) Database Migration resource.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void deleteMongoToCosmosDbMongoVCoreDatabaseMigrationResource(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.databaseMigrationsMongoToCosmosDbvCoreMongoes().delete("testrg", "targetCosmosDbClusterName",
            "migrationRequest", null, com.azure.core.util.Context.NONE);
    }
}
