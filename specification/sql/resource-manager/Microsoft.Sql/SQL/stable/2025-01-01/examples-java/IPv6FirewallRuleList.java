
/**
 * Samples for IPv6FirewallRules ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/IPv6FirewallRuleList.json
     */
    /**
     * Sample code: List IPv6 Firewall Rules.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listIPv6FirewallRules(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getIPv6FirewallRules().listByServer("firewallrulecrudtest-12",
            "firewallrulecrudtest-6285", com.azure.core.util.Context.NONE);
    }
}
