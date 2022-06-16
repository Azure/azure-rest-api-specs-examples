import com.azure.core.util.Context;

/** Samples for FirewallRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/FirewallRuleDelete.json
     */
    /**
     * Sample code: FirewallRuleDelete.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void firewallRuleDelete(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.firewallRules().delete("TestGroup", "testserver", "rule1", Context.NONE);
    }
}
