/** Samples for VirtualNetworkRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/VirtualNetworkRulesCreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a virtual network rule.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createOrUpdateAVirtualNetworkRule(
        com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager
            .virtualNetworkRules()
            .define("vnet-firewall-rule")
            .withExistingServer("TestGroup", "vnet-test-svr")
            .withVirtualNetworkSubnetId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet")
            .withIgnoreMissingVnetServiceEndpoint(false)
            .create();
    }
}
