import com.azure.core.util.Context;

/** Samples for FirewallRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/FirewallRuleGet.json
     */
    /**
     * Sample code: FirewallRuleList.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void firewallRuleList(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.firewallRules().getWithResponse("testrg", "testserver", "rule1", Context.NONE);
    }
}
