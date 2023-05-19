/** Samples for FirewallRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/Firewall/preview/2021-12-01-preview/examples/FirewallRuleGet.json
     */
    /**
     * Sample code: Get a firewall rule.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getAFirewallRule(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.firewallRules().getWithResponse("TestGroup", "testserver", "rule1", com.azure.core.util.Context.NONE);
    }
}
