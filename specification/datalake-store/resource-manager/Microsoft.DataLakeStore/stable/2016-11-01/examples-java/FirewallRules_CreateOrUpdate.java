
/**
 * Samples for FirewallRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/
     * FirewallRules_CreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates the specified firewall rule. During update, the firewall rule with the specified
     * name will be replaced with this new firewall rule.
     * 
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void
        createsOrUpdatesTheSpecifiedFirewallRuleDuringUpdateTheFirewallRuleWithTheSpecifiedNameWillBeReplacedWithThisNewFirewallRule(
            com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager.firewallRules().define("test_rule").withExistingAccount("contosorg", "contosoadla")
            .withStartIpAddress("1.1.1.1").withEndIpAddress("2.2.2.2").create();
    }
}
