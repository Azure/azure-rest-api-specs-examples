
/**
 * Samples for FirewallRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * FirewallRuleDelete.json
     */
    /**
     * Sample code: FirewallRuleDelete.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        firewallRuleDelete(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.firewallRules().delete("testrg", "testserver", "rule1", com.azure.core.util.Context.NONE);
    }
}
