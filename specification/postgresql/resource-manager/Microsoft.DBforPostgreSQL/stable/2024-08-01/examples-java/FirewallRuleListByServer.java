
/**
 * Samples for FirewallRules ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/
     * FirewallRuleListByServer.json
     */
    /**
     * Sample code: FirewallRuleList.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void firewallRuleList(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.firewallRules().listByServer("testrg", "testserver", com.azure.core.util.Context.NONE);
    }
}
