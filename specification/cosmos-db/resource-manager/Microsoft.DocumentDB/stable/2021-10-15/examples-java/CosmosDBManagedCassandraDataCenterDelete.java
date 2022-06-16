import com.azure.core.util.Context;

/** Samples for CassandraDataCenters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBManagedCassandraDataCenterDelete.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraDataCenterDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBManagedCassandraDataCenterDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getCassandraDataCenters()
            .delete("cassandra-prod-rg", "cassandra-prod", "dc1", Context.NONE);
    }
}
