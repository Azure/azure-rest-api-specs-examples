/** Samples for FirewallRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/FirewallRules_Get.json
     */
    /**
     * Sample code: Gets the specified Data Lake Analytics firewall rule.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void getsTheSpecifiedDataLakeAnalyticsFirewallRule(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager
            .firewallRules()
            .getWithResponse("contosorg", "contosoadla", "test_rule", com.azure.core.util.Context.NONE);
    }
}
