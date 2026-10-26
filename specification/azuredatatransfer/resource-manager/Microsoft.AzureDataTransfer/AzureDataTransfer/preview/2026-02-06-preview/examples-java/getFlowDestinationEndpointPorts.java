
/**
 * Samples for Flows GetDestinationEndpointPorts.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/getFlowDestinationEndpointPorts.json
     */
    /**
     * Sample code: Get the flow stream destination endpoint ports.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        getTheFlowStreamDestinationEndpointPorts(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.flows().getDestinationEndpointPortsWithResponse("testRG", "testConnection", "testFlow",
            com.azure.core.util.Context.NONE);
    }
}
