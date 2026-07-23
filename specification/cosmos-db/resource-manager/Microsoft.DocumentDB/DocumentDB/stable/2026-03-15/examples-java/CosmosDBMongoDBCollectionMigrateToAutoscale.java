
/**
 * Samples for MongoDBResources MigrateMongoDBCollectionToAutoscale.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBCollectionMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionMigrateToAutoscale.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBMongoDBCollectionMigrateToAutoscale(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().migrateMongoDBCollectionToAutoscale("rg1", "ddb1", "databaseName",
            "collectionName", com.azure.core.util.Context.NONE);
    }
}
