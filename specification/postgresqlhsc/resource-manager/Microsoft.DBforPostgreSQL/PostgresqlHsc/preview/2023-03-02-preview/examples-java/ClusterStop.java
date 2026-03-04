
/**
 * Samples for Clusters Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-03-02-preview/ClusterStop.json
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
