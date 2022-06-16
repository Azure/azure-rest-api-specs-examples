import com.azure.core.util.Context;

/** Samples for SqlResources GetSqlDatabaseThroughput. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBSqlDatabaseThroughputGet.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseThroughputGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlDatabaseThroughputGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .getSqlDatabaseThroughputWithResponse("rg1", "ddb1", "databaseName", Context.NONE);
    }
}
