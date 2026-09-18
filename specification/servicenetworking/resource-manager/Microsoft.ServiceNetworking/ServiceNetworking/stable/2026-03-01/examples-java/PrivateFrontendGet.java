
/**
 * Samples for FrontendsInterface Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/PrivateFrontendGet.json
     */
    /**
     * Sample code: Get Private Frontend.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void
        getPrivateFrontend(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.frontendsInterfaces().getWithResponse("rg1", "tc1", "pfe1", com.azure.core.util.Context.NONE);
    }
}
