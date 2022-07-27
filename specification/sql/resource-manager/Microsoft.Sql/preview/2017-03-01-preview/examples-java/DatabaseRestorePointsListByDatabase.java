import com.azure.core.util.Context;

/** Samples for RestorePoints ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/DatabaseRestorePointsListByDatabase.json
     */
    /**
     * Sample code: List database restore points.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDatabaseRestorePoints(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getRestorePoints()
            .listByDatabase("sqlcrudtest-6730", "sqlcrudtest-9007", "3481", Context.NONE);
    }
}
