
/**
 * Samples for Migrations Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/MigrationsCancel.json
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
