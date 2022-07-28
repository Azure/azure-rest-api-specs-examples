import com.azure.core.util.Context;

/** Samples for RestorableDroppedDatabases Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/RestorableDroppedDatabaseGet.json
     */
    /**
     * Sample code: Get a restorable dropped database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getARestorableDroppedDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getRestorableDroppedDatabases()
            .getWithResponse(
                "restorabledroppeddatabasetest-1257",
                "restorabledroppeddatabasetest-2389",
                "restorabledroppeddatabasetest-7654,131403269876900000",
                Context.NONE);
    }
}
