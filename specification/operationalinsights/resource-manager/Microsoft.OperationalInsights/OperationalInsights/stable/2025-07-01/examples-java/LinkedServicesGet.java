
/**
 * Samples for LinkedServices Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LinkedServicesGet.json
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
