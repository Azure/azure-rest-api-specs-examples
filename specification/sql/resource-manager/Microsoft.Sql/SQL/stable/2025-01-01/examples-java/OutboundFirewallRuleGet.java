
/**
 * Samples for OutboundFirewallRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/OutboundFirewallRuleGet.json
     */
    /**
     * Sample code: Gets outbound firewall rule.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsOutboundFirewallRule(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getOutboundFirewallRules().getWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645",
            "server.database.windows.net", com.azure.core.util.Context.NONE);
    }
}
