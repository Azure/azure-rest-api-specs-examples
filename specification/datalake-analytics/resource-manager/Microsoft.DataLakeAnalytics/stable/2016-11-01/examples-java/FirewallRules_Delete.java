/** Samples for FirewallRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/FirewallRules_Delete.json
     */
    /**
     * Sample code: Deletes the specified firewall rule.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void deletesTheSpecifiedFirewallRule(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager
            .firewallRules()
            .deleteWithResponse("contosorg", "contosoadla", "test_rule", com.azure.core.util.Context.NONE);
    }
}
