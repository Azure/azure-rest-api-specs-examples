
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.ManagedServerDnsAliasCreation;

/** Samples for ManagedServerDnsAliases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedServerDnsAliasCreateOrUpdate.
     * json
     */
    /**
     * Sample code: Create managed server DNS alias.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createManagedServerDNSAlias(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedServerDnsAliases().createOrUpdate("Default", "dns-mi",
            "dns-alias-mi", new ManagedServerDnsAliasCreation(), Context.NONE);
    }
}
