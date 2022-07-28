import com.azure.core.util.Context;

/** Samples for Databases Resume. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/ResumeDatabase.json
     */
    /**
     * Sample code: Resumes a database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resumesADatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .resume("Default-SQL-SouthEastAsia", "testsvr", "testdwdb", Context.NONE);
    }
}
