import com.azure.resourcemanager.cosmosdbforpostgresql.models.Cluster;

/** Samples for Clusters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterUpdate.json
     */
    /**
     * Sample code: Update multiple configuration settings of the cluster.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void updateMultipleConfigurationSettingsOfTheCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        Cluster resource =
            manager
                .clusters()
                .getByResourceGroupWithResponse("TestGroup", "testcluster", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withAdministratorLoginPassword("newpassword")
            .withCoordinatorVCores(16)
            .withNodeCount(4)
            .withNodeVCores(16)
            .apply();
    }
}
