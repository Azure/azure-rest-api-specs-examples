import com.azure.core.util.Context;

/** Samples for MongoDBResources MigrateMongoDBDatabaseToManualThroughput. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBMongoDBDatabaseMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabaseMigrateToManualThroughput.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBDatabaseMigrateToManualThroughput(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getMongoDBResources()
            .migrateMongoDBDatabaseToManualThroughput("rg1", "ddb1", "databaseName", Context.NONE);
    }
}
