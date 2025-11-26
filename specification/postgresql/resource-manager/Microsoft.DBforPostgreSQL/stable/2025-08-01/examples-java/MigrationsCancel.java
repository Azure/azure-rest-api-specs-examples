
/**
 * Samples for Migrations Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/MigrationsCancel.
     * json
     */
    /**
     * Sample code: Cancel an active migration.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        cancelAnActiveMigration(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.migrations().cancelWithResponse("exampleresourcegroup", "exampleserver", "examplemigration",
            com.azure.core.util.Context.NONE);
    }
}
