import com.azure.core.util.Context;

/** Samples for FirewallRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/FirewallRuleGet.json
     */
    /**
     * Sample code: FirewallRuleGet.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void firewallRuleGet(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.firewallRules().getWithResponse("TestGroup", "testserver", "rule1", Context.NONE);
    }
}
