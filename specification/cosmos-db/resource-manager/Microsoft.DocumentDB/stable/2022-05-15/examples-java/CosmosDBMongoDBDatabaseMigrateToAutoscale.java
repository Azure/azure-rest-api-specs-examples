import com.azure.core.util.Context;

/** Samples for MongoDBResources MigrateMongoDBDatabaseToAutoscale. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBMongoDBDatabaseMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabaseMigrateToAutoscale.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBDatabaseMigrateToAutoscale(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getMongoDBResources()
            .migrateMongoDBDatabaseToAutoscale("rg1", "ddb1", "databaseName", Context.NONE);
    }
}
