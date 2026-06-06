
/**
 * Samples for TrafficManagerUserMetricsKeys Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/TrafficManagerUserMetricsKeys-DELETE.json
     */
    /**
     * Sample code: TrafficManagerUserMetricsKeys-DELETE.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void
        trafficManagerUserMetricsKeysDELETE(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getTrafficManagerUserMetricsKeys().deleteWithResponse(com.azure.core.util.Context.NONE);
    }
}
