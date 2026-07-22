
/**
 * Samples for CassandraDataCenters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBManagedCassandraDataCenterDelete.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraDataCenterDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBManagedCassandraDataCenterDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraDataCenters().delete("cassandra-prod-rg", "cassandra-prod", "dc1",
            com.azure.core.util.Context.NONE);
    }
}
