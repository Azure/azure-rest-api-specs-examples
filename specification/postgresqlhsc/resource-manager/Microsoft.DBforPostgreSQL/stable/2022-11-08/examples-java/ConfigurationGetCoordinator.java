/** Samples for Configurations GetCoordinator. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ConfigurationGetCoordinator.json
     */
    /**
     * Sample code: Get configuration details for coordinator.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void getConfigurationDetailsForCoordinator(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager
            .configurations()
            .getCoordinatorWithResponse(
                "TestResourceGroup", "testcluster", "array_nulls", com.azure.core.util.Context.NONE);
    }
}
