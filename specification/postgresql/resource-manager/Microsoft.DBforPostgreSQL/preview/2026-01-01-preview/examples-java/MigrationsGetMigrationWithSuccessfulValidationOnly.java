
/**
 * Samples for Migrations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/MigrationsGetMigrationWithSuccessfulValidationOnly.json
     */
    /**
     * Sample code: Get information about a migration with successful validation only.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getInformationAboutAMigrationWithSuccessfulValidationOnly(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.migrations().getWithResponse("exampleresourcegroup", "exampleserver", "examplemigration",
            com.azure.core.util.Context.NONE);
    }
}
