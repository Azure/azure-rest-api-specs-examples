
/**
 * Samples for Profiles List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/Profile-GET-BySubscription.json
     */
    /**
     * Sample code: ListBySubscription.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void listBySubscription(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getProfiles().list(com.azure.core.util.Context.NONE);
    }
}
