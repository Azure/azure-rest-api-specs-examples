import com.azure.core.util.Context;

/** Samples for FirewallRules ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/FirewallRuleListByServer.json
     */
    /**
     * Sample code: FirewallRuleList.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void firewallRuleList(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.firewallRules().listByServer("TestGroup", "testserver", Context.NONE);
    }
}
