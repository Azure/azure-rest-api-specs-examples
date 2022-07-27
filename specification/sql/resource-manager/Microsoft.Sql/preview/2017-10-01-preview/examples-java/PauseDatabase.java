import com.azure.core.util.Context;

/** Samples for Databases Pause. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/PauseDatabase.json
     */
    /**
     * Sample code: Pauses a database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void pausesADatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .pause("Default-SQL-SouthEastAsia", "testsvr", "testdwdb", Context.NONE);
    }
}
