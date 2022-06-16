import com.azure.core.util.Context;

/** Samples for CassandraResources GetCassandraTable. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBCassandraTableGet.json
     */
    /**
     * Sample code: CosmosDBCassandraTableGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCassandraTableGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getCassandraResources()
            .getCassandraTableWithResponse("rg1", "ddb1", "keyspaceName", "tableName", Context.NONE);
    }
}
