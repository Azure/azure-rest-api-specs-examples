
/**
 * Samples for OutboundFirewallRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/OutboundFirewallRuleCreate.json
     */
    /**
     * Sample code: Approve or reject a outbound firewall rule with a given name.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        approveOrRejectAOutboundFirewallRuleWithAGivenName(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getOutboundFirewallRules().createOrUpdate("sqlcrudtest-7398", "sqlcrudtest-4645",
            "server.database.windows.net", com.azure.core.util.Context.NONE);
    }
}
