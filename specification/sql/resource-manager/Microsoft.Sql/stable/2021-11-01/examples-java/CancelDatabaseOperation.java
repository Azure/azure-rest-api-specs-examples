
import java.util.UUID;

/**
 * Samples for DatabaseOperations Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CancelDatabaseOperation.json
     */
    /**
     * Sample code: Cancel the database management operation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cancelTheDatabaseManagementOperation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseOperations().cancelWithResponse("sqlcrudtest-7398",
            "sqlcrudtest-6661", "testdb", UUID.fromString("f779414b-e748-4925-8cfe-c8598f7660ae"),
            com.azure.core.util.Context.NONE);
    }
}
