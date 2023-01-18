/** Samples for VirtualNetworkRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/VirtualNetworkRulesCreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a virtual network rule.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void createOrUpdateAVirtualNetworkRule(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
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
