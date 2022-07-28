import com.azure.core.util.Context;

/** Samples for ServerDnsAliases ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ServerDnsAliasList.json
     */
    /**
     * Sample code: List server DNS aliases.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listServerDNSAliases(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServerDnsAliases()
            .listByServer("Default", "dns-alias-server", Context.NONE);
    }
}
