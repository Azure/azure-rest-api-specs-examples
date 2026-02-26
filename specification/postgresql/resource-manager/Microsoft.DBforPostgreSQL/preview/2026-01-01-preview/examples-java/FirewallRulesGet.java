
/**
 * Samples for FirewallRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/FirewallRulesGet.json
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
