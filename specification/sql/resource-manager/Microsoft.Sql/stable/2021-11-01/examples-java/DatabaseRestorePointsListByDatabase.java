
/**
 * Samples for RestorePoints ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseRestorePointsListByDatabase.
     * json
     */
    /**
     * Sample code: List database restore points.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDatabaseRestorePoints(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getRestorePoints().listByDatabase("sqlcrudtest-6730",
            "sqlcrudtest-9007", "3481", com.azure.core.util.Context.NONE);
    }
}
