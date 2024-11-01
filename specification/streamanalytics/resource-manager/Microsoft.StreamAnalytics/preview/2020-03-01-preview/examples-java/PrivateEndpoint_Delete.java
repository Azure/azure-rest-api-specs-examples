
/**
 * Samples for PrivateEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2020-03-01-preview/examples/
     * PrivateEndpoint_Delete.json
     */
    /**
     * Sample code: Delete a private endpoint.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        deleteAPrivateEndpoint(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.privateEndpoints().delete("sjrg", "testcluster", "testpe", com.azure.core.util.Context.NONE);
    }
}
