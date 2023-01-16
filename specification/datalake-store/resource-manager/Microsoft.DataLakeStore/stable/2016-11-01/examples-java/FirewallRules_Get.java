/** Samples for FirewallRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/FirewallRules_Get.json
     */
    /**
     * Sample code: Gets the specified Data Lake Store firewall rule.
     *
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void getsTheSpecifiedDataLakeStoreFirewallRule(
        com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager
            .firewallRules()
            .getWithResponse("contosorg", "contosoadla", "test_rule", com.azure.core.util.Context.NONE);
    }
}
