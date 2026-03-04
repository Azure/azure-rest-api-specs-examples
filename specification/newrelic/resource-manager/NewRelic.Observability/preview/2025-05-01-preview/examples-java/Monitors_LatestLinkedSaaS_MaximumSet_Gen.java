
/**
 * Samples for Monitors LatestLinkedSaaS.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/Monitors_LatestLinkedSaaS_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_LatestLinkedSaaS_MaximumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsLatestLinkedSaaSMaximumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().latestLinkedSaaSWithResponse("rgopenapi", "ipxmlcbonyxtolzejcjshkmlron",
            com.azure.core.util.Context.NONE);
    }
}
