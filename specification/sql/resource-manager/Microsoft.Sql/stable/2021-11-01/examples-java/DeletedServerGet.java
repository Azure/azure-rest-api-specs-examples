
/**
 * Samples for DeletedServers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DeletedServerGet.json
     */
    /**
     * Sample code: Get deleted server.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDeletedServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDeletedServers().getWithResponse("japaneast",
            "sqlcrudtest-d-1414", com.azure.core.util.Context.NONE);
    }
}
