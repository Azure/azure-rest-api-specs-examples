import com.azure.core.util.Context;
import java.util.UUID;

/** Samples for DatabaseOperations Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/CancelDatabaseOperation.json
     */
    /**
     * Sample code: Cancel the database management operation.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cancelTheDatabaseManagementOperation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabaseOperations()
            .cancelWithResponse(
                "sqlcrudtest-7398",
                "sqlcrudtest-6661",
                "testdb",
                UUID.fromString("f779414b-e748-4925-8cfe-c8598f7660ae"),
                Context.NONE);
    }
}
