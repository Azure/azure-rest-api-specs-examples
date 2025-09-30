
/**
 * Samples for DatabaseMigrationsMongoToCosmosDbvCoreMongo GetForScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
     * CosmosDbMongoListByScopeDatabaseMigration.json
     */
    /**
     * Sample code: Get Mongo to CosmosDb Mongo(vCore) database Migration without the expand parameter.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void getMongoToCosmosDbMongoVCoreDatabaseMigrationWithoutTheExpandParameter(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.databaseMigrationsMongoToCosmosDbvCoreMongoes().getForScope("testrg", "targetCosmosDbClusterName",
            com.azure.core.util.Context.NONE);
    }
}
