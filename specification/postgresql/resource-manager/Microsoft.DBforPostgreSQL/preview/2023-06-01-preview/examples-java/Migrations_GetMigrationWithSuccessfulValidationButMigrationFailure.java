/** Samples for Migrations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/Migrations_GetMigrationWithSuccessfulValidationButMigrationFailure.json
     */
    /**
     * Sample code: Migrations_GetMigrationWithSuccessfulValidationButMigrationFailure.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void migrationsGetMigrationWithSuccessfulValidationButMigrationFailure(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .migrations()
            .getWithResponse(
                "ffffffff-ffff-ffff-ffff-ffffffffffff",
                "testrg",
                "testtarget",
                "testmigrationwithsuccessfulvalidationbutmigrationfailure",
                com.azure.core.util.Context.NONE);
    }
}
