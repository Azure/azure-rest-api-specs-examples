import com.azure.resourcemanager.cosmosdbforpostgresql.fluent.models.ServerConfigurationInner;

/** Samples for Configurations UpdateOnNode. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ConfigurationUpdateNode.json
     */
    /**
     * Sample code: Update single configuration of nodes.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void updateSingleConfigurationOfNodes(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager
            .configurations()
            .updateOnNode(
                "TestResourceGroup",
                "testcluster",
                "array_nulls",
                new ServerConfigurationInner().withValue("off"),
                com.azure.core.util.Context.NONE);
    }
}
