
/**
 * Samples for Migrations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/MigrationsGet.json
     */
    /**
     * Sample code: Get information about a migration.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        getInformationAboutAMigration(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.migrations().getWithResponse("exampleresourcegroup", "exampleserver", "examplemigration",
            com.azure.core.util.Context.NONE);
    }
}
