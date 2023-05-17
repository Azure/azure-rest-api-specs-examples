/** Samples for TrafficControllerInterface GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/TrafficControllerGet.json
     */
    /**
     * Sample code: Get Traffic Controller.
     *
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void getTrafficController(
        com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager
            .trafficControllerInterfaces()
            .getByResourceGroupWithResponse("rg1", "tc1", com.azure.core.util.Context.NONE);
    }
}
