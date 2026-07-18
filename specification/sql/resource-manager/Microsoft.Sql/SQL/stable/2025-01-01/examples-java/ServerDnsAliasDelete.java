
/**
 * Samples for ServerDnsAliases Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerDnsAliasDelete.json
     */
    /**
     * Sample code: Delete server DNS alias.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteServerDNSAlias(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerDnsAliases().delete("Default", "dns-alias-server", "dns-alias-name-1",
            com.azure.core.util.Context.NONE);
    }
}
