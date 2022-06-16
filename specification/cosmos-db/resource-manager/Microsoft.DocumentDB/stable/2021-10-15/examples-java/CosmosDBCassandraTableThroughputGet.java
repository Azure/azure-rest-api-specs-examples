import com.azure.core.util.Context;

/** Samples for CassandraResources GetCassandraTableThroughput. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBCassandraTableThroughputGet.json
     */
    /**
     * Sample code: CosmosDBCassandraTableThroughputGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCassandraTableThroughputGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getCassandraResources()
            .getCassandraTableThroughputWithResponse("rg1", "ddb1", "keyspaceName", "tableName", Context.NONE);
    }
}
