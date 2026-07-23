
/**
 * Samples for CassandraDataCenters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBManagedCassandraDataCenterList.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraDataCenterList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBManagedCassandraDataCenterList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraDataCenters().list("cassandra-prod-rg", "cassandra-prod",
            com.azure.core.util.Context.NONE);
    }
}
