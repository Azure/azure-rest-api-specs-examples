
/**
 * Samples for FrontendsInterface Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FrontendGet.json
     */
    /**
     * Sample code: Get Frontend.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void getFrontend(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.frontendsInterfaces().getWithResponse("rg1", "tc1", "fe1", com.azure.core.util.Context.NONE);
    }
}
