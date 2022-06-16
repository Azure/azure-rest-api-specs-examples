import com.azure.core.util.Context;

/** Samples for FirewallRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/FirewallRuleGet.json
     */
    /**
     * Sample code: FirewallRuleGet.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void firewallRuleGet(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.firewallRules().getWithResponse("TestGroup", "testserver", "rule1", Context.NONE);
    }
}
