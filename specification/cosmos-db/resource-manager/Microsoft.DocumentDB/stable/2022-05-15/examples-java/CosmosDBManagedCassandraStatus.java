import com.azure.core.util.Context;

/** Samples for CassandraClusters Status. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBManagedCassandraStatus.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraStatus.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBManagedCassandraStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getCassandraClusters()
            .statusWithResponse("cassandra-prod-rg", "cassandra-prod", Context.NONE);
    }
}
