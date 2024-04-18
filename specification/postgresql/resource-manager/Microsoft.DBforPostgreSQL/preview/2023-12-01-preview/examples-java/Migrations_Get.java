
/**
 * Samples for Migrations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-12-01-preview/examples/
     * Migrations_Get.json
     */
    /**
     * Sample code: Migrations_Get.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void migrationsGet(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.migrations().getWithResponse("ffffffff-ffff-ffff-ffff-ffffffffffff", "testrg", "testtarget",
            "testmigration", com.azure.core.util.Context.NONE);
    }
}
