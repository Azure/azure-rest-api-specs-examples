
/**
 * Samples for SqlMigrationServices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * DeleteSqlMigrationService.json
     */
    /**
     * Sample code: Delete SQL Migration Service.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void deleteSQLMigrationService(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.sqlMigrationServices().delete("testrg", "service1", com.azure.core.util.Context.NONE);
    }
}
