import com.azure.core.util.Context;

/** Samples for CassandraResources GetCassandraKeyspaceThroughput. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBCassandraKeyspaceThroughputGet.json
     */
    /**
     * Sample code: CosmosDBCassandraKeyspaceThroughputGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCassandraKeyspaceThroughputGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getCassandraResources()
            .getCassandraKeyspaceThroughputWithResponse("rg1", "ddb1", "keyspaceName", Context.NONE);
    }
}
