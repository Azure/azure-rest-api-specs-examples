import com.azure.resourcemanager.cosmosdbforpostgresql.fluent.models.ServerConfigurationInner;

/** Samples for Configurations UpdateOnCoordinator. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ConfigurationUpdateCoordinator.json
     */
    /**
     * Sample code: Update single configuration of coordinator.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void updateSingleConfigurationOfCoordinator(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager
            .configurations()
            .updateOnCoordinator(
                "TestResourceGroup",
                "testcluster",
                "array_nulls",
                new ServerConfigurationInner().withValue("on"),
                com.azure.core.util.Context.NONE);
    }
}
