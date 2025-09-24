
/**
 * Samples for MigrationServices GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/
     * GetMigrationService.json
     */
    /**
     * Sample code: Get Migration Service.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void getMigrationService(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.migrationServices().getByResourceGroupWithResponse("testrg", "service1",
            com.azure.core.util.Context.NONE);
    }
}
