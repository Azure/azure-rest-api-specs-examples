
/**
 * Samples for Databases Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetVCoreDatabase.json
     */
    /**
     * Sample code: Gets a database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsADatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().getWithResponse("Default-SQL-SouthEastAsia",
            "testsvr", "testdb", com.azure.core.util.Context.NONE);
    }
}
