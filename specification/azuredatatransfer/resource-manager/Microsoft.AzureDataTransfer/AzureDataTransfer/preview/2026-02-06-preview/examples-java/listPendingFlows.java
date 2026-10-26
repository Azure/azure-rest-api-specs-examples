
/**
 * Samples for Connections ListPendingFlowsList.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/listPendingFlows.json
     */
    /**
     * Sample code: Lists a list of pending flows.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void listsAListOfPendingFlows(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.connections().listPendingFlowsList("testRG", "testConnection", com.azure.core.util.Context.NONE);
    }
}
