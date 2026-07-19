
/**
 * Samples for FirewallRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FirewallRuleDelete.json
     */
    /**
     * Sample code: Delete a firewall rule.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteAFirewallRule(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getFirewallRules().deleteWithResponse("firewallrulecrudtest-9886",
            "firewallrulecrudtest-2368", "firewallrulecrudtest-7011", com.azure.core.util.Context.NONE);
    }
}
