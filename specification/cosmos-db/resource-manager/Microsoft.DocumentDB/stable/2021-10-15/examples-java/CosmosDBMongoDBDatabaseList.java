import com.azure.core.util.Context;

/** Samples for MongoDBResources ListMongoDBDatabases. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBMongoDBDatabaseList.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabaseList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBDatabaseList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getMongoDBResources()
            .listMongoDBDatabases("rgName", "ddb1", Context.NONE);
    }
}
