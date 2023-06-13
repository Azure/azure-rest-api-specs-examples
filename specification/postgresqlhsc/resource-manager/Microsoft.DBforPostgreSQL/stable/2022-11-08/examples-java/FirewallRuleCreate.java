/** Samples for FirewallRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/FirewallRuleCreate.json
     */
    /**
     * Sample code: Create a firewall rule of the cluster.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void createAFirewallRuleOfTheCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager
            .firewallRules()
            .define("rule1")
            .withExistingServerGroupsv2("TestGroup", "pgtestsvc4")
            .withStartIpAddress("0.0.0.0")
            .withEndIpAddress("255.255.255.255")
            .create();
    }
}
