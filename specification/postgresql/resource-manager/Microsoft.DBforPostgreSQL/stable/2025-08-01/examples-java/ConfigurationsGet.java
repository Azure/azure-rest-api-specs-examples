
/**
 * Samples for Configurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/ConfigurationsGet.
     * json
     */
    /**
     * Sample code: Get information about a specific configuration (also known as server parameter) of a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getInformationAboutASpecificConfigurationAlsoKnownAsServerParameterOfAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.configurations().getWithResponse("exampleresourcegroup", "exampleserver", "array_nulls",
            com.azure.core.util.Context.NONE);
    }
}
