
/**
 * Samples for ManagedServerDnsAliases Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedServerDnsAliasDelete.json
     */
    /**
     * Sample code: Delete managed server DNS alias.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteManagedServerDNSAlias(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedServerDnsAliases().delete("Default", "dns-mi", "dns-alias-mi",
            com.azure.core.util.Context.NONE);
    }
}
