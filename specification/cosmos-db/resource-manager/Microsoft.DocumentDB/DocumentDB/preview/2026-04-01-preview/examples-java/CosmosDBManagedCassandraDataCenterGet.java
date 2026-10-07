
/**
 * Samples for CassandraDataCenters Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBManagedCassandraDataCenterGet.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraDataCenterGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBManagedCassandraDataCenterGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraDataCenters().getWithResponse("cassandra-prod-rg", "cassandra-prod", "dc1",
            com.azure.core.util.Context.NONE);
    }
}
