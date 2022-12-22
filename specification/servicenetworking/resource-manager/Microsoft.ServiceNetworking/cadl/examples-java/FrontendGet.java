import com.azure.core.util.Context;

/** Samples for FrontendsInterface Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/FrontendGet.json
     */
    /**
     * Sample code: Get Frontend.
     *
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void getFrontend(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.frontendsInterfaces().getWithResponse("rg1", "TC1", "publicIp1", Context.NONE);
    }
}
