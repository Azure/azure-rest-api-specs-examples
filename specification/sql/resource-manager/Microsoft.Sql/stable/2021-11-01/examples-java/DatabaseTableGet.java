
import com.azure.core.util.Context;

/** Samples for DatabaseTables Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseTableGet.json
     */
    /**
     * Sample code: Get database table.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDatabaseTable(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseTables().getWithResponse("myRG", "serverName",
            "myDatabase", "dbo", "table1", Context.NONE);
    }
}
