
/**
 * Samples for MaintenanceWindowOptionsOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/GetMaintenanceWindowOptions.json
     */
    /**
     * Sample code: Gets a list of available maintenance windows for a selected database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAListOfAvailableMaintenanceWindowsForASelectedDatabase(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getMaintenanceWindowOptionsOperations().getWithResponse("Default-SQL-SouthEastAsia",
            "testsvr", "testdb", "current", com.azure.core.util.Context.NONE);
    }
}
