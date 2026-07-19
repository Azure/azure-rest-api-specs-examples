
import com.azure.resourcemanager.sql.models.ManagedServerDnsAliasCreation;

/**
 * Samples for ManagedServerDnsAliases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedServerDnsAliasCreateOrUpdate.json
     */
    /**
     * Sample code: Create managed server DNS alias.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createManagedServerDNSAlias(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedServerDnsAliases().createOrUpdate("Default", "dns-mi", "dns-alias-mi",
            new ManagedServerDnsAliasCreation(), com.azure.core.util.Context.NONE);
    }
}
