
/**
 * Samples for TrafficControllerInterface ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/TrafficControllersGet.json
     */
    /**
     * Sample code: Get Traffic Controllers.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void
        getTrafficControllers(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.trafficControllerInterfaces().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
