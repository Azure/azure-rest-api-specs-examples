
/**
 * Samples for ManagedServerDnsAliases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedServerDnsAliasGet.json
     */
    /**
     * Sample code: Get managed server DNS alias.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getManagedServerDNSAlias(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedServerDnsAliases().getWithResponse("Default", "dns-mi", "dns-alias-mi",
            com.azure.core.util.Context.NONE);
    }
}
