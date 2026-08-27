
/**
 * Samples for Connections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/Connections_Get.json
     */
    /**
     * Sample code: Get a connection.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getAConnection(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.connections().getWithResponse("exampleRG", "exampleWorkspace", "aksClusterConnection",
            com.azure.core.util.Context.NONE);
    }
}
