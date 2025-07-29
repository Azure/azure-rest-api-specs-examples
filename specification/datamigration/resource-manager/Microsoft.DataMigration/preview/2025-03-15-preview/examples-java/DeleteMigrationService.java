
/**
 * Samples for MigrationServices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * DeleteMigrationService.json
     */
    /**
     * Sample code: Delete Migration Service.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void deleteMigrationService(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.migrationServices().delete("testrg", "service1", com.azure.core.util.Context.NONE);
    }
}
