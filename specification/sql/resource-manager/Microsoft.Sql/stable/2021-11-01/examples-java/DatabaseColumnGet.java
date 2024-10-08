
/**
 * Samples for DatabaseColumns Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseColumnGet.json
     */
    /**
     * Sample code: Get database column.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDatabaseColumn(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseColumns().getWithResponse("myRG", "serverName",
            "myDatabase", "dbo", "table1", "column1", com.azure.core.util.Context.NONE);
    }
}
