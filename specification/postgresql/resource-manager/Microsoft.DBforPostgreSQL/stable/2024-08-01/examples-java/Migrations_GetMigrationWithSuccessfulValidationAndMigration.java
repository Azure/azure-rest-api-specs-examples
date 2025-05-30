
/**
 * Samples for Migrations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/
     * Migrations_GetMigrationWithSuccessfulValidationAndMigration.json
     */
    /**
     * Sample code: Migrations_GetMigrationWithSuccessfulValidationAndMigration.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void migrationsGetMigrationWithSuccessfulValidationAndMigration(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.migrations().getWithResponse("ffffffff-ffff-ffff-ffff-ffffffffffff", "testrg", "testtarget",
            "testmigrationwithsuccessfulvalidationandmigration", com.azure.core.util.Context.NONE);
    }
}
