/** Samples for FirewallRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/FirewallRules_CreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates the specified firewall rule.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void createsOrUpdatesTheSpecifiedFirewallRule(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager
            .firewallRules()
            .define("test_rule")
            .withExistingAccount("contosorg", "contosoadla")
            .withStartIpAddress("1.1.1.1")
            .withEndIpAddress("2.2.2.2")
            .create();
    }
}
