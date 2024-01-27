
import com.azure.core.util.Context;

/** Samples for ManagedServerDnsAliases ListByManagedInstance. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedServerDnsAliasList.json
     */
    /**
     * Sample code: List managed server DNS aliases.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedServerDNSAliases(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedServerDnsAliases().listByManagedInstance("Default",
            "dns-mi", Context.NONE);
    }
}
