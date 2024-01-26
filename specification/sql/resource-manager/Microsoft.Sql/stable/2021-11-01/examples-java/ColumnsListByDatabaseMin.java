
import com.azure.core.util.Context;

/** Samples for DatabaseColumns ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ColumnsListByDatabaseMin.json
     */
    /**
     * Sample code: List database columns.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDatabaseColumns(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseColumns().listByDatabase("myRG", "serverName",
            "myDatabase", null, null, null, null, null, Context.NONE);
    }
}
