
/**
 * Samples for SqlMigrationServices ListMigrations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
     * ListMigrationsBySqlMigrationService.json
     */
    /**
     * Sample code: List database migrations attached to the service.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void listDatabaseMigrationsAttachedToTheService(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.sqlMigrationServices().listMigrations("testrg", "service1", com.azure.core.util.Context.NONE);
    }
}
