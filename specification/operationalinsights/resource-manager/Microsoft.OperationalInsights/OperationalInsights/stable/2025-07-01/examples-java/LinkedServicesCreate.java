
/**
 * Samples for LinkedServices CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/LinkedServicesCreate.json
     */
    /**
     * Sample code: LinkedServicesCreate.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void linkedServicesCreate(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.linkedServices().define("Cluster").withExistingWorkspace("mms-eus", "TestLinkWS")
            .withWriteAccessResourceId(
                "/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/clusters/testcluster")
            .create();
    }
}
