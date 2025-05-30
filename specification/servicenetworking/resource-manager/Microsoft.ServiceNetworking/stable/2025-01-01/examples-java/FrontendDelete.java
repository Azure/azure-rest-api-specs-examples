
/**
 * Samples for FrontendsInterface Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FrontendDelete.json
     */
    /**
     * Sample code: Delete Frontend.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void deleteFrontend(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.frontendsInterfaces().delete("rg1", "tc1", "fe1", com.azure.core.util.Context.NONE);
    }
}
