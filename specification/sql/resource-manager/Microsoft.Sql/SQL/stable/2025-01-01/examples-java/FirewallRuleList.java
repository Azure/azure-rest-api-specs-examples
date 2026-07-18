
/**
 * Samples for FirewallRules ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FirewallRuleList.json
     */
    /**
     * Sample code: List Firewall Rules.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listFirewallRules(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getFirewallRules().listByServer("firewallrulecrudtest-12", "firewallrulecrudtest-6285",
            com.azure.core.util.Context.NONE);
    }
}
