import com.azure.core.util.Context;

/** Samples for CassandraResources ListCassandraKeyspaces. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBCassandraKeyspaceList.json
     */
    /**
     * Sample code: CosmosDBCassandraKeyspaceList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCassandraKeyspaceList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getCassandraResources()
            .listCassandraKeyspaces("rgName", "ddb1", Context.NONE);
    }
}
