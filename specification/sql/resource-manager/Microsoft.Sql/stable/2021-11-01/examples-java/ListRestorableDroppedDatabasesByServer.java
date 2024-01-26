
import com.azure.core.util.Context;

/** Samples for RestorableDroppedDatabases ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ListRestorableDroppedDatabasesByServer.json
     */
    /**
     * Sample code: Gets a list of restorable dropped databases.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAListOfRestorableDroppedDatabases(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getRestorableDroppedDatabases()
            .listByServer("Default-SQL-SouthEastAsia", "testsvr", Context.NONE);
    }
}
