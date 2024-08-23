
/**
 * Samples for ServerDnsAliases ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerDnsAliasList.json
     */
    /**
     * Sample code: List server DNS aliases.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listServerDNSAliases(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerDnsAliases().listByServer("Default", "dns-alias-server",
            com.azure.core.util.Context.NONE);
    }
}
