
/**
 * Samples for TrafficManagerUserMetricsKeys CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/TrafficManagerUserMetricsKeys-PUT.json
     */
    /**
     * Sample code: TrafficManagerUserMetricsKeys-PUT.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void
        trafficManagerUserMetricsKeysPUT(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getTrafficManagerUserMetricsKeys()
            .createOrUpdateWithResponse(com.azure.core.util.Context.NONE);
    }
}
