
import com.azure.core.util.Context;

/** Samples for DatabaseColumns ListByTable. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseColumnListByTable.json
     */
    /**
     * Sample code: List database columns.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDatabaseColumns(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseColumns().listByTable("myRG", "serverName",
            "myDatabase", "dbo", "table1", null, Context.NONE);
    }
}
