
/**
 * Samples for PrivateEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2020-03-01-preview/examples/
     * PrivateEndpoint_Get.json
     */
    /**
     * Sample code: Get a private endpoint.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAPrivateEndpoint(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.privateEndpoints().getWithResponse("sjrg", "testcluster", "testpe", com.azure.core.util.Context.NONE);
    }
}
