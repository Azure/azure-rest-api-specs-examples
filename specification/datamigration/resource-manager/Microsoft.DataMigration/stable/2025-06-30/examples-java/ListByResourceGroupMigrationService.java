
/**
 * Samples for MigrationServices ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/
     * ListByResourceGroupMigrationService.json
     */
    /**
     * Sample code: Get Migration Services in the Resource Group.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void
        getMigrationServicesInTheResourceGroup(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.migrationServices().listByResourceGroup("testrg", com.azure.core.util.Context.NONE);
    }
}
