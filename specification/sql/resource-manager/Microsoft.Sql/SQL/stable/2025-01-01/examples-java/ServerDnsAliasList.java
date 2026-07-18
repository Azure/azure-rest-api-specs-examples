
/**
 * Samples for ServerDnsAliases ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerDnsAliasList.json
     */
    /**
     * Sample code: List server DNS aliases.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listServerDNSAliases(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerDnsAliases().listByServer("Default", "dns-alias-server",
            com.azure.core.util.Context.NONE);
    }
}
