
/**
 * Samples for ManagedDatabaseColumns Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedDatabaseColumnGet.json
     */
    /**
     * Sample code: Get managed database column.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagedDatabaseColumn(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseColumns().getWithResponse("myRG",
            "myManagedInstanceName", "myDatabase", "dbo", "table1", "column1", com.azure.core.util.Context.NONE);
    }
}
