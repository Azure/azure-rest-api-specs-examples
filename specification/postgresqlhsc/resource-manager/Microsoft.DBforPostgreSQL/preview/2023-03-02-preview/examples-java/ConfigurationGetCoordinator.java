
/**
 * Samples for Configurations GetCoordinator.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/
     * ConfigurationGetCoordinator.json
     */
    /**
     * Sample code: Get configuration details for coordinator.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void getConfigurationDetailsForCoordinator(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.configurations().getCoordinatorWithResponse("TestResourceGroup", "testcluster", "array_nulls",
            com.azure.core.util.Context.NONE);
    }
}
