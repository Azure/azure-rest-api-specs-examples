
/**
 * Samples for IPv6FirewallRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/IPv6FirewallRuleDelete.json
     */
    /**
     * Sample code: Delete an IPv6 firewall rule.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteAnIPv6FirewallRule(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getIPv6FirewallRules().deleteWithResponse("firewallrulecrudtest-9886",
            "firewallrulecrudtest-2368", "firewallrulecrudtest-7011", com.azure.core.util.Context.NONE);
    }
}
