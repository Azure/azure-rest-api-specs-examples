
/**
 * Samples for MaintenanceWindowOptionsOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetMaintenanceWindowOptions.json
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
