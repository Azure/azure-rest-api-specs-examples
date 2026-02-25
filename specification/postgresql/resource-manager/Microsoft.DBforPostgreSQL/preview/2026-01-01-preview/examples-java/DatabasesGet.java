
/**
 * Samples for Databases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/DatabasesGet.json
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
