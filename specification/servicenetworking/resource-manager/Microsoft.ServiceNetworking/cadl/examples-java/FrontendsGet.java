/** Samples for FrontendsInterface ListByTrafficController. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/FrontendsGet.json
     */
    /**
     * Sample code: Get Frontends.
     *
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void getFrontends(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.frontendsInterfaces().listByTrafficController("rg1", "tc1", com.azure.core.util.Context.NONE);
    }
}
