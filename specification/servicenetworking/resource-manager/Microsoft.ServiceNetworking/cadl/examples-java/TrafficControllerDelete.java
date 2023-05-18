/** Samples for TrafficControllerInterface Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/TrafficControllerDelete.json
     */
    /**
     * Sample code: Delete Traffic Controller.
     *
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void deleteTrafficController(
        com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.trafficControllerInterfaces().delete("rg1", "tc1", com.azure.core.util.Context.NONE);
    }
}
