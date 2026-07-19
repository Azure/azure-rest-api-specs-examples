
import java.util.UUID;

/**
 * Samples for DatabaseOperations Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CancelDatabaseOperation.json
     */
    /**
     * Sample code: Cancel the database management operation.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void cancelTheDatabaseManagementOperation(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseOperations().cancelWithResponse("sqlcrudtest-7398", "sqlcrudtest-6661",
            "testdb", UUID.fromString("f779414b-e748-4925-8cfe-c8598f7660ae"), com.azure.core.util.Context.NONE);
    }
}
