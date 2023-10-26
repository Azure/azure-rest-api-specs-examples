/** Samples for LinkedServices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/LinkedServicesDelete.json
     */
    /**
     * Sample code: LinkedServicesDelete.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void linkedServicesDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.linkedServices().delete("rg1", "TestLinkWS", "Cluster", com.azure.core.util.Context.NONE);
    }
}
