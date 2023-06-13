import com.azure.resourcemanager.cosmosdbforpostgresql.models.Cluster;

/** Samples for Clusters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterScaleStorage.json
     */
    /**
     * Sample code: Scale up storage.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void scaleUpStorage(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        Cluster resource =
            manager
                .clusters()
                .getByResourceGroupWithResponse("TestGroup", "testcluster", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withNodeStorageQuotaInMb(2097152).apply();
    }
}
