
import com.azure.resourcemanager.sql.models.ManagedServerDnsAliasAcquisition;

/**
 * Samples for ManagedServerDnsAliases Acquire.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedServerDnsAliasAcquire.json
     */
    /**
     * Sample code: Acquire managed server DNS alias.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void acquireManagedServerDNSAlias(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedServerDnsAliases().acquire("Default", "new-mi",
            "dns-alias-mi",
            new ManagedServerDnsAliasAcquisition().withOldManagedServerDnsAliasResourceId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/old-mi/dnsAliases/alias1"),
            com.azure.core.util.Context.NONE);
    }
}
