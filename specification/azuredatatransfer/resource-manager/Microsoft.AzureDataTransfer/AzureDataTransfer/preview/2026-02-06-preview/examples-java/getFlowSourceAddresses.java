
/**
 * Samples for Flows GetSourceAddresses.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/getFlowSourceAddresses.json
     */
    /**
     * Sample code: Get the flow stream source addresses.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        getTheFlowStreamSourceAddresses(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.flows().getSourceAddressesWithResponse("testRG", "testConnection", "testFlow",
            com.azure.core.util.Context.NONE);
    }
}
