
/**
 * Samples for FirewallRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/FirewallRuleCreate
     * .json
     */
    /**
     * Sample code: FirewallRuleCreate.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        firewallRuleCreate(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.firewallRules().define("rule1").withExistingFlexibleServer("testrg", "testserver")
            .withStartIpAddress("0.0.0.0").withEndIpAddress("255.255.255.255").create();
    }
}
