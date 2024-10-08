
/**
 * Samples for RestorableDroppedDatabases Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetRestorableDroppedDatabase.json
     */
    /**
     * Sample code: Gets a restorable dropped database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsARestorableDroppedDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getRestorableDroppedDatabases().getWithResponse(
            "Default-SQL-SouthEastAsia", "testsvr", "testdb,131403269876900000", com.azure.core.util.Context.NONE);
    }
}
