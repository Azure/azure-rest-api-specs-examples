
/**
 * Samples for Configurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-03-02-preview/ConfigurationGet.json
     */
    /**
     * Sample code: Get configuration details.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void
        getConfigurationDetails(com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.configurations().getWithResponse("TestResourceGroup", "testcluster", "client_encoding",
            com.azure.core.util.Context.NONE);
    }
}
