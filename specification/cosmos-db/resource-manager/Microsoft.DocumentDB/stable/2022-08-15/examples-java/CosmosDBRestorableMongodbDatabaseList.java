import com.azure.core.util.Context;

/** Samples for RestorableMongodbDatabases List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBRestorableMongodbDatabaseList.json
     */
    /**
     * Sample code: CosmosDBRestorableMongodbDatabaseList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBRestorableMongodbDatabaseList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getRestorableMongodbDatabases()
            .list("WestUS", "d9b26648-2f53-4541-b3d8-3044f4f9810d", Context.NONE);
    }
}
