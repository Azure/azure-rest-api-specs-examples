
/**
 * Samples for MigrationServices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/
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
