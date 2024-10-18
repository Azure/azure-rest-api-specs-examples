
/**
 * Samples for VirtualNetworkRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/
     * VirtualNetworkRules_CreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates the specified virtual network rule. During update, the virtual network rule with
     * the specified name will be replaced with this new virtual network rule.
     * 
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void
        createsOrUpdatesTheSpecifiedVirtualNetworkRuleDuringUpdateTheVirtualNetworkRuleWithTheSpecifiedNameWillBeReplacedWithThisNewVirtualNetworkRule(
            com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager.virtualNetworkRules().define("test_virtual_network_rules_name")
            .withExistingAccount("contosorg", "contosoadla").withSubnetId("test_subnetId").create();
    }
}
