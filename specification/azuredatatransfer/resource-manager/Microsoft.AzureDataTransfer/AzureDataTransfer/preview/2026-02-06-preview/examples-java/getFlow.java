
/**
 * Samples for Flows Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/getFlow.json
     */
    /**
     * Sample code: Gets flow resource.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void getsFlowResource(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.flows().getWithResponse("testRG", "testConnection", "testFlow", com.azure.core.util.Context.NONE);
    }
}
