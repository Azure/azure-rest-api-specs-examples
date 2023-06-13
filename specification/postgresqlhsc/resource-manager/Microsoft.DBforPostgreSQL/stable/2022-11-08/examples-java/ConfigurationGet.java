/** Samples for Configurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ConfigurationGet.json
     */
    /**
     * Sample code: Get configuration details.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void getConfigurationDetails(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager
            .configurations()
            .getWithResponse("TestResourceGroup", "testcluster", "client_encoding", com.azure.core.util.Context.NONE);
    }
}
