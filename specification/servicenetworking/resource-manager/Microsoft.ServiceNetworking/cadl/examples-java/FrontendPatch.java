import com.azure.resourcemanager.servicenetworking.models.Frontend;

/** Samples for FrontendsInterface Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/FrontendPatch.json
     */
    /**
     * Sample code: Update Frontend.
     *
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void updateFrontend(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        Frontend resource =
            manager
                .frontendsInterfaces()
                .getWithResponse("rg1", "tc1", "fe1", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
