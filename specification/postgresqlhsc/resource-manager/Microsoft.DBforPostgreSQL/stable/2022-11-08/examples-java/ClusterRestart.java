/** Samples for Clusters Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterRestart.json
     */
    /**
     * Sample code: Restart all servers in the cluster.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void restartAllServersInTheCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.clusters().restart("TestGroup", "testcluster1", com.azure.core.util.Context.NONE);
    }
}
