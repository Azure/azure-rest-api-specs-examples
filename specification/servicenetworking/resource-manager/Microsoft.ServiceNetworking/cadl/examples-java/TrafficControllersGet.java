/** Samples for TrafficControllerInterface ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/TrafficControllersGet.json
     */
    /**
     * Sample code: Get Traffic Controllers.
     *
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void getTrafficControllers(
        com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.trafficControllerInterfaces().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
