
/**
 * Samples for Flows GetDestinationEndpoints.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/getFlowDestinationEndpoints.json
     */
    /**
     * Sample code: Get the flow stream destination endpoints.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        getTheFlowStreamDestinationEndpoints(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.flows().getDestinationEndpointsWithResponse("testRG", "testConnection", "testFlow",
            com.azure.core.util.Context.NONE);
    }
}
