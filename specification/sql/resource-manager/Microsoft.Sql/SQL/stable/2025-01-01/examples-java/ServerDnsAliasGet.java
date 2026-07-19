
/**
 * Samples for ServerDnsAliases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerDnsAliasGet.json
     */
    /**
     * Sample code: Get server DNS alias.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getServerDNSAlias(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerDnsAliases().getWithResponse("Default", "dns-alias-server", "dns-alias-name-1",
            com.azure.core.util.Context.NONE);
    }
}
