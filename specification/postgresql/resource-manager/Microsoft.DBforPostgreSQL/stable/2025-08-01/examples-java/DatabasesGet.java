
/**
 * Samples for Databases Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/DatabasesGet.json
     */
    /**
     * Sample code: Get information about an existing database.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getInformationAboutAnExistingDatabase(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.databases().getWithResponse("exampleresourcegroup", "exampleserver", "exampledatabase",
            com.azure.core.util.Context.NONE);
    }
}
