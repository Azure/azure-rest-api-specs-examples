
/**
 * Samples for ServerUsages ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerUsageList.json
     */
    /**
     * Sample code: List servers usages.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listServersUsages(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerUsages().listByServer("sqlcrudtest-6730",
            "sqlcrudtest-9007", com.azure.core.util.Context.NONE);
    }
}
