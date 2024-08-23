
/**
 * Samples for MaintenanceWindowOptionsOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetMaintenanceWindowOptions.json
     */
    /**
     * Sample code: Gets a list of available maintenance windows for a selected database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAListOfAvailableMaintenanceWindowsForASelectedDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getMaintenanceWindowOptionsOperations().getWithResponse(
            "Default-SQL-SouthEastAsia", "testsvr", "testdb", "current", com.azure.core.util.Context.NONE);
    }
}
