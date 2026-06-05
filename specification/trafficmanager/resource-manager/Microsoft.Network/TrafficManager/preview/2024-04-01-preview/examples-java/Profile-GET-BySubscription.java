
/**
 * Samples for Profiles List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/Profile-GET-BySubscription.json
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
