
import com.azure.core.util.Context;

/** Samples for Databases ListInaccessibleByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ListVCoreInaccessibleDatabasesByServer.json
     */
    /**
     * Sample code: Gets a list of inaccessible databases in a logical server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getsAListOfInaccessibleDatabasesInALogicalServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases()
            .listInaccessibleByServer("Default-SQL-SouthEastAsia", "testsvr", Context.NONE);
    }
}
