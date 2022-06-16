import com.azure.core.util.Context;

/** Samples for FirewallRules ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/FirewallRuleListByServer.json
     */
    /**
     * Sample code: FirewallRuleList.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void firewallRuleList(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.firewallRules().listByServer("TestGroup", "testserver", Context.NONE);
    }
}
