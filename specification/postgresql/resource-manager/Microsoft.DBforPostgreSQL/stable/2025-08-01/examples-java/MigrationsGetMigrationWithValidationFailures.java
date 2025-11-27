
/**
 * Samples for Migrations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * MigrationsGetMigrationWithValidationFailures.json
     */
    /**
     * Sample code: Get information about a migration with validation failures.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getInformationAboutAMigrationWithValidationFailures(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.migrations().getWithResponse("exampleresourcegroup", "exampleserver", "examplemigration",
            com.azure.core.util.Context.NONE);
    }
}
