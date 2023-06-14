/** Samples for Clusters Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterStart.json
     */
    /**
     * Sample code: Start all servers in the cluster.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void startAllServersInTheCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.clusters().start("TestGroup", "testcluster1", com.azure.core.util.Context.NONE);
    }
}
