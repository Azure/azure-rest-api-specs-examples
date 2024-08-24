
/**
 * Samples for ManagedDatabaseTables Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedDatabaseTableGet.json
     */
    /**
     * Sample code: Get managed database table.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagedDatabaseTable(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseTables().getWithResponse("myRG",
            "myManagedInstanceName", "myDatabase", "dbo", "table1", com.azure.core.util.Context.NONE);
    }
}
