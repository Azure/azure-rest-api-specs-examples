import com.azure.core.util.Context;

/** Samples for RestorePoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/DatabaseRestorePointsDelete.json
     */
    /**
     * Sample code: Deletes a restore point.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletesARestorePoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getRestorePoints()
            .deleteWithResponse(
                "Default-SQL-SouthEastAsia", "testserver", "testDatabase", "131546477590000000", Context.NONE);
    }
}
