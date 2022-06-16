import com.azure.core.util.Context;

/** Samples for FirewallRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/FirewallRuleDelete.json
     */
    /**
     * Sample code: FirewallRuleDelete.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void firewallRuleDelete(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.firewallRules().delete("TestGroup", "testserver", "rule1", Context.NONE);
    }
}
