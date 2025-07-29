
/**
 * Samples for MigrationServices CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * CreateOrUpdateMigrationServiceMAX.json
     */
    /**
     * Sample code: Create or Update Migration Service with maximum parameters.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void createOrUpdateMigrationServiceWithMaximumParameters(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.migrationServices().define("testagent").withRegion("northeurope").withExistingResourceGroup("testrg")
            .create();
    }
}
