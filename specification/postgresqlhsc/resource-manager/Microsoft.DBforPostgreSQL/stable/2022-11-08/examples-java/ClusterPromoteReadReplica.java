/** Samples for Clusters PromoteReadReplica. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterPromoteReadReplica.json
     */
    /**
     * Sample code: Promote read replica cluster to an independent read-write cluster.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void promoteReadReplicaClusterToAnIndependentReadWriteCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.clusters().promoteReadReplica("TestGroup", "testcluster1", com.azure.core.util.Context.NONE);
    }
}
