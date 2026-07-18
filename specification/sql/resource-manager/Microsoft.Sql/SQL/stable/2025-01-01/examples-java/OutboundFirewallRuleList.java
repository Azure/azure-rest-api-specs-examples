
/**
 * Samples for OutboundFirewallRules ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/OutboundFirewallRuleList.json
     */
    /**
     * Sample code: Gets list of outbound firewall rules on a server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsListOfOutboundFirewallRulesOnAServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getOutboundFirewallRules().listByServer("sqlcrudtest-7398", "sqlcrudtest-4645",
            com.azure.core.util.Context.NONE);
    }
}
