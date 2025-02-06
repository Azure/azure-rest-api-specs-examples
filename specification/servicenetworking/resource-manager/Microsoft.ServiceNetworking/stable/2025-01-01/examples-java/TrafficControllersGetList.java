
/**
 * Samples for TrafficControllerInterface List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/TrafficControllersGetList.json
     */
    /**
     * Sample code: Get Traffic Controllers List.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void
        getTrafficControllersList(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.trafficControllerInterfaces().list(com.azure.core.util.Context.NONE);
    }
}
