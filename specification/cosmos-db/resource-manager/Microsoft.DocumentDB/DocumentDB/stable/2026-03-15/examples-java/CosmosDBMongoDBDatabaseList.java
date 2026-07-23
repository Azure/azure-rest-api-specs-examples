
/**
 * Samples for MongoDBResources ListMongoDBDatabases.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBDatabaseList.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabaseList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBDatabaseList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().listMongoDBDatabases("rgName", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
