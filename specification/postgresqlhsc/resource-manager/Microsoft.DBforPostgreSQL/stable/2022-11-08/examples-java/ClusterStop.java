/** Samples for Clusters Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterStop.json
     */
    /**
     * Sample code: Stop all servers in the cluster.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void stopAllServersInTheCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.clusters().stop("TestGroup", "testcluster1", com.azure.core.util.Context.NONE);
    }
}
