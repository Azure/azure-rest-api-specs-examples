/** Samples for FirewallRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/FirewallRules_Delete.json
     */
    /**
     * Sample code: Deletes the specified firewall rule from the specified Data Lake Store account.
     *
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void deletesTheSpecifiedFirewallRuleFromTheSpecifiedDataLakeStoreAccount(
        com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager
            .firewallRules()
            .deleteWithResponse("contosorg", "contosoadla", "test_rule", com.azure.core.util.Context.NONE);
    }
}
