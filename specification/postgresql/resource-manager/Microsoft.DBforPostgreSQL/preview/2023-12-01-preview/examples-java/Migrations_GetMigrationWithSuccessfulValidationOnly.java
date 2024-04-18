
/**
 * Samples for Migrations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-12-01-preview/examples/
     * Migrations_GetMigrationWithSuccessfulValidationOnly.json
     */
    /**
     * Sample code: Migrations_GetMigrationWithSuccessfulValidationOnly.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void migrationsGetMigrationWithSuccessfulValidationOnly(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.migrations().getWithResponse("ffffffff-ffff-ffff-ffff-ffffffffffff", "testrg", "testtarget",
            "testmigrationwithsuccessfulvalidationonly", com.azure.core.util.Context.NONE);
    }
}
