
/**
 * Samples for MaintenanceWindowsOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetMaintenanceWindows.json
     */
    /**
     * Sample code: Gets maintenance window settings for a selected database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getsMaintenanceWindowSettingsForASelectedDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getMaintenanceWindowsOperations().getWithResponse(
            "Default-SQL-SouthEastAsia", "testsvr", "testdb", "current", com.azure.core.util.Context.NONE);
    }
}
