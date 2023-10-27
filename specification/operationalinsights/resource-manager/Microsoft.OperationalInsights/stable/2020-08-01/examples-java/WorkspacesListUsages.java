/** Samples for Usages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesListUsages.json
     */
    /**
     * Sample code: UsagesList.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void usagesList(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.usages().list("rg1", "TestLinkWS", com.azure.core.util.Context.NONE);
    }
}
