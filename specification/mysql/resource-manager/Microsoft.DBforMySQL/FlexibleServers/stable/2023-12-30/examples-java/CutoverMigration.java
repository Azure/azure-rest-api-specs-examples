
/**
 * Samples for ServersMigration CutoverMigration.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/
     * CutoverMigration.json
     */
    /**
     * Sample code: Cutover migration for MySQL import.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void
        cutoverMigrationForMySQLImport(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.serversMigrations().cutoverMigration("testrg", "mysqltestserver", com.azure.core.util.Context.NONE);
    }
}
