
import com.azure.core.util.Context;

/** Samples for ManagedServerDnsAliases Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedServerDnsAliasGet.json
     */
    /**
     * Sample code: Get managed server DNS alias.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagedServerDNSAlias(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedServerDnsAliases().getWithResponse("Default", "dns-mi",
            "dns-alias-mi", Context.NONE);
    }
}
