import com.azure.core.util.Context;

/** Samples for FirewallRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/FirewallRuleDelete.json
     */
    /**
     * Sample code: Delete a firewall rule.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void deleteAFirewallRule(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.firewallRules().delete("TestGroup", "testserver", "rule1", Context.NONE);
    }
}
