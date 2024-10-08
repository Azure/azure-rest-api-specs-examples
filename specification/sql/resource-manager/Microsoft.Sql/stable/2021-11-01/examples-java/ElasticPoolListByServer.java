
/**
 * Samples for ElasticPools ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ElasticPoolListByServer.json
     */
    /**
     * Sample code: Get all elastic pools in a server.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAllElasticPoolsInAServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getElasticPools().listByServer("sqlcrudtest-2369",
            "sqlcrudtest-8069", null, com.azure.core.util.Context.NONE);
    }
}
