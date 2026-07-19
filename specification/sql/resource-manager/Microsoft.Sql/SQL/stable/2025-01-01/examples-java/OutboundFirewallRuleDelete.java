
/**
 * Samples for OutboundFirewallRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/OutboundFirewallRuleDelete.json
     */
    /**
     * Sample code: Deletes a outbound firewall rule with a given name.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        deletesAOutboundFirewallRuleWithAGivenName(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getOutboundFirewallRules().delete("sqlcrudtest-7398", "sqlcrudtest-6661",
            "server.database.windows.net", com.azure.core.util.Context.NONE);
    }
}
