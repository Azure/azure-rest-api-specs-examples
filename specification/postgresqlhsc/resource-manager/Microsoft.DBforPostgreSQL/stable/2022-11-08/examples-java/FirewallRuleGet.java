/** Samples for FirewallRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/FirewallRuleGet.json
     */
    /**
     * Sample code: Get the firewall rule of the cluster.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void getTheFirewallRuleOfTheCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.firewallRules().getWithResponse("TestGroup", "pgtestsvc4", "rule1", com.azure.core.util.Context.NONE);
    }
}
