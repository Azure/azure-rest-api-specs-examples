
/**
 * Samples for ServerDnsAliases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerDnsAliasCreateOrUpdate.json
     */
    /**
     * Sample code: Create server DNS alias.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createServerDNSAlias(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerDnsAliases().createOrUpdate("Default", "dns-alias-server", "dns-alias-name-1",
            com.azure.core.util.Context.NONE);
    }
}
