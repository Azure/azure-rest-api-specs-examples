
/**
 * Samples for Migrations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/Migrations_Delete.
     * json
     */
    /**
     * Sample code: Migrations_Delete.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void migrationsDelete(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.migrations().deleteWithResponse("ffffffff-ffff-ffff-ffff-ffffffffffff", "testrg", "testtarget",
            "testmigration", com.azure.core.util.Context.NONE);
    }
}
