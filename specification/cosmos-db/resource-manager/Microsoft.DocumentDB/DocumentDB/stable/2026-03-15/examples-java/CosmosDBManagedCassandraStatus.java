
/**
 * Samples for CassandraClusters Status.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBManagedCassandraStatus.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraStatus.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBManagedCassandraStatus(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraClusters().statusWithResponse("cassandra-prod-rg", "cassandra-prod",
            com.azure.core.util.Context.NONE);
    }
}
