
/**
 * Samples for ManagedDatabaseColumns ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedColumnsListByDatabaseMin.json
     */
    /**
     * Sample code: List managed database columns.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedDatabaseColumns(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseColumns().listByDatabase("myRG", "serverName",
            "myDatabase", null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
