
/**
 * Samples for Maintenances Read.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/Maintenance/preview/2023-10-01-preview/examples/
     * MaintenanceRead.json
     */
    /**
     * Sample code: Read a maintenance.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void readAMaintenance(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.maintenances().readWithResponse("TestGroup", "testserver", "_T9Q-TS8",
            com.azure.core.util.Context.NONE);
    }
}
