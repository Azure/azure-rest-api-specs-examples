
/**
 * Samples for Flows Enable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/enableFlow.json
     */
    /**
     * Sample code: Enables the specified flow.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void enablesTheSpecifiedFlow(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.flows().enable("testRG", "testConnection", "testFlow", com.azure.core.util.Context.NONE);
    }
}
