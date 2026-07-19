
import java.util.UUID;

/**
 * Samples for ElasticPoolOperations Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CancelElasticPoolOperation.json
     */
    /**
     * Sample code: Cancel the elastic pool management operation.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void cancelTheElasticPoolManagementOperation(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPoolOperations().cancelWithResponse("sqlcrudtest-7398", "sqlcrudtest-6661",
            "testpool", UUID.fromString("f779414b-e748-4925-8cfe-c8598f7660ae"), com.azure.core.util.Context.NONE);
    }
}
