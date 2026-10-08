
/**
 * Samples for MongoDBResources MigrateMongoDBCollectionToManualThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBMongoDBCollectionMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionMigrateToManualThroughput.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBMongoDBCollectionMigrateToManualThroughput(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().migrateMongoDBCollectionToManualThroughput("rg1", "ddb1",
            "databaseName", "collectionName", com.azure.core.util.Context.NONE);
    }
}
