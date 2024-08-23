
/**
 * Samples for OutboundFirewallRules ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/OutboundFirewallRuleList.json
     */
    /**
     * Sample code: Gets list of outbound firewall rules on a server.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsListOfOutboundFirewallRulesOnAServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getOutboundFirewallRules().listByServer("sqlcrudtest-7398",
            "sqlcrudtest-4645", com.azure.core.util.Context.NONE);
    }
}
