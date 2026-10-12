
/**
 * Samples for TrafficManagerUserMetricsKeys Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/TrafficManagerUserMetricsKeys-DELETE.json
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
