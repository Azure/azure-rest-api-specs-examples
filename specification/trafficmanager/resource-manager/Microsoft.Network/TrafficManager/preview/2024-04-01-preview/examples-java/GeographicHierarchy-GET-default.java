
/**
 * Samples for GeographicHierarchies GetDefault.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/GeographicHierarchy-GET-default.json
     */
    /**
     * Sample code: GeographicHierarchy-GET-default.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void geographicHierarchyGETDefault(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getGeographicHierarchies().getDefaultWithResponse(com.azure.core.util.Context.NONE);
    }
}
