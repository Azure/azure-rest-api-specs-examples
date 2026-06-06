
/**
 * Samples for TrafficManagerUserMetricsKeys Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/TrafficManagerUserMetricsKeys-GET.json
     */
    /**
     * Sample code: TrafficManagerUserMetricsKeys-GET.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void
        trafficManagerUserMetricsKeysGET(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getTrafficManagerUserMetricsKeys().getWithResponse(com.azure.core.util.Context.NONE);
    }
}
