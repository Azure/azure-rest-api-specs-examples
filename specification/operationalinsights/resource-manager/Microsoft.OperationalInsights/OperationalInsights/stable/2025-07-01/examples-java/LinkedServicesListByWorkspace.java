
/**
 * Samples for LinkedServices ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/LinkedServicesListByWorkspace.json
     */
    /**
     * Sample code: LinkedServicesListByWorkspace.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void
        linkedServicesListByWorkspace(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.linkedServices().listByWorkspace("mms-eus", "TestLinkWS", com.azure.core.util.Context.NONE);
    }
}
