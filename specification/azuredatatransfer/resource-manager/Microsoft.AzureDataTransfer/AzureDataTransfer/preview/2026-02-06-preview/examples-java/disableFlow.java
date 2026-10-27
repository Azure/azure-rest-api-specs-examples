
/**
 * Samples for Flows Disable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/disableFlow.json
     */
    /**
     * Sample code: Disables the specified flow.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void disablesTheSpecifiedFlow(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.flows().disable("testRG", "testConnection", "testFlow", com.azure.core.util.Context.NONE);
    }
}
