
/**
 * Samples for ManagedDatabaseQueries Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceQueryGet.json
     */
    /**
     * Sample code: Obtain query properties.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void obtainQueryProperties(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseQueries().getWithResponse("sqlcrudtest-7398",
            "sqlcrudtest-4645", "database_1", "42", com.azure.core.util.Context.NONE);
    }
}
