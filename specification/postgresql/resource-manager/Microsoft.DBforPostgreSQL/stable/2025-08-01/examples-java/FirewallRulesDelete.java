
/**
 * Samples for FirewallRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * FirewallRulesDelete.json
     */
    /**
     * Sample code: Delete an existing firewall rule.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        deleteAnExistingFirewallRule(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.firewallRules().delete("exampleresourcegroup", "exampleserver", "examplefirewallrule",
            com.azure.core.util.Context.NONE);
    }
}
