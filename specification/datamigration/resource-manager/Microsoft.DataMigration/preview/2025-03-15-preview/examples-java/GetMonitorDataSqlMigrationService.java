
/**
 * Samples for SqlMigrationServices ListMonitoringData.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * GetMonitorDataSqlMigrationService.json
     */
    /**
     * Sample code: Retrieve the Monitoring Data.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void retrieveTheMonitoringData(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.sqlMigrationServices().listMonitoringDataWithResponse("testrg", "service1",
            com.azure.core.util.Context.NONE);
    }
}
