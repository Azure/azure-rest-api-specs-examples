
import java.util.UUID;

/**
 * Samples for ElasticPoolOperations Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CancelElasticPoolOperation.json
     */
    /**
     * Sample code: Cancel the elastic pool management operation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cancelTheElasticPoolManagementOperation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getElasticPoolOperations().cancelWithResponse("sqlcrudtest-7398",
            "sqlcrudtest-6661", "testpool", UUID.fromString("f779414b-e748-4925-8cfe-c8598f7660ae"),
            com.azure.core.util.Context.NONE);
    }
}
