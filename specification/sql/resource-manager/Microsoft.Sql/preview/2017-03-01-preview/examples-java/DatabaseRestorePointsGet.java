import com.azure.core.util.Context;

/** Samples for RestorePoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/DatabaseRestorePointsGet.json
     */
    /**
     * Sample code: Gets a database restore point.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsADatabaseRestorePoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getRestorePoints()
            .getWithResponse(
                "Default-SQL-SouthEastAsia", "testserver", "testDatabase", "131546477590000000", Context.NONE);
    }
}
