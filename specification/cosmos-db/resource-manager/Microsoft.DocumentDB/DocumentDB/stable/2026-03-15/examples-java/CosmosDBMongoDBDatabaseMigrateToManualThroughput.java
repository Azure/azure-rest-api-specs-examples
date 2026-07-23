
/**
 * Samples for MongoDBResources MigrateMongoDBDatabaseToManualThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBDatabaseMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabaseMigrateToManualThroughput.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBMongoDBDatabaseMigrateToManualThroughput(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().migrateMongoDBDatabaseToManualThroughput("rg1", "ddb1",
            "databaseName", com.azure.core.util.Context.NONE);
    }
}
