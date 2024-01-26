
import com.azure.core.util.Context;

/** Samples for ManagedServerDnsAliases Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedServerDnsAliasDelete.json
     */
    /**
     * Sample code: Delete managed server DNS alias.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteManagedServerDNSAlias(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedServerDnsAliases().delete("Default", "dns-mi",
            "dns-alias-mi", Context.NONE);
    }
}
