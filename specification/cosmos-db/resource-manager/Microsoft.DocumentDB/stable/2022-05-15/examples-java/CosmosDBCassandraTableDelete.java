import com.azure.core.util.Context;

/** Samples for CassandraResources DeleteCassandraTable. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBCassandraTableDelete.json
     */
    /**
     * Sample code: CosmosDBCassandraTableDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCassandraTableDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getCassandraResources()
            .deleteCassandraTable("rg1", "ddb1", "keyspaceName", "tableName", Context.NONE);
    }
}
