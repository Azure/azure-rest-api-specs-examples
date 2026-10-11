
/**
 * Samples for TrafficManagerUserMetricsKeys Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/TrafficManagerUserMetricsKeys-GET.json
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
