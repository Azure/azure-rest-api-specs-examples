
/**
 * Samples for MongoDBResources MigrateMongoDBDatabaseToAutoscale.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBDatabaseMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabaseMigrateToAutoscale.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBMongoDBDatabaseMigrateToAutoscale(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().migrateMongoDBDatabaseToAutoscale("rg1", "ddb1", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
