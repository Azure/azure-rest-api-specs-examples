import com.azure.resourcemanager.cosmosdbforpostgresql.models.Cluster;

/** Samples for Clusters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterScaleCompute.json
     */
    /**
     * Sample code: Scale compute up or down.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void scaleComputeUpOrDown(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        Cluster resource =
            manager
                .clusters()
                .getByResourceGroupWithResponse("TestGroup", "testcluster", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withNodeVCores(16).apply();
    }
}
