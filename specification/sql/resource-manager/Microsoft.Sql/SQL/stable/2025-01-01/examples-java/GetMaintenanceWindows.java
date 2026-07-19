
/**
 * Samples for MaintenanceWindowsOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetMaintenanceWindows.json
     */
    /**
     * Sample code: Gets maintenance window settings for a selected database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsMaintenanceWindowSettingsForASelectedDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getMaintenanceWindowsOperations().getWithResponse("Default-SQL-SouthEastAsia",
            "testsvr", "testdb", "current", com.azure.core.util.Context.NONE);
    }
}
