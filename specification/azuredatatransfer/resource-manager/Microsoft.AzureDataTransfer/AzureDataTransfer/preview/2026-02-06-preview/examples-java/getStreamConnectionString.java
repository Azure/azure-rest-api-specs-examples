
/**
 * Samples for Flows GetStreamConnectionString.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/getStreamConnectionString.json
     */
    /**
     * Sample code: Get the flow stream connection string.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        getTheFlowStreamConnectionString(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.flows().getStreamConnectionStringWithResponse("testRG", "testConnection", "testFlow",
            com.azure.core.util.Context.NONE);
    }
}
