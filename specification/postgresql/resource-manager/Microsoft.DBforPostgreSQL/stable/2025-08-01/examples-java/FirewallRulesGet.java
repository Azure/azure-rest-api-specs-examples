
/**
 * Samples for FirewallRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/FirewallRulesGet.
     * json
     */
    /**
     * Sample code: Get information about a firewall rule in a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getInformationAboutAFirewallRuleInAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.firewallRules().getWithResponse("exampleresourcegroup", "exampleserver", "examplefirewallrule",
            com.azure.core.util.Context.NONE);
    }
}
