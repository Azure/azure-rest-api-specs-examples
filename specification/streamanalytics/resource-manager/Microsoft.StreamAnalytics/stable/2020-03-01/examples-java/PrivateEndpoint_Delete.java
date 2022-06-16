import com.azure.core.util.Context;

/** Samples for PrivateEndpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/PrivateEndpoint_Delete.json
     */
    /**
     * Sample code: Delete a private endpoint.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void deleteAPrivateEndpoint(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.privateEndpoints().delete("sjrg", "testcluster", "testpe", Context.NONE);
    }
}
