
import com.azure.core.util.Context;

/** Samples for ServerDnsAliases Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerDnsAliasGet.json
     */
    /**
     * Sample code: Get server DNS alias.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getServerDNSAlias(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerDnsAliases().getWithResponse("Default",
            "dns-alias-server", "dns-alias-name-1", Context.NONE);
    }
}
