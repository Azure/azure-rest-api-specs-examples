
/**
 * Samples for Flows Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/deleteFlow.json
     */
    /**
     * Sample code: Deletes the flow resource.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void deletesTheFlowResource(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.flows().delete("testRG", "testConnection", "testFlow", com.azure.core.util.Context.NONE);
    }
}
