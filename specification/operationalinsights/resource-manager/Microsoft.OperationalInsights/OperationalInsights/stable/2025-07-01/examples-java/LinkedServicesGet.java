
/**
 * Samples for LinkedServices Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/LinkedServicesGet.json
     */
    /**
     * Sample code: LinkedServicesGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void linkedServicesGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.linkedServices().getWithResponse("mms-eus", "TestLinkWS", "Cluster", com.azure.core.util.Context.NONE);
    }
}
