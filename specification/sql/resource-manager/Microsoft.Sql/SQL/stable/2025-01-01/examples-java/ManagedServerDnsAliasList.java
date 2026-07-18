
/**
 * Samples for ManagedServerDnsAliases ListByManagedInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedServerDnsAliasList.json
     */
    /**
     * Sample code: List managed server DNS aliases.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listManagedServerDNSAliases(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedServerDnsAliases().listByManagedInstance("Default", "dns-mi",
            com.azure.core.util.Context.NONE);
    }
}
