
/**
 * Samples for MigrationServices CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
     * CreateOrUpdateMigrationServiceMIN.json
     */
    /**
     * Sample code: Create or Update Migration Service with minimum parameters.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void createOrUpdateMigrationServiceWithMinimumParameters(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.migrationServices().define("testagent").withRegion("northeurope").withExistingResourceGroup("testrg")
            .create();
    }
}
